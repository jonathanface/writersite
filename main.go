package main

import (
	"jonathanface/api"
	"jonathanface/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	uiDirectory = "ui"
	apiPath     = "/api"
)

// main.go (add this near the top, before main)

func buildRouter(newsHandler gin.HandlerFunc, booksHandler gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	r.GET(apiPath+"/news", newsHandler)
	r.GET(apiPath+"/books", booksHandler)

	// static UI + SPA fallback
	r.Use(static.Serve("/", static.LocalFile(uiDirectory, false)))
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(uiDirectory, "index.html"))
	})
	return r
}

func main() {
	currentMode := models.AppMode(strings.ToLower(os.Getenv("MODE")))
	if currentMode != models.ModeProduction && currentMode != models.ModeStaging {
		if err := godotenv.Load(); err != nil {
			log.Println("warning: .env not loaded:", err)
		}
	}

	r := buildRouter(api.GetNews, api.GetBooks)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8443"
	}
	addr := port
	if !strings.HasPrefix(port, ":") {
		addr = ":" + port
	}

	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
