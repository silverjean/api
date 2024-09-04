package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
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

	userIDOnToken, err := auth.ExtractUserID(req)
	if err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	if userID != userIDOnToken {
		responses.Err(res, http.StatusForbidden, 
			errors.New("it is not possible to update a user other than yours"))
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

	userIDOnToken, err := auth.ExtractUserID(req)
	if err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDOnToken {
		responses.Err(res, http.StatusForbidden, errors.New("it is not possible to delete a user other than yours"))
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

func FollowUser (res http.ResponseWriter, req *http.Request) {
	followerID, err := auth.ExtractUserID(req)
	if err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(req)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Err(res, http.StatusForbidden, errors.New("it is not possible follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)
	if err = repo.Follow(userID, followerID); err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return 
	}

	responses.JSON(res, http.StatusNoContent, nil)
}

func UnfollowUser (res http.ResponseWriter, req *http.Request) {
	followerID, err := auth.ExtractUserID(req)
	if err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(req)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Err(res, http.StatusForbidden, errors.New("it is not possible unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)
	if err = repo.Unfollow(userID, followerID); err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return 
	}

	responses.JSON(res, http.StatusNoContent, nil)
}

func FollowersUser (res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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
	followers, err := repo.FindFollowers(userID)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return 
	}

	responses.JSON(res, http.StatusOK, followers)
}

func FollowingUser (res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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
	following, err := repo.FindFollowing(userID)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return 
	}

	responses.JSON(res, http.StatusOK, following)
}