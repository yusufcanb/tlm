package shell_test

import (
	"github.com/yusufcanb/tlm/shell"
	"testing"
	"time"
)

func Test_WriteCheckpoint(t *testing.T) {
	var err error

	cp := &shell.Checkpoint{
		Message:        "",
		LastCheckpoint: time.Now(),
	}

	err = shell.WriteCheckpoint(cp)
	if err != nil {
		t.Errorf("Error writing checkpoint: %v", err)
	}

}

func Test_GetCheckpoint(t *testing.T) {
	var err error
	var checkpoint *shell.Checkpoint

	cp := &shell.Checkpoint{
		Message:        "hello",
		LastCheckpoint: time.Now(),
	}
	err = shell.WriteCheckpoint(cp)

	checkpoint, err = shell.GetCheckpoint()
	if err != nil {
		t.Errorf("Error getting checkpoint: %v", err)
	}

	if checkpoint == nil {
		t.Errorf("Checkpoint is nil")
	}

	if checkpoint.Message != "hello" {
		t.Errorf("Expected checkpoint.Latest to be 'hello', got '%s'", checkpoint.Message)
	}

}
