package main

import (
	"fmt"
	"net/http"

	berkatbepkg "github.com/berkatauto/berkat-bepkg"
)

func HelloSignUp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFCreateUserWToken("PASETOPRIVATEKEYENV", "MONGOSTRING", "berkatauto", "userLogin", r))
}

func HelloSignIn(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFPostArticle("MONGOSTRING", "berkatauto", "articleSet", r))
}

func HelloDeleting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFPostArticle("MONGOSTRING", "berkatauto", "articleSet", r))
}

func HelloGetArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, berkatbepkg.GCFHandler("MONGOSTRING", "berkatauto", "articleSet", r))
}
