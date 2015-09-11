package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	endpoint "golang.org/x/oauth2/github"
	// newappengine "google.golang.org/appengine"
	// newurlfetch "google.golang.org/appengine/urlfetch"
	. "./router"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

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
	Scopes:   []string{"user:email", "repo", "openid", "profile"},
	Endpoint: endpoint.Endpoint,
	// RedirectURL: "http://10.14.26.102:8080/view",
}

// var oauthStateSt
var oauthStateString = ""

func oauth2Handler(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println(url)
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + url + "'><button>Login with githbub!</button> </a> </body></html>"))
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/view", HandleGitHubCallback)
	http.HandleFunc("/view/getappdata", HandlerView)
	http.HandleFunc("/view/addnewone", HandlerView)
	http.HandleFunc("/view/getapplist", HandlerView)
	http.HandleFunc("/oauth", Oauth2Handler)
	http.HandleFunc("/test1", MySessionHandler)
	http.HandleFunc("/redirect", RedirectHandler)
	// http.HandleFunc("/deletecookie", RemoveHandler)
	http.ListenAndServe(":8080", nil)
}
