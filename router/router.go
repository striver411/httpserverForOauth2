// Package render implements a URL handler system, applying exclusive
// handler to deal with user URL request.
package render

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Page struct {
	Title string
	Body  []byte
}

var conf = &oauth2.Config{
	ClientID:     "9487562b91cf0e58a7f5",
	ClientSecret: "e7de8b20bdc18a0d4c221a319ef1a585b3c187a4",
	// Scopes:       []string{},
	Scopes:   []string{"user:email", "repo", "openid", "profile"},
	Endpoint: endpoint.Endpoint,
	// RedirectURL: "http://10.14.26.102:8080/view",
}

func Oauth2Handler(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println(url)
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + url + "'><button>Login with githbub!</button> </a> </body></html>"))
}

func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	authcode := r.FormValue("code")
	fmt.Println(authcode)
	token, err := conf.Exchange(oauth2.NoContext, authcode)

	if err != nil {
		fmt.Println("token get error")
		log.Fatal(err)
	}
	fmt.Println("token:")
	fmt.Println(token)

	// client := conf.Client(oauth2.NoContext, tok)
	oauthClient := conf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	fmt.Println("user=======\n", user)
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		http.Redirect(w, r, "/redirect", http.StatusUnauthorized)
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
	// fmt.Println(token.Extra("id_token"), token.AccessToken)
	url1 := "/view/a"
	url2 := "/view/b"
	url3 := "/view/c"

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println(session)
		return
	}
	jsonToken, err := json.Marshal(*token)
	fmt.Println("jsonToken:", string(jsonToken))
	if err != nil {
		panic(err.Error())
	}

	session.Values["UserName"] = *user.Login
	session.Values["Token"] = string(jsonToken)
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println(session)
		return
	}

	fmt.Println(session)
	fmt.Println(w)
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + url1 + "'><button>url1</button> </a> <a href='" + url2 + "'><button>url2</button> </a><a href='" + url3 + "'><button>url3</button> </a></body></html>"))
	// http.SetCookie(w, &cookies)
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

func MySessionHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	fmt.Println(r.Cookies())
	session, _ := store.Get(r, "session-name")
	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
	fmt.Println(w)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Auth failed, Redirect"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookies())
	_, err := r.Cookie("session-name")
	if err != nil {
		w.Write([]byte("no such cookie"))
		return
	}
	cookieMonster := &http.Cookie{
		Name:   "session-name",
		MaxAge: -1,
	}
	http.SetCookie(w, cookieMonster)
	http.Redirect(w, r, "/", http.StatusUnauthorized)
}