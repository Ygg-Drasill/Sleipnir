package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
)

func getCurrentVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	return info.Main.Version
}

func getLatestVersion() (string, error) {
	const repoURL = "https://api.github.com/repos/Ygg-Drasill/Sleipnir/releases/latest"
	resp, err := http.Get(repoURL)
	if err != nil {
		return "", fmt.Errorf("error fetching latest version: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error fetching latest version: received status code %d", resp.StatusCode)
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return "", fmt.Errorf("error parsing latest version response: %w", err)
	}

	return release.TagName, nil
}
