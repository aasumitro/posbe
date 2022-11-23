package errors

import "fmt"

var (
	ErrorPasswordNotProvideValidHash = fmt.Errorf("did not provide a valid hash")
	ErrorPasswordUnableToVerify      = fmt.Errorf("unable to verify user password")
)
