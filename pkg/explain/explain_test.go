package explain_test

import (
	"testing"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/pkg/config"
	"github.com/yusufcanb/tlm/pkg/explain"
)

func TestExplain(t *testing.T) {

	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	explain.New(o, "")
}
