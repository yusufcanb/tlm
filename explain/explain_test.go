package explain_test

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/config"
	"github.com/yusufcanb/tlm/explain"
	"testing"
)

func TestExplain(t *testing.T) {

	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	explain.New(o, "")
}
