package suggest_test

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/config"
	"github.com/yusufcanb/tlm/suggest"
	"testing"
)

func TestSuggest(t *testing.T) {
	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	suggest.New(o)
}
