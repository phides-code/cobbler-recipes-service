package main

type ImageSource struct {
	OriginalName string `json:"originalName" dynamodbav:"originalName" validate:"required"`
	UUIDName     string `json:"uuidName" dynamodbav:"uuidName" validate:"required"`
}

type Entity struct {
	Id        string `json:"id,omitempty" dynamodbav:"id"`
	Author    string `json:"author" dynamodbav:"author"`
	Likes     int    `json:"likes" dynamodbav:"likes"`
	CreatedOn uint64 `json:"createdOn" dynamodbav:"createdOn"`
	//
	Title       string      `json:"title" dynamodbav:"title"`
	Description string      `json:"description" dynamodbav:"description"`
	Tags        []string    `json:"tags" dynamodbav:"tags"`
	Ingredients []string    `json:"ingredients" dynamodbav:"ingredients"`
	Steps       []string    `json:"steps" dynamodbav:"steps"`
	PrepTime    string      `json:"prepTime" dynamodbav:"prepTime"`
	ImageSource ImageSource `json:"imageSource" dynamodbav:"imageSource"`
}

type NewEntity struct {
	Author string `json:"author" validate:"required"`
	//
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Tags        []string    `json:"tags" validate:"required"`
	Ingredients []string    `json:"ingredients" validate:"required"`
	Steps       []string    `json:"steps" validate:"required"`
	PrepTime    string      `json:"prepTime" validate:"required"`
	ImageSource ImageSource `json:"imageSource" validate:"required"`
}

type UpdatedEntity struct {
	Likes int `json:"likes" validate:"required"`
	//
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Tags        []string    `json:"tags" validate:"required"`
	Ingredients []string    `json:"ingredients" validate:"required"`
	Steps       []string    `json:"steps" validate:"required"`
	PrepTime    string      `json:"prepTime" validate:"required"`
	ImageSource ImageSource `json:"imageSource" validate:"required"`
}
