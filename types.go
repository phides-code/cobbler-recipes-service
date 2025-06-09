package main

type ImageSource struct {
	OriginalName string `json:"originalName" dynamodbav:"originalName"`
	UUIDName     string `json:"uuidName" dynamodbav:"uuidName"`
}

type Entity struct {
	Id             string `json:"id,omitempty" dynamodbav:"id"`
	Author         string `json:"author" dynamodbav:"author"`
	Likes          int    `json:"likes" dynamodbav:"likes"`
	CreatedOn      uint64 `json:"createdOn" dynamodbav:"createdOn"`
	LowercaseTitle string `json:"lowercaseTitle" dynamodbav:"lowercaseTitle"`
	TagsList       string `json:"tagsList" dynamodbav:"tagsList"`
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
	Likes *int `json:"likes" validate:"required,gte=0"` // Ensure that 0 likes is allowed
}

// ResponseEntity omits LowercaseTitle and TagsList fields
// as they are not needed in the response.
type ResponseEntity struct {
	Id          string      `json:"id,omitempty"`
	Author      string      `json:"author"`
	Likes       int         `json:"likes"`
	CreatedOn   uint64      `json:"createdOn"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Tags        []string    `json:"tags"`
	Ingredients []string    `json:"ingredients"`
	Steps       []string    `json:"steps"`
	PrepTime    string      `json:"prepTime"`
	ImageSource ImageSource `json:"imageSource"`
}
