package suggest

import (
	"testing"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/config"
)

func TestRefineCommand(t *testing.T) {
	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	s := New(o)

	if s.refineCommand("ls -al") != "ls -al" {
		t.Error("no change should be made if the command is already okay")
	}

	if s.refineCommand("$ ls -al") != "ls -al" {
		t.Error("shell prefix should be removed")
	}

	if s.refineCommand("❯ ls -al") != "ls -al" {
		t.Error("shell prefix should be removed")
	}

	if s.refineCommand(" ls -al") != "ls -al" {
		t.Error("leading space should be removed")
	}
}
