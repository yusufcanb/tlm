package app_test

import (
	"fmt"
	"github.com/yusufcanb/tlm/app"
	"io"
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

	stdout := os.Stdout

	args := os.Args[0:1]
	args = append(args, " version")

	capturedOutput := os.NewFile(0, "tlm.log")
	os.Stdout = capturedOutput

	run(args)

	// Read all contents from the capturedOutput
	contents, err := io.ReadAll(capturedOutput)
	if err != nil {
		fmt.Println("Error reading contents:", err)
		t.Fail()
	}

	os.Stdout = stdout
	_ = capturedOutput.Close()

	// Print the captured output
	t.Log("Captured Output:")
	t.Log(string(contents))

}

func Test_Help(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "help")
	run(args)
}
