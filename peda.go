package berkatbepkg

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/whatsauth/watoken"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataarticle := GetArticle(mconn, collectionname)
	return GCFReturnStruct(dataarticle)
}

func DecodeBase64String(data string) string {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err.Error()
	}
	return string(decoded)
}

func GCFCreateUserWToken(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var response RegisterInfo
	response.Status = false

	// Establish MongoDB connection
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Decode user data from the request body
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)

	// Check for JSON decoding errors
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	// Check if the username already exists
	if IsUsernameExists(MONGOCONNSTRINGENV, dbname, datauser) {
		response.Message = "Username is already exists."
		return GCFReturnStruct(response)
	}

	// Insert user data into the database
	CreateUserAndAddedToken(mconn, collectionname, datauser)
	response.Status = true
	response.Message = "Input Successful with Information: "
	response.Fullname = datauser.Fullname
	response.Username = datauser.Username
	response.Password = datauser.Password
	return GCFReturnStruct(response)
}

func GCFLoginHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false

	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)

	if err != nil {
		Response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(Response)
	}

	if !IsUsernameExists(MONGOCONNSTRINGENV, dbname, datauser) {
		Response.Message = "Username or Password invalid. Please input the correct username and password."
		return GCFReturnStruct(Response)
	}

	if !IsPasswordValid(mconn, collectionname, datauser) {
		Response.Message = "Password Salah"
		return GCFReturnStruct(Response)
	}

	authentication := FindUser(mconn, collectionname, datauser)

	tokenstring, err := watoken.Encode(authentication.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		Response.Message = "Gagal Encode Token : " + err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = "Welcome!"
	Response.Token = tokenstring

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func GCFSearchByCategory(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var categoryarticle Article
	err := json.NewDecoder(r.Body).Decode(&categoryarticle)
	if err != nil {
		return err.Error()
	}
	if categoryarticle.Category == "" {
		return "false"
	}
	return GCFReturnStruct(SearchByCategory(mconn, collectionname, categoryarticle))
}

func GCFSearchByTitle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var searcharticle Article
	err := json.NewDecoder(r.Body).Decode(&searcharticle)
	if err != nil {
		return err.Error()
	}
	find := SearchByTitle(mconn, collectionname, searcharticle)
	return GCFReturnStruct(find)
}

func GCFSearchByTags(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var searcharticle Article
	err := json.NewDecoder(r.Body).Decode(&searcharticle)
	if err != nil {
		return err.Error()
	}
	find := SearchByTags(mconn, collectionname, searcharticle)
	return GCFReturnStruct(find)
}

func GCFSearchByAuthor(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var searcharticle Article
	err := json.NewDecoder(r.Body).Decode(&searcharticle)
	if err != nil {
		return err.Error()
	}
	find := SearchByAuthor(mconn, collectionname, searcharticle)
	return GCFReturnStruct(find)
}

func GCFLoadOneArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var searcharticle Article
	err := json.NewDecoder(r.Body).Decode(&searcharticle)
	if err != nil {
		return err.Error()
	}
	Load := LoadArticle(mconn, collectionname, searcharticle)
	// Date Only Load Day/Month/Year
	Load.Date = time.Date(Load.Date.Year(), Load.Date.Month(), Load.Date.Day(), 0, 0, 0, 0, time.UTC)
	return GCFReturnStruct(Load)
	// Deploy to HTML
}

func GetArticleByLastDate(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, collectionname)
	bylastdate := GetByLastDate(mconn, collectionname, Article{})
	return GCFReturnStruct(bylastdate)
}

func ConvertFileToBase64(file Content) {
	// file.ImageHeader = base64.StdEncoding.EncodeToString([]byte(file.ImageHeader))
	file.Image = base64.StdEncoding.EncodeToString([]byte(file.Image))
}

func GCFPostArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var newarticle Article
	err := json.NewDecoder(r.Body).Decode(&newarticle)
	if err != nil {
		return err.Error()
	}
	response := GCFReturnStruct(newarticle)
	PostArticle(mconn, collectionname, newarticle)
	newarticle.Date = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
	return response
}

func GCFDeleteArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var deleteArticle Article
	err := json.NewDecoder(r.Body).Decode(&deleteArticle)
	if err != nil {
		return err.Error()
	}
	response := GCFReturnStruct("Deleting Successful.")
	DeleteArticle(mconn, collectionname, deleteArticle)
	return response
}

func GCFUpdateArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var updateArticle Article
	err := json.NewDecoder(r.Body).Decode(&updateArticle)
	if err != nil {
		return err.Error()
	}
	response := GCFReturnStruct(updateArticle)
	UpdateArticle(mconn, collectionname, updateArticle)
	return response
}
