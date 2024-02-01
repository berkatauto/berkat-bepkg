package berkatbepkg

import (
	"fmt"
	"testing"
	"time"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// func TestCreateNewUserRole(t *testing.T) {
// 	var userdata User
// 	userdata.Fullname = "Rachma Nurhaliza"
// 	userdata.Username = "rachma"
// 	userdata.Password = "rachma123"
// 	userdata.JournalStatus = "true"
// 	userdata.Role = "admin"
// 	mconn := SetConnection("MONGOSTRING", "berkatauto")
// }

func TestCreateUserWToken(t *testing.T) {
	var userdata User
	userdata.Fullname = "Rachma Nurhaliza"
	userdata.Username = "rachmanurhaliza"
	userdata.Password = "rachma123"
	userdata.JournalStatus = "true"
	userdata.Role = "admin"
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	CreateUserAndAddedToken(mconn, "userLogin", userdata.Fullname, userdata.Username, userdata.Password, userdata.JournalStatus, userdata.Role)
}

func TestPostArticle(t *testing.T) {
	var dataarticle Article
	var tagged Tags
	var isikonten Content
	// var authorname User
	var date = time.Now()
	// var getID RandomNumber
	// Tags
	tagged.Tag = "Tips, " + "Perawatan, " + "Kendaraan, " + "Mobil," + "Motor"
	// Content
	isikonten.Paragraph = "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco labor is nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."
	isikonten.Image = "https://cdn1-production-images-kly.akamaized.net/7PK316VPijt24ISrrQTQzaJK_Eg=/1200x900/smart/filters:quality(75):strip_icc():format(jpeg)/kly-media-production/medias/4239659/original/011356600_1669372008-Mazda-RX7-1999-1600-01.jpg"
	isikonten.VideoContent = "NVI"
	// Craft All to new Article
	dataarticle.Author = "Baba Rafi Gunawan"
	dataarticle.Title = "Tips Merawat Kendaraan Dengan Baik"
	dataarticle.Category = "Belajar Otomotif"
	dataarticle.Tags = tagged
	dataarticle.Content = isikonten
	dataarticle.Date = date.UTC()
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	PostArticle(mconn, "articleSet", dataarticle)
}

func TestUpdateArticle(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	var dataarticle Article
	dataarticle.Author = "Baba Rafi"
	dataarticle.Title = "Tips : Memasang Kaca Film Mobil"
	dataarticle.Category = "Tips"
	dataarticle.Tags.Tag = "Tips, " + "Perawatan, " + "Kendaraan, " + "Mobil"
	UpdateArticle(mconn, "articleSet", dataarticle)
}

func TestDeleteArticle(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	var dataarticle Article
	dataarticle.Title = "Tips : Merawat Suspensi Mobil Agar Awet"
	DeleteArticle(mconn, "articleSet", dataarticle)
}

func TestGetArticle(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	dataarticle := GetArticle(mconn, "articleSet")
	fmt.Println(dataarticle)
}

func TestGetByLastDate(t *testing.T) {
	var dataarticle Article
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	GetByLastDate(mconn, "articleSet", dataarticle)
}

func TestGetOneArticle(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	var dataarticle Article
	dataarticle.Title = "Tips : Memasang Kaca Film Mobil"
	GetOneArticle(mconn, "articleSet", dataarticle)
}

func TestSearchByCategory(t *testing.T) {
	var searchBy Article
	// searchBy.Title = "Tips : Merawat Suspensi Mobil Agar Awet"
	searchBy.Category = "Tips"
	// searchBy.Tags.Tag = "Tips, " + "Perawatan, " + "Kendaraan, " + "Mobil"
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	SearchByCategory(mconn, "articleSet", searchBy)
}

func TestSearchByTitle(t *testing.T) {
	var searchBy Article
	searchBy.Title = "Tips : Memasang Kaca Film Mobil"
	// searchBy.Category = "Tips"
	// searchBy.Tags.Tag = "Tips, " + "Perawatan, " + "Kendaraan, " + "Mobil"
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	SearchByTitle(mconn, "articleSet", searchBy)
}

func TestSearchByTags(t *testing.T) {
	var searchBy Article
	// searchBy.Title = "Tips : Memasang Kaca Film Mobil"
	// searchBy.Category = "Tips"
	searchBy.Tags.Tag = "Tips, " + "Perawatan, " + "Kendaraan, " + "Mobil"
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	SearchByTags(mconn, "articleSet", searchBy)
}

// func SearchByAuthor(t *testing.T) {
// 	var searchBy Article
// 	searchBy.Author = "Baba Rafi"
// 	// searchBy.Category = "Tips"
// 	// searchBy.Tags.Tag = "Tips, " + "Perawatan, " + "Kendaraan, " + "Mobil"
// 	mconn := SetConnection("MONGOSTRING", "berkatauto")
// 	SearchByAuthor(mconn, "articleSet", searchBy)
// }

func TestGeneratePasswordHash(t *testing.T) {
	password := "12345"
	hash, _ := HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("superadmin", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	var userdata User
	userdata.Username = "rachma"
	userdata.Password = "r123"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "userLogin", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	var userdata User
	userdata.Username = "rachma"
	userdata.Password = "r123"
	anu := IsPasswordValid(mconn, "userLogin", userdata)
	fmt.Println(anu)
}
