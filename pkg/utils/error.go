package utils

import (
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
	var errData *ServiceError = nil

	if err != nil {
		errData = &ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return errData
}
