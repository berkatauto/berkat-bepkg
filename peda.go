package berkatbepkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataarticle := GetArticle(mconn, collectionname)
	return GCFReturnStruct(dataarticle)
}

// Ini masih kepakai, jangan dihapus dulu
// func UserRandomNumber(randomnumber string) {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	rand_source := rand.New(source)
// 	for i := 0; i < 5; i++ {
// 		rand_num := rand_source.Int()
// 	}
// }

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

// func GCFBuildContent(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {

// }

func GCFSearchArticleByTags(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var tagarticle Tags
	err := json.NewDecoder(r.Body).Decode(&tagarticle)
	if err != nil {
		return err.Error()
	}
	if tagarticle.Tag == "" {
		return "false"
	}

	tagresult := FindTag(mconn, collectionname, tagarticle)

	if tagresult != (Tags{}) {
		return GCFReturnStruct(tagresult)
	}

	return "false"
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

func GCFImageUploader(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Read the image file
	imagePath := "path/to/your/image.jpg"
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	bucket, err := gridfs.NewBucket(mconn)
	if err != nil {
		log.Fatal(err)
	}

	// Create a file in the GridFS bucket
	uploadStream, err := bucket.OpenUploadStream("image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadStream.Close()

	// Write the image data to the GridFS file
	_, err = uploadStream.Write(imageData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image Upload Success.")

	return "false" // Add this line to fix the "missing return" error

}
