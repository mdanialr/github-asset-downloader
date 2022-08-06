package github_asset

import "net/http"

type Client struct {
	HttpClient *http.Client
	Repo       string
	Token      string
	Arch       string
	Tag        string
}

// Asset inner payload from GitHub api that contain all assets' id.
type Asset struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	FileName string `json:"-"`
	FileExt  string `json:"-"`
}

// release outer payload from GitHub api that contain all assets.
type release struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	TagName string   `json:"tag_name"`
	Assets  []*Asset `json:"assets"`
}
