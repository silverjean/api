package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("\n %s %s %s", req.Method, req.RequestURI, req.Host)
		next(res, req)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if err := auth.TokenValidate(req); err != nil {
			responses.Err(res, http.StatusUnauthorized, err)
			return
		}
		next(res, req)
	}
}