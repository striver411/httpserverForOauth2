package storage

import (
	"crypto/rand"
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func randString(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("negative number")
	}
	const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes), nil
}

func TestStore2Mongo(t *testing.T) {
	dbSession, err := Link2DbByDefault()
	defer dbSession.Close()
	if err != nil {
		t.Errorf("db test failed, error Msg = %v", err)
		return
	}

	// UserName should be unique
	var storageTest = []struct {
		Username   string
		Password   string
		GithubName string
	}{{"a", "b", "c"}}

	testCol, _ := randString(10)
	// testCol := "fortest"

	fmt.Println(testCol)
	c := Link2Collection(dbSession, MONGODB_DB, MONGODB_USER, MONGODB_PWD, testCol, true)
	defer c.DropCollection()
	err = EnsureCollection(c, []string{"username", "githubname"}, true, true, true, true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	for _, i := range storageTest {
		insertSet := UserFormat{
			// Account_id : GetNextSeq("userid")
			Username:   i.Username,
			Password:   i.Password,
			GithubName: i.GithubName,
		}
		fmt.Println(insertSet)
		err := StoreInsert(c, insertSet)
		if err != nil {
			t.Errorf("%v", err)
		}
	}

	result := UserFormat{}
	for _, i := range storageTest {
		err = c.Find(bson.M{"username": i.Username}).One(&result)
		if i.GithubName != result.GithubName || i.Password != result.Password {
			t.Errorf("db store mismatch, when store %s", i.Username)
		}
		fmt.Println(result)
	}

	if err != nil {
		t.Errorf("db test failed, error Msg = %v", err)
		return
	}
}
