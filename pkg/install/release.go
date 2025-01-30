package install

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ReleaseAssetNotFoundErr = errors.New("error reaching release artifacts")

type ReleaseAsset struct {
	Url                string `json:"url"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type Release struct {
	Name       string         `json:"name"`
	TagName    string         `json:"tag_name"`
	Draft      bool           `json:"draft"`
	PreRelease bool           `json:"prerelease"`
	Assets     []ReleaseAsset `json:"assets"`
	CreatedAt  time.Time      `json:"created_at"`
}

func (r *Release) GetDownloadUrlFor(platform, arch string) (string, error) {
	for _, asset := range r.Assets {
		if strings.Contains(asset.BrowserDownloadUrl, fmt.Sprintf("%s_%s", platform, arch)) {
			return asset.BrowserDownloadUrl, nil
		}
	}

	return "", ReleaseAssetNotFoundErr
}

func (r *Release) String() string {
	return r.Name
}
