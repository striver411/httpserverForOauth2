package render

import (
	"fmt"
	"net/http"
)

func AppInfoView(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookies())

	session, err := store.Get(r, "session-name")
	fmt.Println(session)
	if err != nil {
		// http.Error(w, err.Error(), 500)
		http.Redirect(w, r, "/redirect", http.StatusUnauthorized)
		return
	}
	fmt.Println(session)
	_, ok := session.Values["UserName"]
	if !ok {
		http.Redirect(w, r, "/redirect", http.StatusUnauthorized)
		return
	}
	str, ok := session.Values["UserName"].(string)

	if !ok {
		http.Redirect(w, r, "/redirect", http.StatusUnauthorized)
		return
	}

	str1, _ := session.Values["Token"].(string)

	fmt.Println(str1)
	userFromToken, err := checkToken(str1)
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte("<html><body>You are in </br> " + str + "</br>" + userFromToken + "</br> URL = " + r.URL.String() + "</body></html>"))
}

func AddAppView(w http.ResponseWriter, r *http.Request) {
}

func UserInfoView(w http.ResponseWriter, r *http.Request) {
}
