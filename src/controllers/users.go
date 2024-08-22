package controllers

import "net/http"

func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Creating user"))
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