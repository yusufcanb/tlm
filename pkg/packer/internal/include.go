package internal

import (
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

// shouldInclude returns true if the file path matches at least one of the include patterns.
func shouldInclude(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if match, _ := doublestar.Match(pattern, filepath.ToSlash(path)); match {
			return true
		}
	}
	return false
}
