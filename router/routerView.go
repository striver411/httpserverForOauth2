package router

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../storage"
)

type AppListInfo struct {
	AppID   string `json:"appid"`
	AppName string `json:"appname"`
}

type AppStats struct {
	Data []int `json:"data"`
}

type AppStatsRequst struct {
	Granularity string //week, day, hour
	Limit       int
}

type AppIdentifer struct {
	AppID string `json:"appid"`
}

type PostRegisterAppObj struct {
	AppName     string `json:"name"`
	FullPkgName string `json:"pkg_name"`
	UrlIos      string `json:"url_ios"`
	UrlAndroid  string `json:"url_android"`
	UrlYYB      string `json:"url_yyb"`
}

func AppInfoViewHandler(w http.ResponseWriter, r *http.Request) {
	accountID, auth := checkSession(w, r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	r.ParseForm()
	appIDString := r.FormValue("appID")
	appID := int64(0)
	if appIDString != "" {
		var err error
		appID, err = strconv.ParseInt(appIDString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
	applist, err := requestApplist(accountID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if appID != 0 || len(applist) != 0 {
		if appID >= int64(len(applist)) {
			http.Error(w, fmt.Errorf("index out of range").Error(), 400)
			return
		}
		fmt.Println(appID, len(applist))
		appstats, err := requestAppInfo(applist[appID].AppID)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(applist)
		fmt.Println(appstats)
	}
	t, err := template.New("index.html").ParseFiles("ds/site/index.html")
	fmt.Println(t, err)
	data := struct {
		BaseURL
	}{
		baseURL,
	}
	t.Execute(w, data)

	// checkSession(w, r)

}

// str1, _ := session.Values["Token"].(string)

// 	fmt.Println(str1)
// 	userFromToken, err := checkToken(str1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

func AddAppViewHandler(w http.ResponseWriter, r *http.Request) {
	accountID, auth := checkSession(w, r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(accountID)
	t, _ := template.New("addapp.html").ParseFiles("ds/site/addapp.html")
	data := struct {
		BaseURL
		AddAppPostURL string
	}{
		baseURL,
		"/post/addnewappop",
	}
	t.Execute(w, data)
}

func AddAppPostHandler(w http.ResponseWriter, r *http.Request) {
	accountID, auth := checkSession(w, r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	r.ParseForm()
	if r.FormValue("fullpkgname") == "" {
		http.Error(w, fmt.Errorf("Full Package name is not specified").Error(), 400)
		return
	}
	res, err := requestRegisterApp(
		accountID,
		PostRegisterAppObj{
			AppName:     r.FormValue("appname"),
			FullPkgName: r.FormValue("fullpkgname"),
			UrlIos:      r.FormValue("urlios"),
			UrlAndroid:  r.FormValue("urlandroid"),
			UrlYYB:      r.FormValue("urlyyb"),
		},
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(res)
}

func UserInfoUpdateViewHandler(w http.ResponseWriter, r *http.Request) {
	accountID, auth := checkSession(w, r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// storage.StoreInsert(storage.UserFormat{Username: accountID})
	r.ParseForm()
	storage.ModifyUser(
		storage.UserFormat{
			Username: accountID,
		},
		storage.UserFormat{
			Password:    r.FormValue("password"),
			GithubName:  r.FormValue("githubname"),
			RealityName: r.FormValue("realityname"),
			Phone:       r.FormValue("phone"),
			Email:       r.FormValue("email"),
			Wechat:      r.FormValue("wechat"),
			QQAccount:   r.FormValue("qqaccount"),
		},
		false,
	)
}

func UserInfoDisplayViewHandler(w http.ResponseWriter, r *http.Request) {
	accountID, auth := checkSession(w, r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	result, err := storage.FindMatchUser(storage.UserFormat{Username: accountID}, true)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(result)
	t, err := template.New("profile.html").ParseFiles("ds/site/profile.html")
	fmt.Println(t, err)
	data := struct {
		BaseURL
		ModifyProfilePostURL string
	}{
		baseURL,
		"/post/modifyuserinfo",
	}
	t.Execute(w, data)
}
