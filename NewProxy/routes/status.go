package routes

import (
	"github.com/gin-gonic/gin"
	"newproxy/utils"
	"net/http"
	"sync"
	"time"
	"fmt"
)

var (
	activeDownloads = make(map[string]float64)
	totalDownloadedBytes int64
	mu sync.Mutex
)

func RegisterStatusRoute(r *gin.Engine) {
	r.GET("/status", statusHandler)
}

func statusHandler(c *gin.Context) {
    mu.Lock()
    defer mu.Unlock()

    timeout := 10.0
    currentTime := float64(time.Now().Unix())
    filtered := make(map[string]float64)

    for ip, ts := range activeDownloads {
        if currentTime-ts < timeout {
            filtered[ip] = ts
        }
    }
    activeDownloads = filtered

    var totalDownloaded string
    if totalDownloadedBytes >= 1_000_000_000 {
        totalDownloaded = fmt.Sprintf("%.2f GB", float64(totalDownloadedBytes)/1_000_000_000)
    } else {
        totalDownloaded = fmt.Sprintf("%.2f MB", float64(totalDownloadedBytes)/1_000_000)
    }

    c.JSON(http.StatusOK, gin.H{
        "latest_release_version": utils.LatestRelease.ReleaseVersion,
        "proxy_version":          "1.6.2",
        "active_downloads":       len(activeDownloads),
        "total_downloaded":       totalDownloaded,
    })
}