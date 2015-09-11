// Package storage implements a database storage to save user infos.
package storage

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// mongodb configuration
const (
	MONGODB_URL        = "127.0.0.1:27017"
	MONGODB_DB         = "test"
	MONGODB_USER       = "testuser1"
	MONGODB_PWD        = "qwe123123"
	MONGODB_COLLECTION = "deepshare_usercol"
)

// URL/Content struct format
type UserFormat struct {
	// Account_id  int
	Id          string `json:"id" bson:"_id,omitempty"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	GithubName  string `json:"githubname"`
	RealityName string `json:"realityname"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Wechat      string `json:"wechat"`
	QQAccount   string `json:"qqaccount"`
}

// Link2CollectionByDefault 	returns the database server with specfic URL.
func Link2Db(dburl string) (*mgo.Session, error) {
	session, err := mgo.Dial(dburl)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

// Link2CollectionByDefault returns the database server with default URL.
func Link2DbByDefault() (*mgo.Session, error) {
	return Link2Db(MONGODB_URL)
}

// Link2CollectionByDefault returns the collection server with specfic settings.
func Link2Collection(session *mgo.Session, dbname, username, password, collectionname string, auth bool) *mgo.Collection {
	mongoDb := session.DB(dbname)
	if auth {
		mongoDb.Login(username, password)
	}
	return mongoDb.C(collectionname)
}

func EnsureCollection(c *mgo.Collection, key []string, Unique, DropDups, Background, Sparse bool) error {
	return c.EnsureIndex(
		mgo.Index{
			Key:        key,
			Unique:     Unique,
			DropDups:   DropDups,
			Background: Background,
			Sparse:     Sparse,
		},
	)
}

// Link2CollectionByDefault returns the collection server with default settings.
func Link2CollectionByDefault(session *mgo.Session) *mgo.Collection {
	return Link2Collection(session, MONGODB_DB, MONGODB_USER, MONGODB_PWD, MONGODB_COLLECTION, true)
}

// StoreInsert adds a url/content pair into database.
func StoreInsert(c *mgo.Collection, in UserFormat) error {
	err := c.Insert(&in)
	return err
}

func FindMatchUser(c *mgo.Collection, in UserFormat) (string, bool) {
	t := c.Find(in)
	fmt.Println(*t)
	return "", true
}
