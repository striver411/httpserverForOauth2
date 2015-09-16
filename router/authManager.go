package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../storage"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2/bson"
)

func GetAccountID(githubname, token string) (string, error) {
	res, err := storage.FindMatchUser(storage.UserFormat{GithubName: githubname}, false)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return "", err
	}
	if len(res) > 0 && (err == nil || !strings.Contains(err.Error(), "not found")) {
		err = storage.ModifyUser(storage.UserFormat{GithubName: githubname}, storage.UserFormat{Token: token}, false)
		if err != nil {
			return "", err
		}
		return res[0].Id.Hex(), nil
	}
	err = storage.StoreInsert(storage.UserFormat{GithubName: githubname, Token: token})
	if err != nil {
		return "", err
	}
	fmt.Println(strings.Contains(err.Error(), "not found"))
	res, err = storage.FindMatchUser(storage.UserFormat{GithubName: githubname}, false)

	if err != nil {
		return "", err
	}
	return res[0].Id.Hex(), nil
}

func AuthUser(username, password string) (string, bool, error) {
	res, err := storage.FindMatchUser(storage.UserFormat{Username: username}, false)
	if err != nil {
		return "", false, err
	}
	if len(res) == 0 || res[0].Password != password {
		return "", false, nil
	}
	return res[0].Id.Hex(), true, nil
}

func CheckAccountIDExist(accountID string) (bool, error) {
	res, err := storage.FindMatchUser(storage.UserFormat{Id: bson.ObjectIdHex(accountID)}, false)
	if err != nil || len(res) == 0 || res[0].Id.Hex() != accountID {
		return false, err

	}
	return true, nil
}

func CheckToken(jsonToken string) (string, error) {
	var token oauth2.Token
	fmt.Println("jsonToken:", string(jsonToken))
	err := json.Unmarshal([]byte(jsonToken), &token)
	if err != nil {
		return "", err
	}
	oauthClient := conf.Client(oauth2.NoContext, &token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")

	if err != nil {
		return "", err
	}
	return *user.Login, nil
}

func checkSession(w http.ResponseWriter, r *http.Request) (string, bool) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		// http.Error(w, err.Error(), 500)
		return "", false
	}
	fmt.Println(session)
	_, ok := session.Values["AccountID"]
	if !ok {
		return "", false
	}
	str, ok := session.Values["AccountID"].(string)

	if !ok {
		return "", false
	}
	return str, true
}
