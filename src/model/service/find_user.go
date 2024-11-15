package service

import (
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
