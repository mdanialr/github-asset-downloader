package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mdanialr/github-asset-downloader/internal/github_asset"
	"github.com/mdanialr/github-asset-downloader/pkg/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	TOKEN  string
	REPO   string
	ARCH   string
	TAG    string
	OUTPUT string
	EXT    string
)

func init() {
	flag.StringVar(&TOKEN, "token", "", "Github access token. Required.")
	flag.StringVar(&REPO, "repo", "", "Repository to download. Example: org/hello-world. Required.")
	flag.StringVar(&ARCH, "arch", "", "Architecture to download. amd64 or arm64. only support Linux")
	flag.StringVar(&TAG, "tag", "", "Tag to download. Example: v1.0.0. Can be omitted to download from the latest tag")
	flag.StringVar(&OUTPUT, "output", "", "Output file name. The default will use the name in the asset's entry")
	flag.StringVar(&EXT, "ext", "", "Output file extension. The default will use the extension of the asset's name")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	conf := config.InitConfig(".")
	if err := config.SetupDefault(conf); err != nil {
		log.Fatalln("there are some errors in config:", err)
	}

	cl := &github_asset.Client{
		HttpClient: http.DefaultClient,
		Token:      conf.GetString("token"),
		Repo:       conf.GetString("repo"),
		Arch:       conf.GetString("arch"),
		Tag:        conf.GetString("tag"),
	}
	foundAsset, err := github_asset.GetAssetID(cl)
	if err != nil {
		log.Fatalln("failed to retrieve asset id:", err)
	}

	if conf.GetString("output") != "" {
		foundAsset.FileName = conf.GetString("output")
	}
	if conf.GetString("ext") != "" {
		foundAsset.FileExt = conf.GetString("ext")
	}

	if err = github_asset.DownloadAsset(cl, foundAsset); err != nil {
		log.Fatalln("failed to download asset:", err)
	}
}
