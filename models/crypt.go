package models

import (
	"golang.org/x/crypto/bcrypt"
)

type crypt struct {
	cost int
}

func newCrypt() *crypt {
	return &crypt{cost: bcrypt.DefaultCost}
}

func (c crypt) hash(s string) (string, error) {

	h, err := bcrypt.GenerateFromPassword([]byte(s), c.cost)

	if err != nil {
		return "", err
	}

	return string(h), nil

}

func (c crypt) check(s string, h string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(s))

	if err != nil {
		return false
	}

	return true

}
