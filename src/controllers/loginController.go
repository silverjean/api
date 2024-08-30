package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	var userModel models.User
	if err = json.Unmarshal(reqBody, &userModel); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)
	dbUser, err := repo.FindByEmail(userModel.Email)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	if err = security.ValidatePass(dbUser.Password, userModel.Password); err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.TokenCreate(userModel.ID)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	res.Write([]byte(token))
}