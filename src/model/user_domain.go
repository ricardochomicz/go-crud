package model

import (
	"encoding/json"
	"fmt"
)

func (ud *userDomain) SetId(id string) {
	ud.id = id
}

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

// SetID implements UserDomainInterface.
func (ud *userDomain) SetID(id string) {

}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}

// EncryptPassword é um método da estrutura UserDomain que tem como objetivo
// criptografar a senha do usuário utilizando o algoritmo MD5.
