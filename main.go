package main

import (
	"database/sql"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

const (
	uiDirectory = "./ui/public"
	port        = ":4001"
	apiPath     = "/api"
)

func getAbout(c *gin.Context) {
	type About struct {
		ID        int       `json:"id" db:"id"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
		Body      string    `json:"body" db:"body"`
	}
	var (
		db    *sql.DB
		err   error
		about About
	)
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbHost := os.Getenv("dbHost")
	dbName := os.Getenv("dbName")
	if db, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?parseTime=true"); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	if err = db.QueryRow("SELECT id, updated_at, body FROM about ORDER BY updated_at ASC LIMIT 1").Scan(&about.ID, &about.UpdatedAt, &about.Body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, about)
}

func getNews(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Page not found."})
}
func getContact(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Page not found."})
}
func getBooks(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Page not found."})
}

func main() {
	godotenv.Load("configs.env")
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.SetTrustedProxies([]string{"192.168.0.1"})
	r.Use(static.Serve("/", static.LocalFile(uiDirectory, false)))
	r.GET(apiPath+"/about", getAbout)
	r.GET(apiPath+"/news", getNews)
	r.GET(apiPath+"/books", getBooks)
	r.GET(apiPath+"/contact", getContact)
	r.NoRoute(func(c *gin.Context) {
		c.File(uiDirectory + "/index.html")
	})
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
