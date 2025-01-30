package suggest_test

import (
	"testing"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/pkg/config"
	"github.com/yusufcanb/tlm/pkg/suggest"
)

func TestSuggest(t *testing.T) {
	o, _ := ollama.ClientFromEnvironment()

	con := config.New(o)
	con.LoadOrCreateConfig()

	suggest.New(o, "")
}
