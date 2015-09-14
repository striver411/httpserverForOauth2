// Package router implements a URL handler system, applying exclusive
// handler to deal with user URL request.
package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	endpoint "golang.org/x/oauth2/github"
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	urlGithub := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println(urlGithub)
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + urlGithub + "'><button>Login with githbub!</button> </a> </body></html>"))
}

func GitHubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	authcode := r.FormValue("code")
	token, err := conf.Exchange(oauth2.NoContext, authcode)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println("token:")
	fmt.Println(token)

	// client := conf.Client(oauth2.NoContext, tok)
	oauthClient := conf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
	// fmt.Println(token.Extra("id_token"), token.AccessToken)
	jsonToken, err := json.Marshal(*token)
	if err != nil {
		panic(err.Error())
	}
	accountID, err := GetAccountID(*user.Login, string(jsonToken))
	fmt.Println("getAccountID : \n", accountID, err)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	session.Values["AccountID"] = accountID
	err = session.Save(r, w)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/view/getappdata", http.StatusTemporaryRedirect)
	// http.SetCookie(w, &cookies)
}

func MainViewHandle(w http.ResponseWriter, r *http.Request) {
	url1 := "/view/a"
	url2 := "/view/b"
	url3 := "/view/c"
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + url1 + "'><button>url1</button> </a> <a href='" + url2 + "'><button>url2</button> </a><a href='" + url3 + "'><button>url3</button> </a></body></html>"))
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
