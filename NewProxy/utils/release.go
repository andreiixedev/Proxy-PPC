package utils

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

type ReleaseInfo struct {
    ReleaseTS      float64 `json:"release_ts"`
    ReleaseVersion string  `json:"release_version"`
}

var LatestRelease = ReleaseInfo{}

func FetchLatestRelease() {
    file, err := os.Open("data/release.json")
    if err != nil {
        fmt.Println("Error opening release.json:", err)
        setDefaultRelease()
        return
    }
    defer file.Close()

    var data map[string]string
    if err := json.NewDecoder(file).Decode(&data); err != nil {
        fmt.Println("Error decoding release.json:", err)
        setDefaultRelease()
        return
    }

    releaseVersion := data["release_version"]
    releaseTimeStr := data["release_time"]

    releaseTS := float64(time.Now().Unix())
    if t, err := time.Parse(time.RFC3339, releaseTimeStr); err == nil {
        releaseTS = float64(t.Unix())
    }

    LatestRelease = ReleaseInfo{
        ReleaseTS:      releaseTS,
        ReleaseVersion: releaseVersion,
    }
}

func setDefaultRelease() {
    LatestRelease = ReleaseInfo{
        ReleaseTS:      float64(time.Now().Unix()),
        ReleaseVersion: "N/A",
    }
}
