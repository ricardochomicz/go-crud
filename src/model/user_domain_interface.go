package model

import "github.com/ricardochomicz/go-crud/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetJSONValue() (string, error)
	GetID() string
	SetID(string)
	EncryptPassword()

	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func UserLoginDomain(
	email string,
	password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func UpdateUserDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
