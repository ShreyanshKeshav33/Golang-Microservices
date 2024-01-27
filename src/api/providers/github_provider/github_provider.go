package githubprovider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/clients/restclient"
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/domain/github"
)

// Constants for headers and API URL
const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

// Helper function to get the Authorization header
func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// CreateRepo sends a request to GitHub API to create a new repository.
// It takes a request object, containing details for the new repo, and an access token for authentication.
// It returns either the successful response or an error response.
func CreateRepo(request github.CreateRequest, accessToken string) (*github.CreateResponse, *github.GithubErrorResponse) {
	// Set up headers with Authorization
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	// Send POST request to GitHub API using the restclient.Post function
	response, err := restclient.Post(urlCreateRepo, request, headers)

	// Log and handle errors
	if err != nil {
		log.Println("error when trying to create a new repo in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	// Read the response body
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()

	// Check for errors in the response status code
	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	// Unmarshal the successful response into a CreateResponse struct
	var result github.CreateResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create new repo successful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}

	// Return the successful response
	return &result, nil
}
