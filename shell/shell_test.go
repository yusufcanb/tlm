package shell_test

import (
	"github.com/yusufcanb/tlm/shell"
	"testing"
)

func Test_Shell(t *testing.T) {

	t.Run("shell.Err()", func(t *testing.T) {
		if shell.Err() == "(err)" {
			t.Log("Shell Err() is working")
			return
		}
		t.Fail()
	})

	t.Run("shell.Ok()", func(t *testing.T) {
		if shell.Ok() == "(ok)" {
			t.Log("Shell Ok() is working")
			return
		}
		t.Fail()
	})

}
