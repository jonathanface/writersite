package models

type Book struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Link        string `json:"link"`
	ReleasedOn  int    `json:"released_on" dynamodbav:"released_on"`
}
