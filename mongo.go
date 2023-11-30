package berkatbepkg

import (
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
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

func GetArticle(mongoconn *mongo.Database, collection string) []Article {
	tampilartikel := atdb.GetAllDoc[[]Article](mongoconn, collection)
	return tampilartikel
}

func UpdateArticle(mongoconn *mongo.Database, collection string, articleData Article) interface{} {
	return atdb.ReplaceOneDoc(mongoconn, collection, bson.M{"article_id": articleData.articleID}, articleData)
}

func DeleteArticle(mongoconn *mongo.Database, collection string, articleData Article) interface{} {
	return atdb.DeleteOneDoc(mongoconn, collection, bson.M{"article_id": articleData.articleID})
}

func CreateNewUserRole(mongoconn *mongo.Database, collection string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the user data into the database
	return atdb.InsertOneDoc(mongoconn, collection, userdata)
}

func CreateUserAndAddedToken(PASETOPRIVATEKEYENV string, mongoconn *mongo.Database, collection string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the user data into the database
	atdb.InsertOneDoc(mongoconn, collection, userdata)

	// Create a token for the user
	tokenstring, err := watoken.Encode(userdata.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		return err
	}
	userdata.Token = tokenstring

	// Update the user data in the database
	return atdb.ReplaceOneDoc(mongoconn, collection, bson.M{"username": userdata.Username}, userdata)
}

func FindAuthor(mongoconn *mongo.Database, collection string, author Article) Article {
	filter := bson.M{
		"author": author.Author,
	}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

func FindTag(mongoconn *mongo.Database, collection string, searchtag Tags) Tags {
	filter := bson.M{
		"tag": searchtag.Tag,
	}
	return atdb.GetOneDoc[Tags](mongoconn, collection, filter)
}
