package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/aasumitro/posbe/pkg/errors"
	"golang.org/x/crypto/scrypt"
	"strings"
)

type IPassword interface {
	HashPassword() (string, error)
	ComparePasswords() (bool, error)
}

type Password struct {
	Stored   string
	Supplied string
}

func (p *Password) HashPassword() (string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	shash, err := scrypt.Key([]byte(p.Supplied), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

func (p *Password) ComparePasswords() (bool, error) {
	pwsalt := strings.Split(p.Stored, ".")
	if len(pwsalt) < 2 {
		return false, errors.ErrorPasswordNotProvideValidHash
	}

	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, errors.ErrorPasswordUnableToVerify
	}

	shash, err := scrypt.Key([]byte(p.Supplied), salt, 32768, 8, 1, 32)
	if err != nil {
		return false, errors.ErrorPasswordUnableToVerify
	}

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
