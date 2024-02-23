package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/yusufcanb/tlm/app"
)

//go:embed VERSION
var version string

func main() {
	tlm := app.New(version)
	if err := tlm.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
