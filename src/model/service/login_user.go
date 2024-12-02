package service

import (
	"github.com/ricardochomicz/go-crud/src/configuration/logger"
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init loginUser model", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordService(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("journey", user.GetID()),
		zap.String("journey", "createUser"))
	return user, nil
}
