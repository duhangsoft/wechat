package wechat

import (
	"duhangsoft/wechat/message"
	"net/http"
)

//接受微信消息
func (client *WeChatClient) messageHandler(w http.ResponseWriter, r *http.Request) {

	echostr, ok := client.checkSignature(r)
	if !ok {
		log.Errorln("check mesage From error!", r.Host)
		return
	}
	if echostr != "" {
		w.Write([]byte(echostr))
		return
	}
	var msg message.RMessage

	//得到消息
	msg.ReceiveMessage(r)
	log.Debugln(msg)

	//对消息进行处理得到要回复的消息
	replyMsg := msg.GetReplyMsg()
	//对要回复的消息做点啥－比如记录

	//回复消息
	w.Write(message.Send(replyMsg))
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
