package wechat

import (
	"duhangsoft/wechat/message"
	"net/http"
	"net/url"
)

//接受微信消息
func (client *WeChatClient) messageHandler(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Errorln("get signature failed")
		return
	}
	signature := queryForm["signature"][0]
	timestamp := queryForm["timestamp"][0]
	nonce := queryForm["nonce"][0]
	echostr := queryForm["echostr"]
	if !client.checkSignature(signature, timestamp, nonce) {
		log.Errorln("->s: " + signature + "->t: " + timestamp + "->n:" + nonce)
		return
	}
	if len(echostr) > 0 {
		w.Write([]byte(echostr[0]))
		return
	}
	var msg message.TextMessage

	msg.ReceiveMessage(r)
	log.Debugln(msg)
	msg.FromUser, msg.ToUser = msg.ToUser, msg.FromUser
	log.Debugln("send msg", string(msg.Send()))
	w.Write(msg.Send())
	return
}

func (client *WeChatClient) testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ssss"))
}

func safeHander(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				log.Errorln(e.Error())
			}
		}()
		fn(w, r)
	}
}
