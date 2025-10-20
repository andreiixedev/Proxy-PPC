package routes

import (
    "github.com/gin-gonic/gin" //hellow gin again ;)
    "newproxy/utils"
    "net/http"
)

func RegisterBuildsRoute(r *gin.Engine) {
    r.GET("/client/builds.json", buildsHandler)
}

func buildsHandler(c *gin.Context) {
    utils.FetchLatestRelease()
    c.JSON(http.StatusOK, utils.LatestRelease)
}
