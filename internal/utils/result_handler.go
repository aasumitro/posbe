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

func HandleSingleResult[T any](entity string, data *T, err error) (*T, *ServiceError) {
	svcErr := wrapServiceError(entity, err)
	return data, svcErr
}

func HandleMultipleResults[T any](entity string, data []*T, err error) ([]*T, *ServiceError) {
	svcErr := wrapServiceError(entity, err)
	return data, svcErr
}

func wrapServiceError(entity string, err error) *ServiceError {
	var svcErr *ServiceError
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			svcErr = &ServiceError{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("%s not found", entity),
			}
		default:
			svcErr = &ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}
	return svcErr
}
