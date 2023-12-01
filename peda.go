package berkatbepkg

import (
	"encoding/json"
	"fmt"
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
	CreateUserAndAddedToken(PASETOPRIVATEKEYENV, mconn, collectionname, datauser)
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

	categoryresult := SearchByCategory(mconn, collectionname, categoryarticle)

	if categoryresult != (Article{}) {
		return GCFReturnStruct(categoryresult)
	}

	return "false"
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

func GCFGetOneArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var searcharticle Article
	err := json.NewDecoder(r.Body).Decode(&searcharticle)
	if err != nil {
		return err.Error()
	}
	find := GetOneArticle(mconn, collectionname, searcharticle)
	return GCFReturnStruct(find)
}

// func GetByLastDate(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	dataarticle := GetByLastDate(mconn, collectionname)
// 	return GCFReturnStruct(dataarticle)
// }

func GCFPostArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var newcontent Content
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
	PostArticle(mconn, collectionname, newcontent, newarticle)
	return response
}

func GCFDeleteArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var deleteArticle Article
	err := json.NewDecoder(r.Body).Decode(&deleteArticle)
	if err != nil {
		return err.Error()
	}
	response := GCFReturnStruct(deleteArticle)
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

// func SearchArticleByCategory(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var categoryarticle Article
// 	err := json.NewDecoder(r.Body).Decode(&categoryarticle)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	if categoryarticle.Category == "" {
// 		return "false"
// 	}

// 	categoryresult := SearchByCategory(mconn, collectionname, categoryarticle)

// 	if categoryresult != (Article{}) {
// 		return GCFReturnStruct(categoryresult)
// 	}

// 	return "false"
// }

// func GCFSearchArticleByTags(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var tagarticle Tags
// 	err := json.NewDecoder(r.Body).Decode(&tagarticle)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	if tagarticle.Tag == "" {
// 		return "false"
// 	}

// 	tagresult := SearchByTags(mconn, collectionname, tagarticle)

// 	if tagresult != (Tags{}) {
// 		return GCFReturnStruct(tagresult)
// 	}

// 	return "false"
// }

// func GCFSearchArticleByUserId(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var userarticle Article
// 	err := json.NewDecoder(r.Body).Decode(&userarticle)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	if userarticle.Author == "" {
// 		return "false"
// 	}

// 	author := FindAuthor(mconn, collectionname, userarticle)

// 	if author != (Article{}) {
// 		return GCFReturnStruct(author)
// 	}

// 	return "false"
// }
