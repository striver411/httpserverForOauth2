package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	endpoint "golang.org/x/oauth2/github"
	// newappengine "google.golang.org/appengine"
	// newurlfetch "google.golang.org/appengine/urlfetch"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type Page struct {
	Title string
	Body  []byte
}

var conf = &oauth2.Config{
	ClientID:     "9487562b91cf0e58a7f5",
	ClientSecret: "e7de8b20bdc18a0d4c221a319ef1a585b3c187a4",
	// Scopes:       []string{},
	// Scopes:      []string{"SCOPE1", "SCOPE2"},
	Scopes:      []string{"user:email", "repo", "openid", "profile"},
	Endpoint:    endpoint.Endpoint,
	// RedirectURL: "http://10.14.26.102:8080/view",
}

// var oauthStateSt
var oauthStateString = ""

func oauth2Handler(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println(url)
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + url + "'><button>Login with githbub!</button> </a> </body></html>"))
}


func handleGitHubCallback(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
	fmt.Println("Name:", *user.Email)
	fmt.Println(token.Extra("id_token"), token.AccessToken)
	url1 := "/view/a"
	url2 := "/view/b"
	url3 := "/view/c"
	
	cookie := http.Cookie{ Name: "Username", Value: *user.Email,  Expires: time.Now().Add(time.Hour), HttpOnly: true}
	http.SetCookie(w, &cookie)

	cookie = http.Cookie{ Name: "Token", Value: "",  Expires: time.Now().Add(time.Hour), HttpOnly: true}
	http.SetCookie(w, &cookie)

	cookie = http.Cookie{ Name: "Auth", Value: "",  Expires: time.Now().Add(time.Hour), HttpOnly: true}
	http.SetCookie(w, &cookie)
	fmt.Println(w)
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + url1 + "'><button>url1</button> </a> <a href='" + url2 + "'><button>url2</button> </a><a href='" + url3 + "'><button>url3</button> </a></body></html>"))
	// http.SetCookie(w, &cookies)
}

func handlerView(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookies())
	w.Write([]byte("<html><body>You are in " + r.URL.String() + "</body></html>"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view", handleGitHubCallback)
	http.HandleFunc("/view/a", handlerView)
	http.HandleFunc("/view/b", handlerView)
	http.HandleFunc("/view/c", handlerView)
	http.HandleFunc("/oauth", oauth2Handler)
	http.ListenAndServe(":8080", nil)
}
