package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

// Downloads maps public URL slugs to S3 object keys.
// Add entries here to make files available at /downloads/<slug>
var Downloads = map[string]string{
	"thecow": "THE_COW-jface.pdf",
}

func GetDownload(c *gin.Context) {
	slug := strings.TrimPrefix(c.Param("filepath"), "/")
	if slug == "" {
		c.Status(http.StatusNotFound)
		return
	}

	s3Key, ok := Downloads[slug]
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}

	bucket := os.Getenv("DOWNLOADS_BUCKET")
	if bucket == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "downloads not configured"})
		return
	}

	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1"
	}

	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to load AWS config"})
		return
	}

	clientIP := c.ClientIP()
	referrer := c.GetHeader("Referer")
	log.Printf("download: slug=%s ip=%s referrer=%s", slug, clientIP, referrer)

	presigner := s3.NewPresignClient(s3.NewFromConfig(awsCfg))
	req, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &s3Key,
	}, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		log.Printf("presign error: %v (bucket=%s key=%s)", err, bucket, s3Key)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to generate download URL"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, req.URL)
}
