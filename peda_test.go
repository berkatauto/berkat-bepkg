package berkatbepkg

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// Sewaktu waktu ini masih bisa dipakai, jangan dihapus dulu ya.
// func TestGCFListArticle() {
// 	mconn := SetConnection("MONGOSTRING", "berkatauto")
// 	dataarticle := GetArticle(mconn, "articleSet")
// 	fmt.Println(dataarticle)
// }

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Fullname = "Rachma Nurhaliza"
	userdata.Username = "rachma"
	userdata.Password = "r123"
	userdata.JournalStatus = true
	userdata.Role = "admin"
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	CreateNewUserRole(mconn, "userLogin", userdata)
}

func TestPostArticle(t *testing.T) {
	var dataarticle Article
	var tagged Tags
	var isikonten Content
	// Tags
	tagged.Tag = "Mesin, " + "Ilmu Pengetahuan, " + "Teknologi"
	// Content
	isikonten.Paragraph = "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco labor is nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."
	isikonten.VideoContent = "NVI"
	// Craft All to new Article
	dataarticle.articleID = 
	dataarticle.Author = "Random Person"
	dataarticle.Title = "Pendidikan : Apa Itu Mobil?"
	dataarticle.Category = "Pendidikan Otomotif"
	dataarticle.Tags = tagged
	dataarticle.Content = isikonten
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	PostArticle(mconn, "articleSet", dataarticle)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "12345"
	hash, _ := HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

// func TestUserRandomNumber(t *testing.T) {
// }

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("rachma", privateKey)
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
