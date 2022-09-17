package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.SetTrustedProxies([]string{"192.168.0.1"})
	r.Use(static.Serve("/", static.LocalFile("./ui/public", false)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./ui/public/index.html")
	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
