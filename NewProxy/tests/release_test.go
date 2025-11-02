package tests

import (
	"newproxy/utils"
	"os"
	"testing"
)

func TestFetchLatestRelease_ValidFile(t *testing.T) {
	content := `{"release_version": "v1.0.0", "release_time": "2025-01-01T00:00:00Z"}`
	_ = os.WriteFile("data/release.json", []byte(content), 0644)
	utils.FetchLatestRelease()

	if utils.LatestRelease.ReleaseVersion != "v1.0.0" {
		t.Errorf("Expected version v1.0.0, got %s", utils.LatestRelease.ReleaseVersion)
	}
}

func TestFetchLatestRelease_MissingFile(t *testing.T) {
	_ = os.Remove("data/release.json")
	utils.FetchLatestRelease()

	if utils.LatestRelease.ReleaseVersion == "" {
		t.Errorf("Expected default release info, got empty version")
	}
}
