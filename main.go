package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/yusufcanb/tlm/pkg/app"
	"github.com/yusufcanb/tlm/pkg/shell"
)

//go:embed VERSION
var version string
var sha1ver string

func main() {
	shell.Version = version
	tlm := app.New(version, sha1ver)
	if err := tlm.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
