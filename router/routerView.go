package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../storage"
	"gopkg.in/mgo.v2/bson"
)

type AppListInfo struct {
	AppID   string `json:"appid"`
	AppName string `json:"appname"`
}

type AppStats struct {
	Data []struct {
		Open    int `json:"open"`
		Install int `json:"install"`
	} `json:"data"`
	OpenToday      int `json:"open_today"`
	InstallToday   int `json:"install_today"`
	OpenLastDay    int `json:"open_last_day"`
	InstallLastDay int `json:"install_last_day"`
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

type AppInfoReturnType struct {
	AppList []AppListInfo `json:"applist"`
	AppName string        `json:"appname"`
	AppInfo AppStats      `json:"appstats"`
}

func AppInfoViewHandler(w http.ResponseWriter, r *http.Request) {
	_, auth := checkSession(w, r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	t, err := template.New("index.html").ParseFiles("ds/site/index.html")
	fmt.Println(t, err)
	data := struct {
		BaseURL
	}{
		baseURL,
	}
	t.Execute(w, data)
}

func AppInfoRequestHandler(w http.ResponseWriter, r *http.Request) {
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
	resdata := AppInfoReturnType{AppName: "暂无APP", AppList: applist}
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
		resdata.AppInfo = appstats
		resdata.AppName = applist[appID].AppName
		fmt.Println(applist)
		fmt.Println(appstats)
	}
	resjson, err := json.Marshal(resdata)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resjson)
}

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
	if r.FormValue("pkgname") == "" {
		http.Error(w, fmt.Errorf("Full Package name is not specified").Error(), 400)
		return
	}
	urlios := ""
	urlandroid := ""
	urlyyb := ""
	if r.FormValue("apptype") == "1" {
		urlandroid = r.FormValue("link")
	} else {
		urlios = r.FormValue("link")
	}

	_, err := requestRegisterApp(
		accountID,
		PostRegisterAppObj{
			AppName:     r.FormValue("name"),
			FullPkgName: r.FormValue("pkgname"),
			UrlIos:      urlios,
			UrlAndroid:  urlandroid,
			UrlYYB:      urlyyb,
		},
	)
	if err != nil {
		http.Error(w, fmt.Errorf("创建失败").Error(), 400)
		return
	}
	w.Write([]byte("创建成功"))
}

func UserInfoUpdateViewHandler(w http.ResponseWriter, r *http.Request) {
	accountID, auth := checkSession(w, r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	r.ParseForm()
	// fmt.Println(r)
	err := storage.ModifyUser(
		storage.UserFormat{
			Id: bson.ObjectIdHex(accountID),
		},
		storage.UserFormat{
			Password:    r.FormValue("password"),
			GithubName:  r.FormValue("githubname"),
			RealityName: r.FormValue("name"),
			Phone:       r.FormValue("phone"),
			Email:       r.FormValue("email"),
			Wechat:      r.FormValue("wechat"),
			QQAccount:   r.FormValue("qq"),
		},
		false,
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("修改成功"))
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
