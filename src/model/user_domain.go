package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetJSONValue() (string, error)
	SetID(string)
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *userDomain) SetId(id string) {
	ud.ID = id
}

type userDomain struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
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
func (ud *userDomain) SetID(string) {
	panic("unimplemented")
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

// EncryptPassword é um método da estrutura UserDomain que tem como objetivo
// criptografar a senha do usuário utilizando o algoritmo MD5.
func (ud *userDomain) EncryptPassword() {
	// Cria um novo hash MD5.
	hash := md5.New()
	// Garante que o hash será redefinido após o uso, evitando vazamentos de memória.
	defer hash.Reset()
	// Escreve a senha do usuário no hash. A senha é convertida para bytes.
	hash.Write([]byte(ud.Password))
	// Calcula o hash e converte o resultado para uma string hexadecimal.
	// O resultado é então atribuído de volta à propriedade Password do usuário.
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
