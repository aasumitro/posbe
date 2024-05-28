package utils

import (
	"database/sql"
	"errors"
	"net/http"
)

type ServiceError struct {
	Code    int
	Message any
}

func ValidateDataRow[T any](data *T, err error) (valueData *T, errData *ServiceError) {
	errData = checkError(err)
	return data, errData
}

func ValidateDataRows[T any](data []*T, err error) (valueData []*T, errData *ServiceError) {
	errData = checkError(err)
	return data, errData
}

func checkError(err error) *ServiceError {
	var errData *ServiceError
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			errData = &ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
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
