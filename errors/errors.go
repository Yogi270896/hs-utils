package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type RestAPIError struct {
	Message          string `json:"message"`
	Status           int    `json:"status"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (o RestAPIError) String() string {
	return fmt.Sprintf("%#v", o)
}

func (o RestAPIError) IsNull() bool {
	return reflect.DeepEqual(o, RestAPIError{})
}

func (o RestAPIError) IsNotNull() bool {
	return !reflect.DeepEqual(o, RestAPIError{})
}

func (o RestAPIError) ToJson() string {
	js, serr := json.Marshal(o)
	if serr != nil {
		log.Println("Error while marshalling RestAPIError: ", serr)
	}
	return string(js)
}

func NO_ERROR() RestAPIError {
	var restErr RestAPIError
	return restErr
}

func HasError(restErr *RestAPIError) bool {
	return *restErr != NO_ERROR()
}

func NewError(text string) error {
	return errors.Unwrap(fmt.Errorf(text))
}

func NewRestAPIErrorFromBytes(body []byte) (RestAPIError, error) {
	var result RestAPIError
	if err := json.Unmarshal(body, &result); err != nil {
		return NO_ERROR(), errors.Unwrap(fmt.Errorf("invalid json body"))
	}
	return result, nil
}

func NewBadRequestError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}

func NewNotFoundError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not found",
	}
}

func NewInternalServerError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
}

func NewUnAuthorizedError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "User not authorized to access this resource",
	}
}

func NewForbiddenError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusForbidden,
		Error:   "User forbidden to access this resource",
	}
}

func NewDuplicateRecord(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusConflict,
		Error:   "Duplicate record found in database",
	}
}

func NewRedisNotCache(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Record not found in Redis cache",
	}
}
