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

// func GCFLoginHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	var Response Credential
// 	Response.Status = false
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var datauser User
// 	err := json.NewDecoder(r.Body).Decode(&datauser)
// 	if err != nil {
// 		Response.Message = "error parsing application/json: " + err.Error()
// 	}

// 	if !IsUsernameExists(MONGOCONNSTRINGENV, dbname, datauser) {
// 		Response.Message = "Username or Password invalid. Please input the correct username and password."
// 	}

// 	if IsPasswordValid(mconn, collectionname, datauser) {
// 		Response.Message = "Username or Password invalid. Please input the correct username and password."
// 	} else {
// 		if IsPasswordValid(mconn, collectionname, datauser) {
// 			Response.Status = true
// 			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
// 			if err != nil {
// 				Response.Message = "Gagal Encode Token : " + err.Error()
// 			} else {
// 				Response.Message = "Welcome!"
// 				Response.Token = tokenstring
// 			}
// 		} else {
// 			Response.Message = "Username or Password invalid. Please input the correct username and password."
// 		}
// 	}
// 	return GCFReturnStruct(Response)
// }

// func GCFCreateUserWToken(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	var Response RegisterInfo
// 	Response.Status = false
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var datauser User
// 	err := json.NewDecoder(r.Body).Decode(&datauser)
// 	if err != nil {
// 		return err.Error()
// 	}

// 	if IsUsernameExists(MONGOCONNSTRINGENV, dbname, datauser) {
// 		Response.Message = "Username is already exists."
// 		return GCFReturnStruct(Response)
// 	}

// 	// Hash the password before storing it
// 	hashedPassword, hashErr := HashPassword(datauser.Password)
// 	if hashErr != nil {
// 		return hashErr.Error()
// 	}
// 	datauser.Password = hashedPassword
// 	CreateUserAndAddedToken(mconn, collectionname, datauser.Fullname, datauser.Username, hashedPassword, datauser.JournalStatus, datauser.Role)
// 	fmt.Println("User Creation Succesfull. User Information : ")
// 	Response.Fullname = datauser.Fullname
// 	Response.Username = datauser.Username
// 	Response.Password = datauser.Password
// 	Response.Status = true
// 	return GCFReturnStruct(datauser)
// }

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

	// Hash the user's password
	hash, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		response.Message = "Gagal hash password: " + hashErr.Error()
		return GCFReturnStruct(response)
	}

	// Insert user data into the database
	CreateUserAndAddedToken(mconn, collectionname, datauser.Fullname, datauser.Username, hash, datauser.JournalStatus, datauser.Role)
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

	if IsUsernameExists(MONGOCONNSTRINGENV, dbname, datauser) {
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

// func GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var datauser User
// 	err := json.NewDecoder(r.Body).Decode(&datauser)
// 	if err != nil {
// 		return err.Error()
// 	}

// 	// Hash the password before storing it
// 	hashedPassword, hashErr := HashPassword(datauser.Password)
// 	if hashErr != nil {
// 		return hashErr.Error()
// 	}
// 	datauser.Password = hashedPassword

// 	createErr := CreateNewUserRole(mconn, collectionname, datauser)
// 	fmt.Println(createErr)

// 	return GCFReturnStruct(datauser)
// }

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

func GCFLoadOneArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var searcharticle Article
	err := json.NewDecoder(r.Body).Decode(&searcharticle)
	if err != nil {
		return err.Error()
	}
	Load := LoadArticle(mconn, collectionname, searcharticle)
	// Decoding the base64 string
	Load.Content.Image = DecodeBase64String(Load.Content.Image)
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

// If Available, it will convert the video to base64
func UploadedVideoToBase64(file Content) {
	file.VideoContent = base64.StdEncoding.EncodeToString([]byte(file.VideoContent))
}

func GCFPostArticle(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var newarticle Article
	var fileConvert Content
	err := json.NewDecoder(r.Body).Decode(&newarticle)
	if err != nil {
		return err.Error()
	}
	response := GCFReturnStruct(newarticle)
	PostArticle(mconn, collectionname, newarticle)
	// Automatically If There is an Category
	// It will be added to the tags
	if newarticle.Category != "" {
		newarticle.Tags.Tag = newarticle.Tags.Tag
	}
	// Add category at the first line of title
	newarticle.Title = newarticle.Category + " : " + newarticle.Title

	// Add the date
	newarticle.Date = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
	// If Available, it will convert the image to base64
	fileConvert = newarticle.Content
	ConvertFileToBase64(fileConvert)
	UploadedVideoToBase64(fileConvert)
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

// func WAAuth(Whatsauth Whatsauth, tokenstring string) {
// 	var Username = Whatsauth.Username
// 	var nohp = Whatsauth.No_whatsapp
// 	dt := &wa.TextMessage{
// 		To:       nohp,
// 		IsGroup:  false,
// 		Messages: Username + " berhasil login\nNikmati Web Wisata di kota bandung\nIni token Untuk melanjutkan Proses selanjutnya yah.",
// 	}
// 	atapi.PostStructWithToken[atmessage.Response](dt, tokenstring)
// }
