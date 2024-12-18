package service

import (
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
	"github.com/ricardochomicz/go-crud/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

// CreateUserService implements UserDomainService.

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(
		userId string,
		userDomain model.UserDomainInterface) *rest_err.RestErr

	FindAllUsersService() ([]model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDService(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(userId string) *rest_err.RestErr

	LoginUserService(userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, *rest_err.RestErr)
}
