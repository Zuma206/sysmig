package updates

import (
	"io"
	"net/http"

	"github.com/zuma206/sysmig/utils"
)

type GithubAsset struct {
	Name               string `json:"name"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

func (asset *GithubAsset) Download() *[]byte {
	resp, err := http.Get(asset.BrowserDownloadUrl)
	utils.HandleErr(err)
	data, err := io.ReadAll(resp.Body)
	utils.HandleErr(err)
	return &data
}
