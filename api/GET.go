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

const templateDir = "ui/public/html/"

type htmlTemplate struct {
	Body string `json:"body"`
}

func dbClient() (*dynamodb.Client, error) {
	var err error
	currentMode := models.AppMode(strings.ToLower(os.Getenv("MODE")))
	if currentMode != models.ModeProduction && currentMode != models.ModeStaging {
		if err = godotenv.Load(); err != nil {
			return nil, err
		}
	}
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		opts.Region = os.Getenv("AWS_REGION")
		return nil
	})
	if err != nil {
		return nil, err
	}
	awsCfg.RetryMaxAttempts = 5
	return dynamodb.NewFromConfig(awsCfg), nil
}

func GetNews(c *gin.Context) {
	var (
		err         error
		newsEntries []models.News
	)

	dbClient, err := dbClient()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	out, err := dbClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("jonathanface_news"),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: "news"},
		},
		ScanIndexForward: aws.Bool(false),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
	newsEntries = []models.News{}
	if err = attributevalue.UnmarshalListOfMaps(out.Items, &newsEntries); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
	c.JSON(http.StatusOK, newsEntries)
	return
}

func GetBooks(c *gin.Context) {
	var (
		err         error
		bookEntries []models.Book
	)

	dbClient, err := dbClient()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	out, err := dbClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("jonathanface_books2"),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: "books"},
		},
		ScanIndexForward: aws.Bool(false),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
	bookEntries = []models.Book{}
	if err = attributevalue.UnmarshalListOfMaps(out.Items, &bookEntries); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
	c.JSON(http.StatusOK, bookEntries)
	return
}
