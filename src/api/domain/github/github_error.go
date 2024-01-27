package github

// GithubErrorResponse represents the error response structure from the GitHub API.
type GithubErrorResponse struct {
	StatusCode int           `json:"status_code"`
	Message    string        `json:"message"`
	DocumentUr string        `json:"documentation_url"`
	Errors     []GithubError `json:"errors"`
}

// GithubError represents an individual error in the GitHub API error response.
type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
