package install_test

import (
	"testing"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/pkg/config"
	"github.com/yusufcanb/tlm/pkg/explain"
	"github.com/yusufcanb/tlm/pkg/install"
	"github.com/yusufcanb/tlm/pkg/suggest"
)

const (
	repoOwner = "yusufcanb" // Replace with the owner of the repository
	repoName  = "tlm"       // Replace with the name of the repository
)

func TestInstall(t *testing.T) {

	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	install.New(o, suggest.New(o, ""), explain.New(o, ""))
}

func TestReleaseManager_CanUpgrade(t *testing.T) {

	// upgrade test matrix
	shouldUpgradeMatrix := [][]any{
		{"1.0-rc1", install.Release{Name: "1.0", Draft: false, PreRelease: false}},
		{"1.0-rc2", install.Release{Name: "1.0", Draft: false, PreRelease: false}},
		{"1.0", install.Release{Name: "1.1", Draft: false, PreRelease: false}},
		{"1.0", install.Release{Name: "1.1-rc0", Draft: false, PreRelease: false}},
		{"1.1-rc0", install.Release{Name: "1.1-rc1", Draft: false, PreRelease: false}},
	}

	rm := install.NewReleaseManager(repoOwner, repoName)

	for _, test := range shouldUpgradeMatrix {
		ok, err := rm.CanUpgrade(test[0].(string), test[1].(*install.Release))
		if err != nil {
			t.Fatal(err)
		}
		if !ok {
			t.Fatalf("expected %s to be able to upgrade to %s", test[0], test[1])
		}
		t.Log(test[0], "can upgrade to", test[1].(install.Release).Name)
	}

}

func TestReleaseManager_GetLatest(t *testing.T) {
	rm := install.NewReleaseManager(repoOwner, repoName)
	latest, _ := rm.GetLatest()

	if latest.Name != "1.0" {
		t.Fatal("latest release is not 1.0")
	}

	t.Log("latest release is", latest.Name)
}

func TestReleaseManager_UpgradeTo(t *testing.T) {
}
