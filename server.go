package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	endpoint "golang.org/x/oauth2/github"
	// newappengine "google.golang.org/appengine"
	// newurlfetch "google.golang.org/appengine/urlfetch"
	. "./router"
	"./storage"
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
	dbSession, err := storage.Link2DbByDefault()
	defer dbSession.Close()
	if err != nil {
		log.Fatal("mongodb error, ", err)
	}
	storage.Link2UserCollectionByDefault(dbSession)

	http.HandleFunc("/", LoginHandler)
	http.HandleFunc("/githubfeedback", GitHubCallbackHandler)
	http.HandleFunc("/index", AppInfoViewHandler)
	http.HandleFunc("/addapp", AddAppViewHandler)
	http.HandleFunc("/profile", UserInfoDisplayViewHandler)

	http.HandleFunc("/post/addnewappop", AddAppPostHandler)
	http.HandleFunc("/post/modifyuserinfo", UserInfoUpdateViewHandler)

	http.HandleFunc("/redirect", RedirectHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("ds/assets"))))

	http.ListenAndServe(":8080", nil)

}
