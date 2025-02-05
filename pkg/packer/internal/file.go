package internal

import (
	"os"

	"github.com/bmatcuk/doublestar/v4"
)

// GetContextFilePaths returns the file paths in the given directory that match the include patterns and do not match the exclude patterns.
func GetContextFilePaths(path string, includePatterns []string, excludePatterns []string) ([]string, error) {
	dirFs := os.DirFS(path)
	filePaths := make([]string, 0)

	if len(includePatterns) == 0 {
		includePatterns = []string{"**/*"}
	}

	// Get all files that match the include patterns
	for _, ip := range includePatterns {
		paths, err := doublestar.Glob(dirFs, ip)
		if err != nil {
			return nil, err
		}
		filePaths = append(filePaths, paths...)
	}

	// Exclude files that match the exclude patterns
	for _, ep := range excludePatterns {
		for _, fp := range filePaths {
			shouldExclude, err := doublestar.PathMatch(ep, fp)
			if err != nil {
				return nil, err
			}

			// if the file should be excluded, remove it from the list
			if shouldExclude {

				for i, path := range filePaths {
					if path == fp {
						filePaths = append(filePaths[:i], filePaths[i+1:]...)
						break
					}
				}

				// fmt.Println("Excluded file: ", fp)
			}
		}
	}

	return filePaths, nil
}

// GetFileContent returns the content of a file and the number of tokens in it.
func GetFileContent(filePath string) (string, int, int, error) {
	// Get the file content
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", -1, -1, err
	}

	return string(b), GetTokenCount(string(b)), len(b), nil
}
