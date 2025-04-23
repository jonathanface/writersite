package models

type News struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	File     string `json:"file"`
	PostedOn int    `json:"posted_on" dynamodbav:"posted_on"`
	EditedOn int    `json:"edited_on" dynamodbav:"edited_on"`
}
