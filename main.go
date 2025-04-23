package main

import (
	"jonathanface/api"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	uiDirectory = "./ui/public"
	port        = ":8443"
	apiPath     = "/api"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.SetTrustedProxies([]string{"192.168.0.1"})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	r.Use(static.Serve("/", static.LocalFile(uiDirectory, false)))
	r.GET(apiPath+"/news", api.GetNews)
	r.GET(apiPath+"/books", api.GetBooks)

	r.NoRoute(func(c *gin.Context) {
		c.File(uiDirectory + "/index.html")
	})
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
