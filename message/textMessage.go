package message

import (
	"encoding/xml"
	"net/http"
)

type TextMessage struct {
	XMLName    xml.Name `xml:"xml"`
	ToUser     string   `xml:"ToUserName"`
	FromUser   string   `xml:"FromUserName"`
	CreateTime int64    `xml:CreateTime`
	MsgType    string   `xml:"MsgType"`
	Content    string   `xml:"Content"`
	MsgId      int64    `xml:MsgId`
}

func (msg *TextMessage) ReceiveMessage(r *http.Request) {
	m := make([]byte, 1024)
	_, err := r.Body.Read(m)
	/*
		if err != nil || err.Error() != "EOF" {
			log.Errorln("r.Body.Read err:", err, "m:", m)
			return
		}
	*/
	err = xml.Unmarshal(m, msg)
	if err != nil {
		log.Errorln("receiveMessage Error: " + err.Error())
		return
	}

}
func (msg *TextMessage) Send() []byte {
	m := make([]byte, 1024)
	m, err := xml.Marshal(msg)
	if err != nil {
		log.Errorln("make messag error:", err)
	}
	return m
}
