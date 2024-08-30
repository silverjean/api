package middlewares

import (
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
		fmt.Println("Autenticando...")
		next(res, req)
	}
}