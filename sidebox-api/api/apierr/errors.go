package apierr

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, e responseError) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(e.code)
	json.NewEncoder(w).Encode(struct {
		ErrorName string `json:"errorName"`
	}{
		ErrorName: e.name,
	})
}

type responseError struct {
	name string
	code int
}

var BadAuthorization = responseError{
	name: "bad authorization",
	code: 401,
}

var InvalidBody = responseError{
	name: "invalid body",
	code: 400,
}

var InternalError = responseError{
	name: "internal error",
	code: 500,
}

var DatabaseError = responseError{
	name: "database error",
	code: 500,
}

var BadLogin = responseError{
	name: "bad login",
	code: 401,
}
