package errs

import (
	"fmt"
	"net/http"
)

// Common errors
var (
	ErrMethodNotAllowed = func(methodName string) CustomErr {
		return NewCustomErr(
			fmt.Sprintf("Just accept %v method.", methodName),
			http.StatusMethodNotAllowed,
		)
	}
	ErrSomethingWentWrong = func(err error) CustomErr {
		return NewCustomErr("Something went wrong", http.StatusInternalServerError)
	}
	ErrRequired = func(field string) CustomErr {
		return NewCustomErr(
			fmt.Sprintf("%v is required", field),
			http.StatusBadRequest,
		)
	}
	ErrInvalidType = func(field string) CustomErr {
		return NewCustomErr(
			fmt.Sprintf("%v is invalid type", field),
			http.StatusBadRequest,
		)
	}
	ErrNotFound = func(field string) CustomErr {
		return NewCustomErr(
			fmt.Sprintf("%v not found", field),
			http.StatusNotFound,
		)
	}
	ErrBodyNotAllowed   = NewCustomErr("Body not allowed", http.StatusBadRequest)
	ErrPermissionDenied = NewCustomErr("Permission denied", http.StatusForbidden)
)
