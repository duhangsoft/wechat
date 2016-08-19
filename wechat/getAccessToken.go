package wechat

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type getToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

//获取AccessToken
func (client *WeChatClient) getAccessToken() (accesstoken string, err error) {
	if client.tokenTimestamp > time.Now().Unix() {
		accesstoken = client.AccessToken
		return
	}

	reqUrl := WECHATAPIADD + "token?grant_type=client_credential&appid=" +
		client.appId + "&secret=" + client.appSecret
	resp, err := http.Get(reqUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var ok getToken
	err = json.Unmarshal(body, &ok)
	if err != nil {
		return
	}

	if ok.AccessToken == "" {
		err = errors.New("getAccessToken failed ,errcode:" + strconv.Itoa(ok.Errcode) + "   errmsg: " + ok.Errmsg)
		return
	}

	client.AccessToken = ok.AccessToken
	client.tokenTimestamp = ok.ExpiresIn + time.Now().Unix()
	accesstoken = ok.AccessToken
	return
}
