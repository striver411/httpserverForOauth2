package main

import (
	"net/http"

	. "./render"
)

// newappengine "google.golang.org/appengine"
// newurlfetch "google.golang.org/appengine/urlfetch"

func main() {
	http.HandleFunc("/", Oauth2Handler)
	http.HandleFunc("/view", HandleGitHubCallback)
	http.HandleFunc("/view/getappdata", AppInfoView)
	http.HandleFunc("/view/addnewapp", AddAppView)
	http.HandleFunc("/view/supplementuserinfo", UserInfoView)
	http.HandleFunc("/test1", MySessionHandler)
	http.HandleFunc("/redirect", RedirectHandler)
	http.HandleFunc("/deletecookie", RemoveHandler)
	http.ListenAndServe(":8080", nil)
}
