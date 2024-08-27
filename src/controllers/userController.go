package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	var userModel models.User

	if err = json.Unmarshal(requestBody, &userModel); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	if err = userModel.Prepare("register"); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)

	userModel.ID, err = repo.Create(userModel)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
	}

	responses.JSON(res, http.StatusCreated, userModel)
}

func FindUsers(res http.ResponseWriter, req *http.Request) {
	nameOrNick := strings.ToLower(req.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)
	users, err := repo.Find(nameOrNick)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(res, http.StatusOK, users)

}

func FindUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	
	userID, err := strconv.ParseUint(params["userId"], 10, 64);
	if err != nil {
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
	user, err := repo.FindOne(userID)
	if err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	responses.JSON(res, http.StatusOK, user)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID, err := strconv.ParseUint(params["userId"], 10, 64);
	if err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

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

	if err = userModel.Prepare("edit"); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}		

	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)
	if err = repo.UpdateUser(userID, userModel); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	responses.JSON(res, http.StatusNoContent, nil)
}

func DeleteUsers(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID, err := strconv.ParseUint(params["userId"], 10, 64);
	if err != nil {
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
	if err = repo.Delete(userID); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	responses.JSON(res, http.StatusNoContent, err)
}