package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	uiDirectory = "./ui/public"
	port        = ":8080"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.SetTrustedProxies([]string{"192.168.0.1"})
	r.Use(static.Serve("/", static.LocalFile(uiDirectory, false)))
	r.NoRoute(func(c *gin.Context) {
		c.File(uiDirectory + "/index.html")
	})
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
