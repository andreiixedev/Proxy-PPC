package main

import (
	"github.com/gin-gonic/gin"
	"newproxy/routes"
)

func main() {
	r := gin.Default()

	// Register routes of the proxy
	routes.RegisterStatusRoute(r)
	routes.RegisterForwardRoute(r)
	routes.RegisterBuildsRoute(r)

	r.Run(":5090") //from old proxy :)
}