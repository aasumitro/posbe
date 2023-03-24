package commons

import (
	"errors"
)

var (
	// ErrorDatabaseRowNotExist = fmt.Errorf("row does not exist")

	ErrorPasswordNotProvideValidHash = errors.New("did not provide a valid hash")
	ErrorPasswordUnableToVerify      = errors.New("unable to verify user password")
	ErrorUnableToDelete              = errors.New("unable to delete this data")
)
