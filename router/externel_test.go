package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AppListInfo struct {
	AppID   string `json:"appid"`
	AppName string `json:"appname"`
}

type PostRegisterAppObj struct {
	AppName     string `json:"name"`
	FullPkgName string `json:"pkg_name"`
	UrlIos      string `json:"url_ios"`
	UrlAndroid  string `json:"url_android"`
	UrlYYB      string `json:"url_yyb"`
}

type AppIdentifer struct {
	AppID string `json:"appid"`
}

type appListInfo struct {
	AppID   string `json:"appid"`
	AppName string `json:"appname"`
}

type AppStats struct {
	Data []int `json:"data"`
}

func requestRegisterApp(appID string, appInfo PostRegisterAppObj) (AppIdentifer, error) {
	u, _ := url.Parse("http://deepshare.chinacloudapp.cn:8080/apps/" + appID)
	q := u.Query()
	u.RawQuery = q.Encode()
	jsonStr, err := json.Marshal(appInfo)
	res, err := http.Post(u.String(), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return AppIdentifer{}, err
	}
	if res.StatusCode != 200 {
		return AppIdentifer{}, fmt.Errorf("Failed request, error Msg = %v", res.Status)
	}
	// receive data
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return AppIdentifer{}, err
	}

	fmt.Println(res)
	fmt.Println(string(result))

	// covert data to pointed sturct
	var appIden AppIdentifer
	err = json.Unmarshal(result, &appIden)
	return appIden, nil

}

func requestAppInfo(appID string) (AppStats, error) {
	u, _ := url.Parse("http://deepshare.chinacloudapp.cn:8080/appstats/" + appID)
	q := u.Query()
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return AppStats{}, err
	}

	// receive data
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return AppStats{}, err
	}

	// covert data to pointed sturct
	var appInfo AppStats
	err = json.Unmarshal(result, &appInfo)
	return appInfo, nil
}

func requestApplist(accountID string) ([]AppListInfo, error) {

	u, _ := url.Parse("http://deepshare.chinacloudapp.cn:8080/apps/" + accountID)
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	fmt.Println(res)
	if err != nil {
		return []AppListInfo{}, err
	}

	// receive data
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return []AppListInfo{}, err
	}

	// covert data to pointed sturct
	var appList []AppListInfo
	err = json.Unmarshal(result, &appList)
	return appList, nil
}

func main() {

	// u, _ := url.Parse("http://deepshare.chinacloudapp.cn:8080/apps/testaccount")
	// u, _ := url.Parse("http://deepshare.chinacloudapp.cn:8080/apps/www")
	// q := u.Query()
	// q.Set("username", "user")
	// q.Set("password", "passwd")
	// u.RawQuery = q.Encode()
	// res, err := http.Get(u.String())
	// fmt.Println(res)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// result, err := ioutil.ReadAll(res.Body)
	// var testV []appListInfo
	// err = json.Unmarshal(result, &testV)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Printf("%s", result)

	b := PostRegisterAppObj{
		"a", "bb2", "c", "d", "e",
	}
	a, err := requestRegisterApp("www", b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	c, err := requestAppInfo(a.AppID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	d, err := requestApplist("www")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}
