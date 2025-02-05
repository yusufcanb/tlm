package internal

import (
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

// shouldExclude returns true if the file path matches any of the exclude patterns.
func shouldExclude(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if match, _ := doublestar.Match(pattern, filepath.ToSlash(path)); match {
			return true
		}
	}
	return false
}
