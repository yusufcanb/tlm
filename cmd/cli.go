package main

import (
	"log"
	"os"
	"tlama/pkg/app"
)

func main() {
	tlama := app.New()
	if err := tlama.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
