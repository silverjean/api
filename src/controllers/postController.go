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

	"github.com/gorilla/mux"
)

func CreatePost(res http.ResponseWriter, req *http.Request){
	userId, err := auth.ExtractUserID(req)
	if err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	}

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(reqBody, &post); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userId

	if err = post.Prepare(); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}



	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryPost(db)
	post.ID, err = repo.Create(post)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(res, http.StatusCreated, post)

}
func FindPosts(res http.ResponseWriter, req *http.Request){
	userID, err := auth.ExtractUserID(req)
	if err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryPost(db)
	posts, err := repo.FindAllPosts(userID)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	} 

	responses.JSON(res, http.StatusOK, posts)
}
func FindPost(res http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
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

	repo := repositories.NewRepositoryPost(db)
	post, err := repo.FindPostByID(postID)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(res, http.StatusOK, post)
}
func UpdatePost(res http.ResponseWriter, req *http.Request){
	userID, err := auth.ExtractUserID(req)
	if err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(req)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
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

	repo := repositories.NewRepositoryPost(db)
	postOnDB, err := repo.FindPostByID(postID)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	if postOnDB.AuthorID != userID {
		responses.Err(res, http.StatusForbidden, errors.New("it is not possible update post that is not yours"))
		return
	}

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post

	if err = json.Unmarshal(reqBody, &post); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePost(postID, post); err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(res, http.StatusOK, nil)
}
func DeletePost(res http.ResponseWriter, req *http.Request){

}
