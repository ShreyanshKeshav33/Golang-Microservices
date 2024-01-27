package services

import (
	"strings"

	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/config"
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/domain/github"
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/domain/repositories"
	githubprovider "github.com/ShreyanshKeshav33/Golang-Microservices/src/api/providers/github_provider"
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRequest) (*repositories.CreateResponse, errors.APIError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRequest) (*repositories.CreateResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}
	request := github.CreateRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}
	response, err := githubprovider.CreateRepo(request, config.GetGithubAccessToken())
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
