package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
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

}
func FindPost(res http.ResponseWriter, req *http.Request){

}
func UpdatePost(res http.ResponseWriter, req *http.Request){

}
func DeletePost(res http.ResponseWriter, req *http.Request){

}
