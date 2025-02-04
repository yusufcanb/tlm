package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

// getFileEntries walks the file system starting at rootPath, filters files based on the
// include and exclude patterns, reads file contents, and builds a directory tree.
func getFileEntries(rootPath string, includePatterns []string, excludePatterns []string) ([]FileEntry, *Node, error) {
	var fileEntries []FileEntry
	rootNode := &Node{
		Name:     rootPath,
		IsDir:    true,
		Children: make(map[string]*Node),
	}

	err := filepath.Walk(rootPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Propagate the error.
		}
		if info.IsDir() {
			return nil // Skip directories.
		}

		// base := filepath.Base(filePath)g
		//merge excludePatterns with defaultIgnoreList
		if shouldExclude(filePath, append(excludePatterns, defaultIgnoreList...)) {
			return nil
		}
		if len(includePatterns) > 0 && !shouldInclude(filePath, includePatterns) {
			return nil
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			// Optionally log the error; here we silently skip unreadable files.
			return nil
		}

		relPath, err := filepath.Rel(rootPath, filePath)
		if err != nil {
			relPath = filePath // Fallback to full path if relative conversion fails.
		}

		// Always insert the file into the directory tree.
		insertPath(rootNode, relPath)

		// If the file is binary, do not add it to the fileEntries.
		if isBinary(data) {
			return nil
		}

		fileEntries = append(fileEntries, FileEntry{
			RelPath: relPath,
			Content: string(data),
		})
		return nil
	})

	return fileEntries, rootNode, err
}

// insertPath adds a file path to the directory tree rooted at node.
func insertPath(root *Node, relPath string) {
	parts := strings.Split(relPath, string(os.PathSeparator))
	current := root
	for i, part := range parts {
		if part == "" {
			continue
		}
		isDir := i < len(parts)-1
		if child, exists := current.Children[part]; exists {
			current = child
		} else {
			newNode := &Node{
				Name:     part,
				IsDir:    isDir,
				Children: make(map[string]*Node),
			}
			current.Children[part] = newNode
			current = newNode
		}
	}
}

// buildDirStructure returns an indented string representation of the directory tree.
func buildDirStructure(node *Node, parentPath string) string {
	var sb strings.Builder
	var fullPath string

	// Build the full path for the current node.
	if node.Name != "" {
		if parentPath == "" {
			fullPath = node.Name
		} else {
			fullPath = parentPath + "/" + node.Name
		}

		// Append a trailing slash for directories.
		if node.IsDir {
			// sb.WriteString(fmt.Sprintf("%s/\n", fullPath))
		} else {
			sb.WriteString(fmt.Sprintf("%s\n", fullPath))
		}
	} else {
		// If node.Name is empty (for example, the root node),
		// then just use the parentPath.
		fullPath = parentPath
	}

	// Process children in sorted order.
	var keys []string
	for k := range node.Children {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sb.WriteString(buildDirStructure(node.Children[k], fullPath))
	}

	return sb.String()
}

// shouldExclude returns true if the file path matches any of the exclude patterns.
func shouldExclude(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if match, _ := doublestar.Match(pattern, filepath.ToSlash(path)); match {
			return true
		}
	}
	return false
}

// shouldInclude returns true if the file path matches at least one of the include patterns.
func shouldInclude(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if match, _ := doublestar.Match(pattern, filepath.ToSlash(path)); match {
			return true
		}
	}
	return false
}
