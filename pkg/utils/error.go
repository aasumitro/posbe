package utils

import (
	"github.com/lib/pq"
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
		msg := err.Error()

		if pqErr, ok := err.(*pq.Error); ok {
			msg = pqErr.Error()
		}

		errData = &ServiceError{
			Code:    http.StatusInternalServerError,
			Message: msg,
		}
	}

	return errData
}