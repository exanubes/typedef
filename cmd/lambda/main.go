package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/exanubes/typedef/internal/domain"
	driver "github.com/exanubes/typedef/internal/drivers/lambda"
)

var headers = map[string]string{
	"Content-Type":           "application/json; charset=utf-8",
	"X-Content-Type-Options": "nosniff",
	"Referrer-Policy":        "noreferrer",
}

func handle_request(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	response := events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers:    headers,
	}

	var payload domain.CodegenRequest

	bytes := []byte(request.Body)

	if request.IsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			response.StatusCode = 400
			response.Body = fmt.Sprintf(`{"error": %q}`, fmt.Errorf("Failed to decode payload:\n| %w", err).Error())
			return response, nil
		}

		bytes = decoded
	}

	if err := json.Unmarshal(bytes, &payload); err != nil {
		response.StatusCode = 400
		response.Body = fmt.Sprintf(`{"error": %q}`, fmt.Errorf("Failed to unmarshal payload:\n| %w", err).Error())
		return response, nil
	}

	output, err := driver.Start(ctx, payload)

	if err != nil {
		response.StatusCode = 400
		response.Body = fmt.Sprintf(`{"error": %q}`, fmt.Errorf("Failed to run code generation:\n| %w", err).Error())
		return response, nil
	}

	response_bytes, err := json.Marshal(output)

	if err != nil {
		response.StatusCode = 400
		response.Body = fmt.Sprintf(`{"error": %q}`, fmt.Errorf("Failed to marshal response object:\n| %w", err).Error())
		return response, nil
	}

	response.Body = string(response_bytes)

	return response, nil
}

func main() {
	lambda.Start(handle_request)

}
