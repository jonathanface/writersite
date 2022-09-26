package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func dbConnect() (*sql.DB, error) {
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbHost := os.Getenv("dbHost")
	dbName := os.Getenv("dbName")
	return sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?parseTime=true")
}

func GetAbout(c *gin.Context) {
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
	if db, err = dbConnect(); err != nil {
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

func GetNews(c *gin.Context) {
	type News struct {
		ID       int       `json:"id" db:"id"`
		Title    string    `json:"title" db:"title"`
		Post     string    `json:"post" db:"post"`
		PostedOn time.Time `json:"posted_on" db:"posted_on"`
		EditedOn time.Time `json:"edited_on" db:"edited_on"`
	}
	var (
		db          *sql.DB
		err         error
		newsEntries []News
		rows        *sql.Rows
	)

	if db, err = dbConnect(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	if rows, err = db.Query("SELECT * FROM news ORDER BY posted_on DESC LIMIT 10"); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		news := News{}
		if err = rows.Scan(&news.Title, &news.Post, &news.PostedOn, &news.EditedOn, &news.ID); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}
		newsEntries = append(newsEntries, news)
	}
	c.JSON(http.StatusOK, newsEntries)
}
func GetContact(c *gin.Context) {
	type Contact struct {
		ID        int       `json:"id" db:"id"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
		Body      string    `json:"body" db:"body"`
	}
	var (
		db      *sql.DB
		err     error
		contact Contact
	)
	if db, err = dbConnect(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	if err = db.QueryRow("SELECT id, updated_at, body FROM contact ORDER BY updated_at ASC LIMIT 1").Scan(&contact.ID, &contact.UpdatedAt, &contact.Body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}
func GetBooks(c *gin.Context) {
	type Book struct {
		ID          int       `json:"id" db:"id"`
		Title       string    `json:"title" db:"title"`
		Description string    `json:"description" db:"description"`
		Img         string    `json:"img" db:"img"`
		Link        string    `json:"link" db:"link"`
		ReleasedOn  time.Time `json:"released_on" db:"released_on"`
	}
	var (
		db    *sql.DB
		err   error
		books []Book
		rows  *sql.Rows
	)
	if db, err = dbConnect(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	if err = db.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	if rows, err = db.Query("SELECT * FROM books ORDER BY released_on DESC"); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		book := Book{}
		if err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Img, &book.Link, &book.ReleasedOn); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}
