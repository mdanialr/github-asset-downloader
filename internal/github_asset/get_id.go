package github_asset

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var knownExtensions = []string{"tar.gz", "tgz", "gz", "zip"}

// GetAssetID get the targeted asset's id.
func GetAssetID(cl *Client) (*Asset, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases", cl.Repo)
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", cl.Token))
	req.Header.Add("Accept", "application/vnd.github.v3.raw")

	resp, err := cl.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	by, _ := io.ReadAll(resp.Body)
	log.Println("GitHub API status code:", resp.StatusCode)

	var payload []release
	json.Unmarshal(by, &payload)

	if len(payload) > 0 {
		if cl.Tag == "latest" {
			for _, asset := range payload[0].Assets {
				if strings.Contains(asset.Name, fmt.Sprintf("linux_%s", cl.Arch)) {
					for _, ext := range knownExtensions {
						if strings.HasSuffix(asset.Name, fmt.Sprintf(".%s", ext)) {
							asset.FileExt = fmt.Sprintf(".%s", ext)
							asset.FileName = strings.TrimSuffix(asset.Name, asset.FileExt)

							return asset, nil
						}
					}

					asset.FileExt = filepath.Ext(asset.Name)
					asset.FileName = strings.TrimSuffix(asset.Name, asset.FileExt)

					return asset, nil
				}
			}
		}

		for _, pay := range payload {
			if pay.TagName == cl.Tag || pay.Name == cl.Tag {

				for _, asset := range pay.Assets {
					if strings.Contains(asset.Name, fmt.Sprintf("linux_%s", cl.Arch)) {
						for _, ext := range knownExtensions {
							if strings.HasSuffix(asset.Name, fmt.Sprintf(".%s", ext)) {
								asset.FileExt = fmt.Sprintf(".%s", ext)
								asset.FileName = strings.TrimSuffix(asset.Name, asset.FileExt)

								return asset, nil
							}
						}

						asset.FileExt = filepath.Ext(asset.Name)
						asset.FileName = strings.TrimSuffix(asset.Name, asset.FileExt)

						return asset, nil
					}
				}

			}
		}

		return nil, fmt.Errorf("tag %s with arch %s not found", cl.Tag, cl.Arch)
	}

	return nil, nil
}
