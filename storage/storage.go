// Package storage implements a database storage to save user infos.
package storage

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// mongodb configuration
const (
	MONGODB_URL        = "127.0.0.1:27017"
	MONGODB_DB         = "test"
	MONGODB_USER       = "testuser1"
	MONGODB_PWD        = "qwe123123"
	MONGODB_COLLECTION = "deepshareusercol"
)

// URL/Content struct format
type UserFormat struct {
	// Account_id  int
	Id          string `json:"id" bson:"_id,omitempty"`
	Username    string `json:"username" bson:"username,omitempty"`
	Password    string `json:"password" bson:"password,omitempty"`
	GithubName  string `json:"githubname" bson:"githubname,omitempty"`
	RealityName string `json:"realityname" bson:"realityname,omitempty"`
	Phone       string `json:"phone" bson:"phone,omitempty"`
	Email       string `json:"email" bson:"email,omitempty"`
	Wechat      string `json:"wechat" bson:"wechat,omitempty"`
	QQAccount   string `json:"qqaccount" bson:"qqaccount,omitempty"`
	Token       string `json:"token" bson:"token,omitempty"`
}

var UserCol *mgo.Collection

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
func Link2UserCollection(session *mgo.Session, dbname, username, password, collectionname string, auth bool) *mgo.Collection {
	mongoDb := session.DB(dbname)
	if auth {
		mongoDb.Login(username, password)
	}
	UserCol = mongoDb.C(collectionname)
	return UserCol
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
func Link2UserCollectionByDefault(session *mgo.Session) *mgo.Collection {
	return Link2UserCollection(session, MONGODB_DB, MONGODB_USER, MONGODB_PWD, MONGODB_COLLECTION, true)
}

// StoreInsert adds a url/content pair into database.
func StoreInsert(in UserFormat) error {
	err := UserCol.Insert(&in)
	return err
}

func FindMatchUser(matchPattern UserFormat, all bool) ([]UserFormat, error) {
	result := []UserFormat{}
	if UserCol == nil {
		return result, fmt.Errorf("Collection not prepare!")
	}

	t := UserCol.Find(matchPattern)
	var err error
	if all {
		err = t.All(&result)
	} else {
		one := UserFormat{}
		err = t.One(&one)
		result = append(result, one)
	}
	return result, err
}

func ModifyUser(matchPattern UserFormat, update UserFormat, all bool) error {
	if all {
		_, err := UserCol.UpdateAll(matchPattern, bson.M{"$set": update})
		return err
	}
	err := UserCol.Update(matchPattern, bson.M{"$set": update})
	return err

}
