package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateRepoRequestAsJson is a test function that checks the JSON serialization of a CreateRequest.
func TestCreateRepoRequestAsJson(t *testing.T) {
	// Create a sample CreateRequest
	request := CreateRequest{
		Name:        "golang intro",
		Description: "a golang intro repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}
	//Marsahl takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)

	// Use assertions to check if there is no error during JSON marshaling
	//The purpose of this line is to assert that err is nil, meaning that no error occurred during the execution of json.Marshal(request).
	assert.Nil(t, err)
	// Use assertions to check if the resulting JSON bytes are not nil
	//The purpose of this line is to assert that bytes is not nil. This ensures that the marshaling operation produced some valid data.
	assert.NotNil(t, bytes)

	// Print the JSON string for manual inspection
	fmt.Println(string(bytes))

	var target CreateRequest
	//bytes: The byte slice containing the JSON-encoded data.
	//&target: The address of the target variable (target).
	err = json.Unmarshal(bytes, &target)

	/*// Verify that the resulting JSON string matches the expected JSON string
	assert.EqualValues(t, `{"name":"golang intro","description":"a golang intro repo","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))
	*/
}
