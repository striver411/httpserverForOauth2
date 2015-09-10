package render

import "net/http"

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

func AppInfoView(w http.ResponseWriter, r *http.Request) {
	accountID := "testaccount"
	// checkSession(w, r)

}

// str1, _ := session.Values["Token"].(string)

// 	fmt.Println(str1)
// 	userFromToken, err := checkToken(str1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

func AddAppView(w http.ResponseWriter, r *http.Request) {

}

func UserInfoView(w http.ResponseWriter, r *http.Request) {
}
