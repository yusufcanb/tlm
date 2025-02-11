package internal

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/bmatcuk/doublestar/v4"
)

// GetContextFilePaths returns the file paths in the given directory that match the include patterns and do not match the exclude patterns.
func GetContextFilePaths(path string, includePatterns []string, excludePatterns []string) ([]string, error) {
	dirFs := os.DirFS(path)
	filePathSet := make(map[string]struct{})

	if len(includePatterns) == 0 {
		includePatterns = []string{"**/*"}
	}

	// Get all files that match the include patterns
	for _, ip := range includePatterns {
		paths, err := doublestar.Glob(dirFs, ip)
		if err != nil {
			return nil, err
		}
		for _, p := range paths {
			// Check if the path corresponds to a file.
			info, err := fs.Stat(dirFs, p)
			if err != nil {
				return nil, fmt.Errorf("failed to stat %s: %w", p, err)
			}
			// Skip if it's a directory.
			if info.IsDir() {
				continue
			}
			filePathSet[p] = struct{}{}
		}
	}

	// Exclude files that match the exclude patterns
	for _, ep := range append(excludePatterns, defaultIgnoreList...) {
		for fp := range filePathSet {
			shouldExclude, err := doublestar.PathMatch(ep, fp)
			if err != nil {
				return nil, err
			}

			// if the file should be excluded, remove it from the list
			if shouldExclude {
				delete(filePathSet, fp)
			}
		}
	}

	// Convert the set into a slice.
	filePaths := make([]string, 0, len(filePathSet))
	for fp := range filePathSet {
		filePaths = append(filePaths, fp)
		delete(filePathSet, fp)
	}

	return filePaths, nil
}

// GetFileContent returns the content of a file and the number of tokens in it.
func GetFileContent(baseDir string, filePath string) (string, int, int, error) {
	b, err := os.ReadFile(path.Join(baseDir, filePath))
	if err != nil {
		return "", -1, -1, fmt.Errorf("failed to read file content: %s ", err.Error())
	}

	if isBinary(b) {
		return "__BINARY__", 0, 0, nil
	}

	return string(b), GetTokenCount(string(b)), len(b), nil
}
