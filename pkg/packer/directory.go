package packer

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"text/template"
	"time"

	"github.com/yusufcanb/tlm/pkg/packer/internal"
)

type directoryPacker struct {
	template string
	files    []File
}

func (dp *directoryPacker) Pack(path any, includePatterns []string, ignorePatterns []string) (*Result, error) {
	var numTokens int = 0
	var numChars int = 0

	path, ok := path.(string)
	if !ok {
		return nil, errors.New("path is not a string")
	}

	filePaths, err := internal.GetContextFilePaths(path.(string), includePatterns, ignorePatterns)
	if err != nil {
		return nil, errors.New("failed to get context files")
	}

	for _, fp := range filePaths {
		content, tokens, chars, err := internal.GetFileContent(fp)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return nil, errors.New("failed to get file content")
		}

		dp.files = append(dp.files, File{
			Path:    fp,
			Content: content,
			Chars:   chars,
			Tokens:  tokens,
		})

		numTokens += tokens
		numChars += chars
	}

	return &Result{
		Files:  dp.files,
		Tokens: numTokens,
		Chars:  numChars,
	}, nil
}

func (dp *directoryPacker) Render(result *Result) (string, error) {

	// Prepare the data for the template.
	data := Result{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339Nano),
		Files:       result.Files,
		Tokens:      result.Tokens,
	}

	// Parse and execute the template.
	tmpl, err := template.New("packed").Parse(dp.template)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("Error executing template: %v", err)
	}

	return buf.String(), nil
}

func (dp *directoryPacker) PrintTopFiles(result *Result, top int) {
	// Print the top 5 files
	fmt.Printf("📈 Top 5 Files by Character Count and Token Count:\n──────────────────────────────────────────────────\n")
	sort.Slice(result.Files, func(i, j int) bool {
		return result.Files[i].Tokens > result.Files[j].Tokens
	})

	// Print the top 5 files
	for i, file := range result.Files {
		if i >= 5 {
			break
		}
		fmt.Printf("%d.  %s (%d chars, %d tokens)\n", i+1, file.Path, file.Chars, file.Tokens)
	}
	fmt.Println()
}

func (dp *directoryPacker) PrintContextSummary(result *Result) {
	fmt.Printf("📊 Context Summary:\n───────────────────\n")
	fmt.Printf("Total Files:\t%d\nTotal Chars:\t%d\nTotal Tokens:\t%d\n\n", len(result.Files), result.Chars, result.Tokens)
}
