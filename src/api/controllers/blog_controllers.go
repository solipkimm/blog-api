package controllers

import "net/http"

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List posts"))
}
