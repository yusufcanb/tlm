package main

import (
	"log"
	"os"

	"github.com/yusufcanb/tlama/pkg/app"
)

var version = "1.0"

func main() {
	tlama := app.New(version)
	if err := tlama.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
