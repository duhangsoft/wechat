package route

import (
	"net/http"
)

func SafeHander(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				logger(e.Error())
			}
		}()
		fn(w, r)
	}
}
