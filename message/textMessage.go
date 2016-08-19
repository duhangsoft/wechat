package message

import (
	"encoding/xml"
)

type TextMessage struct {
	XMLName    xml.Name `xml:"xml"`
	ToUser     string   `xml:"ToUserName"`
	FromUser   string   `xml:"FromUserName"`
	CreateTime int64    `xml:CreateTime`
	MsgType    string   `xml:"MsgType"`
	Content    string   `xml:"Content"`
}
