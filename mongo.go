package berkatbepkg

import (
	"context"
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

func FindUser(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func IsUsernameExists(MONGOSTRING, dbname string, datauser User) bool {
	mconn := SetConnection(MONGOSTRING, dbname).Collection("userLogin")
	filter := bson.M{"username": datauser.Username}

	var userdata User
	err := mconn.FindOne(context.Background(), filter).Decode(&userdata)
	return err == nil
}

func IsPasswordValid(MONGOSTRING *mongo.Database, collection string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](MONGOSTRING, collection, filter)
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

func LoadByArticleID(mongoconn *mongo.Database, collection string, articleData Article) Article {
	filter := bson.M{"article_id": articleData.ID}
	return atdb.GetOneDoc[Article](mongoconn, collection, filter)
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

func CreateUserAndAddedToken(MONGOCONNSTRINGEV *mongo.Database, collection string, datauser User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(datauser.Password)
	if err != nil {
		return err
	}
	datauser.Password = hashedPassword
	return atdb.InsertOneDoc(MONGOCONNSTRINGEV, collection, datauser)
}
