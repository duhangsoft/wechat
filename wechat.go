package wechat

import (
	"net/http"
)

const (
	APPNAME  = "duhangsoft"
	APPID    = "wx8fa4fd5bf3e8f928"
	APPSECRT = "acfbff71914c751ee0feeb43e20c4935"
	TOKEN    = "duhangyangguang"
)

func Start() {
	client := NewWeChatClient(APPNAME, APPID, APPSECRT, TOKEN)
	log.Debugln("app secret: ", client.appSecret)
	_, err := client.getAccessToken()
	if err != nil {
		log.Errorln(err)
		return
	}
	//return err
	Path := "/" + client.appName + "/"
	log.Debugln("weChat API: YOUR ADD " + Path)

	mux := http.NewServeMux()
	mux.HandleFunc(Path, safeHander(client.messageHandler))
	mux.HandleFunc(Path+"test/", safeHander(client.testHandler))
	err = http.ListenAndServe(":80", mux)
	log.Errorln(err)
}
