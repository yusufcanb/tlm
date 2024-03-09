package app_test

import (
	"github.com/yusufcanb/tlm/app"
	"os"
	"testing"
)

func run(args []string) {
	tlm := app.New("0.0", "test")

	err := tlm.App.Run(args)
	if err != nil {
		os.Exit(1)
	}
}

func Test_Version(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "version")
	run(args)
}

func Test_Help(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "help")
	run(args)
}
