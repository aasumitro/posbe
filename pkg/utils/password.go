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

const cost, r, p, k = 32768, 8, 1, 32
const maxSplit = 2
const byteSize = 32

func (pwd *Password) HashPassword() (string, error) {
	salt := make([]byte, byteSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	shash, err := scrypt.Key([]byte(pwd.Supplied), salt, cost, r, p, k)
	if err != nil {
		return "", err
	}

	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

func (pwd *Password) ComparePasswords() (bool, error) {
	pwsalt := strings.Split(pwd.Stored, ".")
	if len(pwsalt) < maxSplit {
		return false, errors.ErrorPasswordNotProvideValidHash
	}

	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, errors.ErrorPasswordUnableToVerify
	}

	shash, err := scrypt.Key([]byte(pwd.Supplied), salt, cost, r, p, k)
	if err != nil {
		return false, errors.ErrorPasswordUnableToVerify
	}

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
