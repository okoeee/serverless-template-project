package lambda

import (
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {

	os.Setenv("APP_ENV", "dev")

	request := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/hello",
	}

	response, err := HandleRequest(request)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", response.StatusCode)
	}

}
