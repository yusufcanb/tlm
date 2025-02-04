package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"text/template"
	"time"
)

// PackFilesForLLMs walks the directory tree starting at "path" and, for every file
// that passes the include/exclude filters, reads its content. It then renders a merged
// output using the text/template package.
func PackFilesForLLMs(path string, outputTemplate string, includePatterns []string, excludePatterns []string) (string, string) {
	// Collect file entries and build a directory tree.
	fileEntries, rootNode, err := getFileEntries(path, includePatterns, excludePatterns)
	if err != nil {
		return "", fmt.Sprintf("Error walking the path: %v", err)
	}

	// Sort file entries by their relative path.
	sort.Slice(fileEntries, func(i, j int) bool {
		return fileEntries[i].RelPath < fileEntries[j].RelPath
	})

	// Build a formatted string of the directory structure.
	dirStructure := buildDirStructure(rootNode, "")

	// Prepare the data for the template.
	data := TemplateData{
		GeneratedAt:        time.Now().UTC().Format(time.RFC3339Nano),
		DirectoryStructure: dirStructure,
		Files:              fileEntries,
	}

	// Parse and execute the template.
	tmpl, err := template.New("packed").Parse(outputTemplate)
	if err != nil {
		return "", fmt.Sprintf("Error parsing template: %v", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Sprintf("Error executing template: %v", err)
	}

	return dirStructure, buf.String()

}
