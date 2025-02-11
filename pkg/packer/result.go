package packer

type File struct {
	Path    string // relative path
	Content string // file content
	Chars   int    // number of characters
	Tokens  int    // number of tokens
}

func (f *File) IsBinary() bool {
	return f.Chars < 0 && f.Tokens < 0 && f.Content == "__BINARY__"
}

type Result struct {
	GeneratedAt string // generation time
	Files       []File // list of file paths
	Chars       int    // number of characters
	Tokens      int    // number of tokens
}
