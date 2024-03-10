package install

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/yusufcanb/tlm/shell"
	"net/http"
	"sort"
	"time"
)

type Release struct {
	Name       string    `json:"name"`
	TagName    string    `json:"tag_name"`
	Draft      bool      `json:"draft"`
	PreRelease bool      `json:"prerelease"`
	CreatedAt  time.Time `json:"created_at"`
}

type ReleaseManager struct {
	httpClient *http.Client

	githubApiUrl string
	repo         string
	owner        string

	Releases []Release
	Message  string
}

func (rm *ReleaseManager) getReleases() error {
	var err error
	var resp *http.Response

	url := fmt.Sprintf("%s/repos/%s/%s/releases", rm.githubApiUrl, rm.owner, rm.repo)

	resp, err = rm.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GitHub API request failed with status: %s", resp.Status)
	}

	var releases []Release
	if err = json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return err
	}

	rm.Releases = releases

	return nil
}

func (rm *ReleaseManager) Ping() bool {
	resp, err := http.Get(rm.githubApiUrl)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func (rm *ReleaseManager) IsLatest(version string) (bool, error) {
	for _, release := range rm.Releases {
		if release.TagName == version {
			return true, nil
		}
	}

	return false, nil
}

func (rm *ReleaseManager) GetLatest() (Release, error) {
	err := rm.getReleases()
	if err != nil {
		return Release{}, err
	}

	releases := make([]Release, len(rm.Releases))
	copy(releases, rm.Releases)

	sort.Slice(releases, func(i, j int) bool {
		return releases[i].CreatedAt.After(releases[j].CreatedAt)
	})

	if len(releases) == 0 {
		return Release{}, errors.New("no releases found")
	}

	return releases[0], nil
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
		shell.WriteCheckpoint(cp)
	}

	if now == cp.LastCheckpoint || now.Sub(cp.LastCheckpoint) > 24*time.Hour {
		latest, err := rm.GetLatest()
		if err != nil {
			cp.Message = ""
			cp.LastCheckpoint = time.Now()
			shell.WriteCheckpoint(cp)
			return nil
		}
		yes, err := rm.CanUpgrade(base, &latest)
		if err != nil {
			cp.Message = ""
			cp.LastCheckpoint = time.Now()
			shell.WriteCheckpoint(cp)
			return nil
		}

		if yes {
			cp.Message = "A new version of tlm is available. Run `tlm upgrade` to upgrade."
		} else {
			cp.Message = ""
		}

		cp.LastCheckpoint = time.Now()
		shell.WriteCheckpoint(cp)
	}

	if cp.Message != "" {
		fmt.Println(cp.Message)
	}

	return nil
}

func (rm *ReleaseManager) UpgradeTo(release *Release) error {
	return nil
}

func NewReleaseManager(owner, repo string) *ReleaseManager {
	r := &ReleaseManager{}
	r.httpClient = &http.Client{
		Timeout: 350 * time.Millisecond,
	}

	r.githubApiUrl = "https://api.github.com"
	r.owner = owner
	r.repo = repo

	return r
}
