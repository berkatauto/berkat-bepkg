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

func SearchByAuthor(mongoconn *mongo.Database, collection string, searchBy Article) Article {
	filter := bson.M{"author": searchBy.Author}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

func SearchByCategory(mongoconn *mongo.Database, collection string, searchBy Article) Article {
	filter := bson.M{"category": searchBy.Category}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

func SearchByTitle(mongoconn *mongo.Database, collection string, searchBy Article) Article {
	filter := bson.M{"title": searchBy.Title}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

func SearchByTags(mongoconn *mongo.Database, collection string, searchBy Article) Article {
	filter := bson.M{"tags": searchBy.Tags}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

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

func LoadArticle(mongoconn *mongo.Database, collection string, articleData Article) Article {
	// Load by title if article selected
	if articleData.Title != "" {
		filter := bson.M{"title": articleData.Title}
		return atdb.GetOneDoc[Article](mongoconn, collection, filter)
	}
	return atdb.GetOneDoc[Article](mongoconn, collection, nil)
}

func GetOneArticle(mongoconn *mongo.Database, collection string, articleData Article) Article {
	filter := bson.M{"title": articleData.Title}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

func GetByLastDate(mongoconn *mongo.Database, collection string, articleData Article) Article {
	filter := bson.M{"date": articleData.Date}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

func GetByAuthor(mongoconn *mongo.Database, collection string, articleData Article) Article {
	filter := bson.M{"author": articleData.Author}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
}

func UpdateArticle(mongoconn *mongo.Database, collection string, articleData Article) interface{} {
	filter := bson.M{"title": articleData.Title}
	return atdb.ReplaceOneDoc(mongoconn, collection, filter, articleData)
}

func DeleteArticle(mongoconn *mongo.Database, collection string, articleData Article) interface{} {
	filter := bson.M{"title": articleData.Title}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
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
	// Generate Token
	// Create a token for the user
	tokenstring, err := watoken.Encode(userdata.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		return err
	}
	userdata.Token = tokenstring
	// Update the user data in the database
	return atdb.ReplaceOneDoc(mongoconn, collection, bson.M{"username": userdata.Username}, userdata)
}
