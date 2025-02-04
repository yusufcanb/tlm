package rag

import (
	_ "embed"

	"github.com/yusufcanb/tlm/pkg/rag/internal"
)

//go:embed CONTEXT
var template string

func GetContext(path string, includePatterns []string, excludePatterns []string) (string, string) {
	return internal.PackFilesForLLMs(path, template, []string{}, []string{})
}
