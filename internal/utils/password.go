package utils

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

const (
	// CPU/memory cost parameter
	cost = 32768
	// The block mixing parameter.
	// This controls the amount of memory used by
	// the algorithm and the degree of parallelism.
	mixing = 8
	// The Parallelization parameter.
	// This controls the number of independent memory blocks
	// that are processed in parallel.
	Parallelization = 1
	keyLen          = 32
	maxSplit        = 2
	byteSize        = 32
)

var (
	ErrorPasswordHashNotValid   = fmt.Errorf("password hash not valid")
	ErrorPasswordUnableToVerify = fmt.Errorf("unable to verify password")
)

func MakePassword(par int, supplied string) (string, error) {
	var scryptHash []byte
	var err error

	salt := make([]byte, byteSize)
	if _, err = cryptoRand.Read(salt); err != nil {
		return "", err
	}

	if scryptHash, err = scrypt.Key(
		[]byte(supplied),
		salt, cost, mixing,
		par, keyLen,
	); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s.%s",
		hex.EncodeToString(scryptHash),
		hex.EncodeToString(salt),
	), nil
}

func ComparePassword(par int, stored, supplied string) (bool, error) {
	var scryptHash []byte
	var salt []byte
	var err error

	pwdSalt := strings.Split(stored, ".")
	if len(pwdSalt) < maxSplit {
		return false, ErrorPasswordHashNotValid
	}

	if salt, err = hex.DecodeString(pwdSalt[1]); err != nil {
		return false, ErrorPasswordUnableToVerify
	}

	if scryptHash, err = scrypt.Key(
		[]byte(supplied),
		salt, cost, mixing,
		par, keyLen,
	); err != nil {
		return false, ErrorPasswordUnableToVerify
	}

	return hex.EncodeToString(scryptHash) == pwdSalt[0], nil
}
