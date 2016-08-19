package message

import (
	"net/http"
)

type Message interface {
	ReceiveMessage(r *http.Request)
}
