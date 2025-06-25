package updates

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/zuma206/sysmig/utils"
)

type GithubRelease struct {
	TagName   string         `json:"tag_name"`
	CreatedAt string         `json:"created_at"`
	Assets    []*GithubAsset `json:"assets"`
}

func (release *GithubRelease) GetCreatedAt() *time.Time {
	createdAt, err := time.Parse(time.RFC3339, release.CreatedAt)
	utils.HandleErr(err)
	return &createdAt
}

func (release *GithubRelease) GetAsset(name string) *GithubAsset {
	var result *GithubAsset
	for _, asset := range release.Assets {
		if asset.Name == name {
			result = asset
		}
	}
	return result
}

type GithubReleases []*GithubRelease

func (releases *GithubReleases) GetLatestRelease() *GithubRelease {
	var latestRelease *GithubRelease = nil
	var latestReleaseCreatedAt *time.Time = nil
	for _, release := range *releases {
		releaseCreatedAt := release.GetCreatedAt()
		if latestRelease == nil || releaseCreatedAt.After(*latestReleaseCreatedAt) {
			latestReleaseCreatedAt = releaseCreatedAt
			latestRelease = release
		}
	}
	return latestRelease
}

const releasesUrl = "https://api.github.com/repos/Zuma206/sysmig/releases"

func GetReleases() *GithubReleases {
	resp, err := http.Get(releasesUrl)
	utils.HandleErr(err)
	var releases GithubReleases
	err = json.NewDecoder(resp.Body).Decode(&releases)
	utils.HandleErr(err)
	return &releases
}
