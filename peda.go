package berkatbepkg

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/whatsauth/watoken"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataarticle := GetArticle(mconn, collectionname)
	return GCFReturnStruct(dataarticle)
}

func UserRandomNumber() {
	source := rand.NewSource(time.Now().UnixNano())
	rand_source := rand.New(source)
	for i := 0; i < 5; i++ {
		rand_num := rand_source.Int()
		fmt.Println(rand_num)
	}
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

func GCFCreateUser(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}

	// Hash the password before storing it
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword

	createErr := CreateNewUserRole(mconn, collectionname, datauser)
	fmt.Println(createErr)

	return GCFReturnStruct(datauser)
}

func GCFCreateHandlerTokenPaseto(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword
	CreateNewUserRole(mconn, collectionname, datauser)
	tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		return err.Error()
	}
	datauser.Token = tokenstring
	return GCFReturnStruct(datauser)
}

func GCFCreateAccountAndToken(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword
	CreateUserAndAddedToeken(PASETOPRIVATEKEYENV, mconn, collectionname, datauser)
	return GCFReturnStruct(datauser)
}

func GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}

	// Hash the password before storing it
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword

	createErr := CreateNewUserRole(mconn, collectionname, datauser)
	fmt.Println(createErr)

	return GCFReturnStruct(datauser)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func GCFSearchArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var searcharticle Article
	err := json.NewDecoder(r.Body).Decode(&searcharticle)
	if err != nil {
		return err.Error()
	}
	find := SearchArticle(mconn, collectionname, searcharticle)
	return GCFReturnStruct(find)
}

func GCFPostArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var newarticle Article
	err := json.NewDecoder(r.Body).Decode(&newarticle)
	if err != nil {
		return err.Error()
	}

	// Get PASETO Header Value
	// pasetoValue := r.Header.Set("PASETOPRIVATEKEYENV")
	// Post The Article
	response := GCFReturnStruct(newarticle)
	// response += "PASETO Value: " + pasetoValue
	PostArticle(mconn, collectionname, newarticle)
	return response
}

func GCFBuildContent() {

}

func GCFSearchArticleByUserId(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userarticle Article
	err := json.NewDecoder(r.Body).Decode(&userarticle)
	if err != nil {
		return err.Error()
	}
	if userarticle.Author == "" {
		return "false"
	}

	author := FindAuthor(mconn, collectionname, userarticle)

	if author != (Article{}) {
		return GCFReturnStruct(author)
	}

	return "false"
}
