package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func clientError(status int) (events.APIGatewayProxyResponse, error) {

	errorString := http.StatusText(status)

	response := ResponseStructure{
		Data:         nil,
		ErrorMessage: &errorString,
	}

	responseJson, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Body:       string(responseJson),
		StatusCode: status,
		Headers:    headers,
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())

	errorString := http.StatusText(http.StatusInternalServerError)

	response := ResponseStructure{
		Data:         nil,
		ErrorMessage: &errorString,
	}

	responseJson, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Body:       string(responseJson),
		StatusCode: http.StatusInternalServerError,
		Headers:    headers,
	}, nil
}

func mergeHeaders(baseHeaders, additionalHeaders map[string]string) map[string]string {
	mergedHeaders := make(map[string]string)
	for key, value := range baseHeaders {
		mergedHeaders[key] = value
	}
	for key, value := range additionalHeaders {
		mergedHeaders[key] = value
	}
	return mergedHeaders
}

func toResponseEntity(e *Entity) *ResponseEntity {
	if e == nil {
		return nil
	}
	return &ResponseEntity{
		Id:          e.Id,
		Author:      e.Author,
		Likes:       e.Likes,
		CreatedOn:   e.CreatedOn,
		Title:       e.Title,
		Description: e.Description,
		Tags:        e.Tags,
		Ingredients: e.Ingredients,
		Steps:       e.Steps,
		PrepTime:    e.PrepTime,
		ImageSource: e.ImageSource,
	}
}

func toResponseEntitySlice(entities []Entity) []ResponseEntity {
	responses := make([]ResponseEntity, len(entities))
	for i, e := range entities {
		responses[i] = *toResponseEntity(&e)
	}
	return responses
}

func flattenArray(stringArray []string) string {
	var result string

	for _, str := range stringArray {
		result += str + ","
	}

	return result
}

func sortEntitiesByLikes(entities []Entity) []Entity {
	if len(entities) == 0 {
		return entities
	}

	sortedEntities := make([]Entity, len(entities))
	copy(sortedEntities, entities)

	for i := 0; i < len(sortedEntities)-1; i++ {
		for j := 0; j < len(sortedEntities)-i-1; j++ {
			if sortedEntities[j].Likes < sortedEntities[j+1].Likes {
				sortedEntities[j], sortedEntities[j+1] = sortedEntities[j+1], sortedEntities[j]
			}
		}
	}

	return sortedEntities
}
