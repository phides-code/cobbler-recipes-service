package main

type ImageSource struct {
	OriginalName string `json:"originalName" dynamodbav:"originalName"`
	UUIDName     string `json:"uuidName" dynamodbav:"uuidName"`
}

type Entity struct {
	Id          string      `json:"id,omitempty" dynamodbav:"id"`
	Author      string      `json:"author" dynamodbav:"author"`
	Title       string      `json:"title" dynamodbav:"title"`
	Description string      `json:"description" dynamodbav:"description"`
	Tags        []string    `json:"tags" dynamodbav:"tags"`
	Ingredients []string    `json:"ingredients" dynamodbav:"ingredients"`
	Steps       []string    `json:"steps" dynamodbav:"steps"`
	Likes       int         `json:"likes" dynamodbav:"likes"`
	PrepTime    string      `json:"prepTime" dynamodbav:"prepTime"`
	ImageSource ImageSource `json:"imageSource" dynamodbav:"imageSource"`
	CreatedOn   uint64      `json:"createdOn" dynamodbav:"createdOn"`
}

type NewEntity struct {
	Author      string      `json:"author" validate:"required"`
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Tags        []string    `json:"tags" validate:"required"`
	Ingredients []string    `json:"ingredients" validate:"required"`
	Steps       []string    `json:"steps" validate:"required"`
	Likes       int         `json:"likes" validate:"required"`
	PrepTime    string      `json:"prepTime" validate:"required"`
	ImageSource ImageSource `json:"imageSource" validate:"required"`
}

type UpdatedEntity struct {
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Tags        []string    `json:"tags" validate:"required"`
	Ingredients []string    `json:"ingredients" validate:"required"`
	Steps       []string    `json:"steps" validate:"required"`
	Likes       int         `json:"likes" validate:"required"`
	PrepTime    string      `json:"prepTime" validate:"required"`
	ImageSource ImageSource `json:"imageSource" validate:"required"`
}
