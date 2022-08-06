# GitHub Release Asset Downloader
CLI app to download assets from GitHub Releases. Can choose which version (_using tag_) and which arch to choose. Only
supports Linux. Thanks to [Viper](https://github.com/spf13/viper) you can use both arguments and config file but putting
`GitHub Access Token` in the config file is recommended instead putting it in the arguments.

## How to Use
1. Download the binary from [GitHub Releases](https://github.com/mdanialr/github-asset-downloader/releases)
2. Extract then run
```bash
tar -xvf github-asset-downloader....tar.gz
./github-asset-downloader --help # to see options
```
### Example
Put `token` in config file with the name `app.yml`. __The file name `app.yml`__ is mandatory otherwise
[Viper](https://github.com/spf13/viper) will not find it.
```yaml
# config file app.yml

token: <YOUR_TOKEN>
```
```bash
# download asset linux amd64 from the latest release
./github-asset-downloader --repo username/repo-name
# download asset linux amd64 from the release tagged as v1.0.0
./github-asset-downloader --repo username/repo-name --tag v1.0.0
# download asset linux amd64 from the release tagged as v1.0.0 and output filename as my-asset
./github-asset-downloader --repo username/repo-name --tag v1.0.0 --output my-asset
# NOTE: file extensions will always be extracted from the asset's real name
```
