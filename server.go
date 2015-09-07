package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

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
	RedirectURL: "http://10.14.26.102:8080/view",
}

// var oauthStateSt
var oauthStateString = ""

func oauth2Handler(w http.ResponseWriter, r *http.Request) {

	// first try
	// url := conf.AuthCodeURL("")

	url := conf.AuthCodeURL("mycode")
	fmt.Println(url)
	// Handle the exchange code to initiate a transport.
	// tok, err := conf.Exchange(oauth2.NoContext, "authorization-code")
	// fmt.Println(tok)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := conf.Client(oauth2.NoContext, tok)
	// client.Get("...")

	// fmt.Printf("Visit the URL for the auth dialog: %v", url)
	//redirect user to that page
	// http.Redirect(url)
	// http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	w.Write([]byte("<html><title>Golang Login github Example</title> <body> <a href='" + url + "'><button>Login with githbub!</button> </a> </body></html>"))
	// // second try
	// 	// Redirect user to consent page to ask for permission
	// 	// for the scopes specified above.
	// 	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	// 	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// 	// Use the authorization code that is pushed to the redirect URL.
	// 	// NewTransportWithCode will do the handshake to retrieve
	// 	// an access token and initiate a Transport that is
	// 	// authorized and authenticated by the retrieved token.
	// 	var code string
	// 	code = "https://github.com/login/oauth/authorize"
	// 	// _, err := fmt.Scan(&code)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	tok, err := conf.Exchange(oauth2.NoContext, code)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	client := conf.Client(oauth2.NoContext, tok)
	// 	client.Get("...")

	// var code string
	// if _, err := fmt.Scan(&code); err != nil {
	// 	log.Fatal(err)
	// }
	// tok, err := conf.Exchange(oauth2.NoContext, code)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := conf.Client(oauth2.NoContext, tok)
	// client.Get("...")
}

var userInfoTemplate = template.Must(template.New("").Parse(`
<html><body>
This app is now authenticated to access your github user info.  Your details are:<br />
{{.}}
</body></html>
`))

func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	// // Get the code from the response
	// code := r.FormValue("code")

	// t := &oauth2.Transport{oauth.Config: config}
	// conf.TokenSource(ctx, t)

	// // Exchange the received code for a token
	// t.Exchange(code)

	// //now get user data based on the Transport which has the token
	// resp, _ := t.Client().Get("http://10.14.26.102:8080/view")

	// buf := make([]byte, 1024)
	// resp.Body.Read(buf)
	// userInfoTemplate.Execute(w, string(buf))
}

const htmlIndex = `<html><body>
Logged in with <a href="/login">GitHub</a>
</body></html>
`

// /
func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "333", "444")
}

func handleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("1!!!")
	// fmt.Println(r)
	fmt.Println(r.Form)
	state := r.FormValue("state")
	fmt.Println(state)
	// if state != oauthStateString {
	// 	fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }

	// // code := r.URL.Query().Get("code")

	// // Exchanging the code for a token
	// token, err := conf.Exchange(oauth2.NoContext, "authorization-code")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Println(token)
	code := r.URL.Query().Get("code")
	fmt.Println(code)
	authcode := r.FormValue("code")
	fmt.Println(authcode)
	token, err := conf.Exchange(oauth2.NoContext, authcode)
	if err != nil {
		fmt.Println("11111")
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
	fmt.Println(token.Extra("id_token"), token.AccessToken)
	// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	// session, _ := app.GlobalSessions.SessionStart(w, r)
	// defer session.SessionRelease(w)

	// session.Set("id_token", token.Extra("id_token"))
	// session.Set("access_token", token.AccessToken)
	// session.Set("profile", profile)

	// // Redirect to logged in page
	// http.Redirect(w, r, "/user", http.StatusMovedPermanently)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", handleGitHubCallback)
	http.HandleFunc("/view1", viewHandler)
	http.HandleFunc("/oauth", oauth2Handler)
	// http.Post(url, bodyType, body)
	http.ListenAndServe(":8080", nil)
}
