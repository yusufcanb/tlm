package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/yusufcanb/tlama/pkg/app"
)

//go:embed VERSION
var version string

//go:embed Modelfile.explain
var explainModelfile string

//go:embed Modelfile.suggest
var suggestModelfile string

func main() {
	tlm := app.New(version, explainModelfile, suggestModelfile)
	if err := tlm.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
