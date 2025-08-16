package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type ServiceError struct {
	Code    int
	Message any
}

func ValidateDataRow[T any](coll string, data *T, err error) (valueData *T, errData *ServiceError) {
	errData = checkError(coll, err)
	return data, errData
}

func ValidateDataRows[T any](coll string, data []*T, err error) (valueData []*T, errData *ServiceError) {
	errData = checkError(coll, err)
	return data, errData
}

func checkError(coll string, err error) *ServiceError {
	var errData *ServiceError
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			errData = &ServiceError{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("%s not found", coll),
			}
		default:
			errData = &ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}
	return errData
}
