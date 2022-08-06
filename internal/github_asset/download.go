package github_asset

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadAsset downloads the given asset.
func DownloadAsset(cl *Client, asset *Asset) error {
	url := fmt.Sprintf("https://%s:@api.github.com/repos/%s/releases/assets/%d", cl.Token, cl.Repo, asset.Id)
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/octet-stream")

	resp, err := cl.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download asset: %s", err)
	}
	defer resp.Body.Close()
	log.Println("Response status:", resp.Status)

	fl, err := os.Create(fmt.Sprintf("%s%s", asset.FileName, asset.FileExt))
	if err != nil {
		return fmt.Errorf("failed to create file: %s", err)
	}

	io.Copy(fl, resp.Body)
	return nil
}
