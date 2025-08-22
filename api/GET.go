package api

import (
	"context"
	"jonathanface/models"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func newDynamoClient() (*dynamodb.Client, error) {
	// Load .env for non-prod/staging
	currentMode := models.AppMode(strings.ToLower(os.Getenv("MODE")))
	if currentMode != models.ModeProduction && currentMode != models.ModeStaging {
		_ = godotenv.Load()
	}

	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1"
	}

	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	awsCfg.RetryMaxAttempts = 5

	return dynamodb.NewFromConfig(awsCfg), nil
}

func GetNews(c *gin.Context) {
	ctx := c.Request.Context()

	ddb, err := newDynamoClient()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	table := os.Getenv("NEWS_TABLE")
	if table == "" {
		table = "jonathanface_news"
	}

	out, err := ddb.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(table),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: "news"},
		},
		ScanIndexForward: aws.Bool(false), // newest first if you have a sort key
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	var newsEntries []models.News
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &newsEntries); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newsEntries)
}

func GetBooks(c *gin.Context) {
	ctx := c.Request.Context()

	ddb, err := newDynamoClient()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	table := os.Getenv("BOOKS_TABLE")
	if table == "" {
		table = "jonathanface_books2"
	}

	out, err := ddb.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(table),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: "books"},
		},
		ScanIndexForward: aws.Bool(false),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	var bookEntries []models.Book
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &bookEntries); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookEntries)
}
