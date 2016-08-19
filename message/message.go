package message

import (
	"encoding/xml"
	"net/http"
)

type Message interface {
	//Send() []byte
}

type RMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUser       string   `xml:"ToUserName"`
	FromUser     string   `xml:"FromUserName"`
	CreateTime   int64    `xml:CreateTime`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int64    `xml:MsgId`
	PicUrl       string   `xml:"PicUrl"`
	MediaId      string   `xml:"MediaId"`
	Format       string   `xml:"Format"`
	Recognition  string   `xml:"Recognition"`
	ThumbMediaId string   `xml:"ThumbMediaId"`
	Location_X   float64  `xml:Location_X`
	Location_Y   float64  `xml:Location_Y`
	Scale        int32    `xml:Scale`
	Label        string   `xml:"Label"`
	Title        string   `xml:"Title"`
	Description  string   `xml:"Description"`
	Url          string   `xml:"Url"`
	Event        string   `xml:"Event"`
	EventKey     string   `xml:"EventKey"`
	Ticket       string   `xml:"Ticket"`
	Latitude     float64  `xml:Latitude`
	Longitude    float64  `xml:Longitude`
	Precision    float64  `xml:Precision`
}

func (msg *RMessage) ReceiveMessage(r *http.Request) {
	m := make([]byte, 1024)
	_, err := r.Body.Read(m)

	err = xml.Unmarshal(m, msg)
	if err != nil {
		log.Errorln("receiveMessage Error: " + err.Error())
		return
	}

}

//自动回复逻辑
func (msg *RMessage) GetReplyMsg() Message {
	return &TextMessage{
		FromUser:   msg.ToUser,
		ToUser:     msg.FromUser,
		CreateTime: msg.CreateTime,
		MsgType:    "text",
		Content:    "建设中。。。",
	}
}

func Send(msg Message) []byte {
	m := make([]byte, 1024)
	m, err := xml.Marshal(msg)
	if err != nil {
		log.Errorln("make messag error:", err)
	}
	return m
}
