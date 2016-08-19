package wechat

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	WECHATAPIADD = "https://api.weixin.qq.com/cgi-bin/"
)

var WeChat WeChatClient

type WeChatClient struct {
	appName        string
	appId          string
	appSecret      string
	token          string
	AccessToken    string
	tokenTimestamp int64
}

func NewWeChatClient(appName, appId, appSecret, token string) *WeChatClient {
	return &WeChatClient{
		appName:   appName,
		token:     token,
		appId:     appId,
		appSecret: appSecret,
	}

}

//验证消息来源
func (client *WeChatClient) checkSignature(r *http.Request) (echostr string, ok bool) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Errorln("get signature failed")
		return
	}
	signature := strings.Join(queryForm["signature"], "")
	timestamp := strings.Join(queryForm["timestamp"], "")
	nonce := strings.Join(queryForm["nonce"], "")
	echostr = strings.Join(queryForm["echostr"], "")

	str := []string{client.token, timestamp, nonce}
	sort.Strings(str)
	temStr := strings.Join(str, "")
	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(temStr))
	if signature == fmt.Sprintf("%x", sha1Inst.Sum(nil)) {
		ok = true
		return
	}
	return

}
