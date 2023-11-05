package berkatbepkg

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataarticle := GetArticle(mconn, collectionname)
	return GCFReturnStruct(dataarticle)
}

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Welcome!"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah! Silahkan Coba Lagi."
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// func CreateUser(mongoenv, dbname, collname string, r *http.Request) string {
// 	var response Credential
// 	response.Status = false
// 	mconn := SetConnection(mongoenv, dbname)
// 	var datauser User
// 	err := json.NewDecoder(r.Body).Decode(&datauser)
// 	if err != nil {
// 		response.Message = "error parsing application/json: " + err.Error()
// 	} else {
// 		response.Status = true
// 		hash, hashErr := HashPassword(datauser.Password)
// 		if hashErr != nil {
// 			response.Message = "Gagal Hash Password" + err.Error()
// 		}
// 		(mconn, collname, datauser.Username, datauser.Role, hash)
// 		response.Message = "Berhasil Input data"
// 	}
// 	return GCFReturnStruct(response)
// }
