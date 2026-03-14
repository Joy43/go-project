package controllers

import (
	"encoding/json"
	"go-jwt-auth/config"
	"go-jwt-auth/models"
	"net/http"
)

func CreatePostnew(w http.ResponseWriter,r *http.Request){
username:= r.Context().Value("username").(string)
var user models.User
if err:=config.DB.Where("username=?",username).First(&user).Error;err!=nil{
	http.Error(w,"User not found",http.StatusNotFound)
	return
}
var comment models.Comment
if err:=json.NewDecoder(r.Body).Decode(&comment);err!=nil{
	http.Error(w,"Invalid request payload",http.StatusBadRequest)
	return
}
comment.UserID=user.ID
if err:=config.DB.Create(&comment).Error;err!=nil{
	http.Error(w,"Error saving comment to database",http.StatusInternalServerError)
	return
}
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(comment)
}