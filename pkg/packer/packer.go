package packer

import (
	_ "embed"
)

//go:embed xml.tmpl
var xmlTemplate string

type Packer interface {
	Pack(path any, includePatterns []string, ignorePatterns []string) (*Result, error)
	Render(result *Result) (string, error)
	PrintTopFiles(result *Result, top int)
	PrintContextSummary(result *Result)
}

func New() Packer {
	return &directoryPacker{
		template: xmlTemplate,
	}
}
