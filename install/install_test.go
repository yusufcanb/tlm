package install_test

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/config"
	"github.com/yusufcanb/tlm/install"
	"testing"
)

func TestInstall(t *testing.T) {

	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	install.New(o, "", "")
}
