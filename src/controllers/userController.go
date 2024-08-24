package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewRepositoryUser(db)

	userID, err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	res.Write([]byte(fmt.Sprintf("Inserted ID: %d", userID)))
}

func FindUsers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Finding all users"))
}

func FindUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Finding one user"))
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Updating one user"))
}

func DeleteUsers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Deleting one user"))
}