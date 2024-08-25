package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(res http.ResponseWriter, statusCode int, data interface{}) {
	res.WriteHeader(statusCode)

	if err := json.NewEncoder(res).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Err(res http.ResponseWriter, statusCode int, err error) {
	JSON(res, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}