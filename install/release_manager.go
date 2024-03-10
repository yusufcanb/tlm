package install

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/yusufcanb/tlm/shell"
	"net/http"
	"runtime"
	"time"
)

var GithubAPIAccessError = errors.New("error accessing GitHub API")

type ReleaseManager struct {
	httpClient *http.Client

	githubApiUrl string
	repo         string
	owner        string

	platform       string
	arch           string
	deploymentPath string

	Releases []Release
	Message  string
}

func (rm *ReleaseManager) GetLatest() (*Release, error) {
	var err error
	var resp *http.Response
	var latestRelease Release

	resp, err = rm.httpClient.Get(fmt.Sprintf("%s/repos/%s/%s/releases/latest", rm.githubApiUrl, rm.owner, rm.repo))
	if err != nil {
		return nil, GithubAPIAccessError
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, GithubAPIAccessError
	}

	if err = json.NewDecoder(resp.Body).Decode(&latestRelease); err != nil {
		return nil, err
	}

	return &latestRelease, nil
}

func (rm *ReleaseManager) CanUpgrade(base string, to *Release) (bool, error) {
	var err error

	if to.Draft || to.PreRelease {
		return false, nil
	}

	var baseVersion, toVersion *semver.Version
	baseVersion, err = semver.NewVersion(base)
	if err != nil {
		return false, fmt.Errorf("invalid base version: %s", base)
	}

	toVersion, err = semver.NewVersion(to.Name)
	if err != nil {
		return false, fmt.Errorf("invalid to version: %s", to.Name)
	}

	if baseVersion.Major() != toVersion.Major() {
		return false, nil
	}

	return toVersion.GreaterThan(baseVersion), nil
}

func (rm *ReleaseManager) CheckForUpdates(base string) error {
	var cp *shell.Checkpoint
	var err error

	now := time.Now()

	cp, err = shell.GetCheckpoint()
	if errors.Is(err, shell.CheckpointFileNotExistErr) {
		cp = &shell.Checkpoint{
			Message:        "",
			LastCheckpoint: now,
		}
		_ = shell.WriteCheckpoint(cp)
	}

	if now == cp.LastCheckpoint || now.Sub(cp.LastCheckpoint) > 3*time.Hour {
		latest, err := rm.GetLatest()
		if err != nil {
			cp.Message = ""
			cp.LastCheckpoint = time.Now()
			_ = shell.WriteCheckpoint(cp)
			return nil
		}

		yes, err := rm.CanUpgrade(base, latest)
		if err != nil {
			cp.Message = ""
			cp.LastCheckpoint = time.Now()
			_ = shell.WriteCheckpoint(cp)
			return nil
		}

		if yes {
			cp.Message = fmt.Sprintf("A new version of tlm is available (%s)\nPlease run the installation script to get the latest version.", latest.Name)

		} else {
			cp.Message = ""
		}

		cp.LastCheckpoint = time.Now()
		_ = shell.WriteCheckpoint(cp)
	}

	if cp.Message != "" {
		fmt.Println(cp.Message)
	}

	return nil
}

func NewReleaseManager(owner, repo string) *ReleaseManager {
	rm := &ReleaseManager{}
	rm.httpClient = &http.Client{}

	rm.githubApiUrl = "https://api.github.com"
	rm.owner = owner
	rm.repo = repo

	rm.platform = runtime.GOOS
	rm.arch = runtime.GOARCH

	return rm
}
