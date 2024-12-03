package view

import (
	"github.com/ricardochomicz/go-crud/src/model"
	"github.com/ricardochomicz/go-crud/src/model/response"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

func ConvertDomainsToResponses(
	userDomains []model.UserDomainInterface,
) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, userDomain := range userDomains {
		userResponses = append(userResponses, ConvertDomainToResponse(userDomain))
	}
	return userResponses
}
