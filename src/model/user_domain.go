package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email,
		password,
		name,
		age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
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
func (ud *userDomain) EncryptPassword() {
	// Cria um novo hash MD5.
	hash := md5.New()
	// Garante que o hash será redefinido após o uso, evitando vazamentos de memória.
	defer hash.Reset()
	// Escreve a senha do usuário no hash. A senha é convertida para bytes.
	hash.Write([]byte(ud.password))
	// Calcula o hash e converte o resultado para uma string hexadecimal.
	// O resultado é então atribuído de volta à propriedade Password do usuário.
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
