package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func checkSession(w http.ResponseWriter, r *http.Request) string {
	session, err := store.Get(r, "session-name")
	fmt.Println(session)
	if err != nil {
		// http.Error(w, err.Error(), 500)
		http.Redirect(w, r, "/redirect", http.StatusUnauthorized)
		return ""
	}
	fmt.Println(session)
	_, ok := session.Values["AccountID"]
	if !ok {
		http.Redirect(w, r, "/redirect", http.StatusUnauthorized)
		return ""
	}
	str, ok := session.Values["AccountID"].(string)

	if !ok {
		http.Redirect(w, r, "/redirect", http.StatusUnauthorized)
		return ""
	}
	return str
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

func requestAppInfo(appID string) (AppStats, error) {
	u, _ := url.Parse("http://deepshare.chinacloudapp.cn:8080/appstats/" + appID)
	q := u.Query()
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return AppStats{}, err
	}
	if res.StatusCode != 200 {
		return AppStats{}, fmt.Errorf("Failed request, error Msg = %v", res.Status)
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

func requestRegisterApp(accountID string, appInfo PostRegisterAppObj) (AppIdentifer, error) {
	u, _ := url.Parse("http://deepshare.chinacloudapp.cn:8080/apps/" + accountID)
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
