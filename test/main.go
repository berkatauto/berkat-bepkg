package main

import (
	"fmt"
	"net/http"

	berkatbepkg "github.com/berkatauto/berkat-bepkg"
)

func HelloSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "3600")
	if r.Method == "OPTIONS" {

		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFCreateUserWToken("MONGOSTRING", "berkatauto", "userLogin", r))
}

func HelloSignIn(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFLoginHandler("PASETOPRIVATEKEYENV", "MONGOSTRING", "berkatauto", "userLogin", r))
}

func HelloPostArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFPostArticle("MONGOSTRING", "berkatauto", "articleSet", r))
}

func HelloEditArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFUpdateArticle("MONGOSTRING", "berkatauto", "articleSet", r))
}

func HelloDeleting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFDeleteArticle("MONGOSTRING", "berkatauto", "articleSet", r))
}

func HelloGetArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFHandler("MONGOSTRING", "berkatauto", "articleSet", r))
}

func HelloUpdateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}
	fmt.Fprintf(w, berkatbepkg.GCFUpdateArticle("MONGOSTRING", "berkatauto", "articleSet", r))
}

func HelloGetArticleByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}
	fmt.Fprintf(w, berkatbepkg.GetArticleByLastDate("MONGOSTRING", "berkatauto", "articleSet"))
}

func main() {
	http.HandleFunc("/signup", HelloSignUp)
	http.HandleFunc("/signin", HelloSignIn)
	http.HandleFunc("/postarticle", HelloPostArticle)
	http.HandleFunc("/editarticle", HelloEditArticle)
	http.HandleFunc("/deletearticle", HelloDeleting)
	http.HandleFunc("/getarticle", HelloGetArticle)
	http.HandleFunc("/updatearticle", HelloUpdateArticle)
	http.ListenAndServe(":8080", nil)
}
