package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	request := make(map[string]interface{})
	request["name"] = "nuv"

	response := HelloWorld(request)

	if response["body"] != "Hello, nuv!" {
		t.Errorf("Expected 'Hello, nuv!', got %s", response["body"])
	}
}
