package common

import (
	"golang.org/x/crypto/bcrypt"
)

type HashService interface {
	Create(plain string) (string, error)
	Validate(hash string, plain string) bool
}

type hashService struct {
}

func NewHashService() HashService {
	return &hashService{}
}

func (s hashService) Create(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s hashService) Validate(hash string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
