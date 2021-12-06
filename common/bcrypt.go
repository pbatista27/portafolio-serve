package common

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct{}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (b Bcrypt) Encriptar(password string) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		return hash, err
	}

	return hash, nil
}

func (b Bcrypt) Descriptar(password, dbpassword string) (bool, error) {

	pass := []byte(password)
	dbpass := []byte(dbpassword)

	err := bcrypt.CompareHashAndPassword(dbpass, pass)

	if err != nil {
		return false, err
	}

	return true, nil
}
