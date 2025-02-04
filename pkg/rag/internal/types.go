package internal

// FileEntry holds a file's relative path and its content.
type FileEntry struct {
	RelPath string
	Content string
}

// Node is used to build the directory tree.
type Node struct {
	Name     string
	IsDir    bool
	Children map[string]*Node
}

// TemplateData is used to pass all needed information to the text/template.
type TemplateData struct {
	GeneratedAt        string
	DirectoryStructure string
	Files              []FileEntry
}
