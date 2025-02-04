package internal

import (
	"fmt"
	"testing"
)

func TestShouldExclude(t *testing.T) {
	tests := []struct {
		name     string   // file path to test
		patterns []string // additional exclude patterns
		expected bool     // expected outcome: true if the file should be excluded
	}{
		{"e2e/file.txt", []string{"e2e/.venv/**"}, false},
		{"e2e/markdown.md", []string{"e2e/.venv/**"}, false},
		{"e2e/sub/test.py", []string{"e2e/.venv/**"}, false},
		{"e2e/sub/test2.py", []string{"e2e/.venv/**"}, false},
		{"e2e/.venv/Scripts/activate.bat", []string{"e2e/.venv/**"}, true},
		{"e2e/.venv/Lib/site-packages/requests-2.32.3.dist-info/WHEEL", []string{}, true},
		{"src/temp/file.txt", []string{"**/temp/**"}, true},  // Matches a "temp" folder anywhere.
		{"src/Temp/file.txt", []string{"**/temp/**"}, false}, // Case-sensitive mismatch.
		{"node_modules/package/index.js", []string{"node_modules/**"}, true},
		{"vendor/package/file.go", []string{"vendor/**"}, true},
		{".git/config", []string{}, true},                          // Excluded via defaultIgnoreList (".git/**").
		{"docs/manual.md", []string{"**/temp/**"}, false},          // Does not match any exclude pattern.
		{"build/output.log", []string{"**/*.log"}, true},           // Matches log file exclusion.
		{"cache/data.tmp", []string{"cache/**", "**/*.tmp"}, true}, // Matches multiple exclude patterns.

		// Windows-style path: backslashes should be normalized.
		{"node_modules\\lib\\index.js", []string{"node_modules/**"}, true},

		// Single-character wildcard: "?" matching.
		{"logs/app1.log", []string{"logs/app?.log"}, true},
		{"logs/app.log", []string{"logs/app?.log"}, true},

		// Matching files with a specific extension via wildcard.
		{"backup/file.bak", []string{"**/*.bak"}, true},

		// Nested directories: using ** vs. a single *.
		{"a/b/c/d.txt", []string{"a/**/c/d.txt"}, true},
		{"a/b/c/d.txt", []string{"a/*/c/d.txt"}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf(" (%s)", tt.name), func(t *testing.T) {
			// Merge custom patterns with the default ignore list.
			result := shouldExclude(tt.name, append(tt.patterns, defaultIgnoreList...))
			if result != tt.expected {
				t.Errorf("shouldExclude(%q, patterns=%v) = %v; want %v", tt.name, tt.patterns, result, tt.expected)
			}
		})
	}
}

func TestShouldInclude(t *testing.T) {
	tests := []struct {
		name     string   // file path to test
		patterns []string // include patterns to use
		expected bool     // expected outcome: true if the file should be included
	}{
		// Original test cases:
		{"file.txt", []string{"*.txt"}, true},
		{"file.txt", []string{"*.md"}, false},
		{"file.md", []string{"*.md", "*.txt"}, true},
		{"file.md", []string{}, false},

		// Extended test cases:
		{"src/main.go", []string{"src/*.go"}, true},        // Direct match in "src" folder.
		{"src/sub/main.go", []string{"src/*.go"}, false},   // Does not match files in a subdirectory.
		{"src/sub/main.go", []string{"src/**/*.go"}, true}, // ** matches nested directories.
		{"FILE.TXT", []string{"*.txt"}, false},             // Case-sensitive mismatch.
		{"scripts/build.sh", []string{"scripts/*.sh"}, true},
		{"scripts/build.sh", []string{"scripts/*.js"}, false},
		{"docs/README.md", []string{"**/README.md"}, true},
		{"docs/README.MD", []string{"**/README.md"}, false}, // Case-sensitive mismatch.
		{"assets/image.png", []string{"assets/*.png", "*.jpg"}, true},
		{"assets/image.jpeg", []string{"assets/*.png", "*.jpg"}, false},

		// Using "?" wildcard and character classes.
		{"file1.txt", []string{"file?.txt"}, true},
		{"file12.txt", []string{"file?.txt"}, false},
		{"file1.txt", []string{"file[0-9].txt"}, true},
		{"file10.txt", []string{"file[0-9].txt"}, false},

		// Nested directories with specific patterns.
		{"a/b/c.go", []string{"a/*/*.go"}, true},
		{"a/b/d/c.go", []string{"a/*/*.go"}, false},

		// Windows-style paths.
		{"src\\main.go", []string{"src/*.go"}, true},
		{"src\\sub\\main.go", []string{"src/**/*.go"}, true},

		// Matching only the file name in a nested directory.
		{"dir1/dir2/file.txt", []string{"file.txt"}, false},
		{"dir1/dir2/file.txt", []string{"**/file.txt"}, true},

		// Using universal wildcards: "*" should not match across directories.
		{"file.txt", []string{"*"}, true},
		{"dir/file.txt", []string{"*"}, false},
		{"dir/file.txt", []string{"**"}, true},

		// Hidden file matching.
		{".env", []string{".*"}, true},
		{".env", []string{"env"}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf(" (%s)", tt.name), func(t *testing.T) {
			result := shouldInclude(tt.name, tt.patterns)
			if result != tt.expected {
				t.Errorf("shouldInclude(%q, patterns=%v) = %v; want %v", tt.name, tt.patterns, result, tt.expected)
			}
		})
	}
}
