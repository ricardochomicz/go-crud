package service

import (
	"github.com/ricardochomicz/go-crud/src/configuration/logger"
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDService(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById service", zap.String("journey", "findUserById"))
	return ud.userRepository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailService(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail service", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordService(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword service", zap.String("journey", "findUserByEmailAndPassword"))
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
