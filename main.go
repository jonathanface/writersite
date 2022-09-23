package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"jonathanface/api"
)

const (
	uiDirectory = "./ui/public"
	port        = ":4001"
	apiPath     = "/api"
)

func main() {
	godotenv.Load("configs.env")
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.SetTrustedProxies([]string{"192.168.0.1"})
	r.Use(static.Serve("/", static.LocalFile(uiDirectory, false)))
	r.GET(apiPath+"/about", api.GetAbout)
	r.GET(apiPath+"/news", api.GetNews)
	r.GET(apiPath+"/books", api.GetBooks)
	r.GET(apiPath+"/contact", api.GetContact)
	r.NoRoute(func(c *gin.Context) {
		c.File(uiDirectory + "/index.html")
	})
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
