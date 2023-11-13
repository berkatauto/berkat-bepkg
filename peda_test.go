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
	userdata.Fullname = "Adam Ghafara"
	userdata.Username = "adam"
	userdata.Password = "12345"
	userdata.JournalStatus = true
	userdata.Role = "admin"
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	CreateNewUserRole(mconn, "userLogin", userdata)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "12345"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

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
	hasil, err := watoken.Encode("adam", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "berkatauto")
	var userdata User
	userdata.Username = "adam"
	userdata.Password = "12345"

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
	userdata.Username = "adam"
	userdata.Password = "12345"
	anu := IsPasswordValid(mconn, "userLogin", userdata)
	fmt.Println(anu)
}
