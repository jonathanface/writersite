package main

import (
	"jonathanface/api"
	"jonathanface/models"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	uiDirectory = "/ui/"
	apiPath     = "/api"
)

func main() {
	currentMode := models.AppMode(strings.ToLower(os.Getenv("MODE")))
	if currentMode != models.ModeProduction && currentMode != models.ModeStaging {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	port := os.Getenv("PORT")
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
