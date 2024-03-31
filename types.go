package main

type Ingredient struct {
	Id      string
	Content string
}

type Step struct {
	Id         string
	Content    string
	StepNumber int
}

type Entity struct {
	Id          string       `json:"id" dynamodbav:"id"`
	AuthorId    string       `json:"authorId" dynamodbav:"authorId"`
	Title       string       `json:"title" dynamodbav:"title"`
	Description string       `json:"description" dynamodbav:"description"`
	FoodType    string       `json:"foodType" dynamodbav:"foodType"`
	Cuisine     string       `json:"cuisine" dynamodbav:"cuisine"`
	Ingredients []Ingredient `json:"ingredients" dynamodbav:"ingredients"`
	Steps       []Step       `json:"steps" dynamodbav:"steps"`
	LikedBy     []string     `json:"likedBy" dynamodbav:"likedBy"`
}

type NewOrUpdatedEntity struct {
	AuthorId    string       `json:"authorId" validate:"required"`
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"description" validate:"required"`
	FoodType    string       `json:"foodType" validate:"required"`
	Cuisine     string       `json:"cuisine" validate:"required"`
	Ingredients []Ingredient `json:"ingredients" validate:"required"`
	Steps       []Step       `json:"steps" validate:"required"`
	LikedBy     []string     `json:"likedBy" validate:"required"`
}
