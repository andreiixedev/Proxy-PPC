package routes

import (
    "github.com/gin-gonic/gin"
    "io"
    "net/http"
    "time"
) //btw same shitty import gin

func RegisterForwardRoute(r *gin.Engine) {
    r.GET("/", forwardHandler)
} // main proxy url

func forwardHandler(c *gin.Context) {
    targetURL := c.Query("url")
    if targetURL == "" {
        c.String(http.StatusBadRequest, "Unsupported request")
        return
    }

    ip := c.ClientIP()
    mu.Lock()
    activeDownloads[ip] = float64(time.Now().Unix())
    mu.Unlock()

    resp, err := http.Get(targetURL)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error fetching the requested URL: %v", err)
        return
    }
    defer resp.Body.Close()

    content, err := io.ReadAll(resp.Body)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error reading response: %v", err)
        return
    }

    mu.Lock()
    totalDownloadedBytes += int64(len(content))
    mu.Unlock()

    c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), content)
}