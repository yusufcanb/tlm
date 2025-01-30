package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/Masterminds/semver"
)

// GithubAPIAccessErr error that occurs when there is a problem accessing the GitHub API.
var GithubAPIAccessErr = errors.New("error accessing GitHub API")

// ReleaseManager manages the releases of tlm's GitHub repository.
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

// GetLatest fetches the latest release from the GitHub API.
func (rm *ReleaseManager) GetLatest() (*Release, error) {
	var err error
	var resp *http.Response
	var latestRelease Release

	resp, err = rm.httpClient.Get(fmt.Sprintf("%s/repos/%s/%s/releases/latest", rm.githubApiUrl, rm.owner, rm.repo))
	if err != nil {
		return nil, GithubAPIAccessErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, GithubAPIAccessErr
	}

	if err = json.NewDecoder(resp.Body).Decode(&latestRelease); err != nil {
		return nil, err
	}

	return &latestRelease, nil
}

// CanUpgrade checks if an upgrade is possible from a base version to a target version.
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

// CheckForUpdates checks for updates and writes a checkpoint.
func (rm *ReleaseManager) CheckForUpdates(base string) error {
	var err error
	var latest *Release

	latest, err = rm.GetLatest()
	if err != nil {
		return fmt.Errorf("error fetching latest release: %s", err)
	}

	_, err = rm.CanUpgrade(base, latest)
	if err != nil {
		return fmt.Errorf("error checking for upgrade: %s", err)
	}

	return nil
}

// NewReleaseManager is a constructor for the ReleaseManager struct.
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
