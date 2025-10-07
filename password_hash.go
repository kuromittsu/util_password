package util_password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func passwordHash(rawPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	switch err {
	case bcrypt.ErrPasswordTooLong:
		return "", errors.New("password too long")
		// return "", errors.New("password too long (> 72 bytes)")
	default:
		return string(hash), nil
	}
}

func passwordCompare(hashedPassword, rawPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	switch err {
	case bcrypt.ErrMismatchedHashAndPassword:
		return errors.New("hash and password not match")
	case bcrypt.ErrHashTooShort:
		return errors.New("hash is too short")
	default:
		return nil
	}
}
