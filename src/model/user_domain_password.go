package model

import (
	"crypto/md5"
	"encoding/hex"
)

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
