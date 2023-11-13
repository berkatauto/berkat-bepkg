package berkatbepkg

import (
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GetArticle(mongoconn *mongo.Database, collection string) []Article {
	tampilartikel := atdb.GetAllDoc[[]Article](mongoconn, collection)
	return tampilartikel
}

func GetNameAndPassowrd(mongoconn *mongo.Database, collection string) []User {
	user := atdb.GetAllDoc[[]User](mongoconn, collection)
	return user
}

func SearchArticle(mongoconn *mongo.Database, collection string, searcharticle Article) Article {
	filter := bson.M{"Title": searcharticle.Title,
		"Category": searcharticle.Category,
		"Tags":     searcharticle.Tags}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

// func ResultArticle(mongoconn *mongo.Database, collection string, resultsearch Article) Article {
// 	showresult := atdb.GetAllDoc[Article](mongoconn, collection)

// }

func GetAllUser(mongoconn *mongo.Database, collection string) []User {
	user := atdb.GetAllDoc[[]User](mongoconn, collection)
	return user
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mongoconn, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}

func PostArticle(mongoconn *mongo.Database, collection string, articleData Article) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, articleData)
}
