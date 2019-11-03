package app

import (
	"fmt"
	"log"
	"net/http"
)

func _logHTTPError(
	r *http.Request,
	w *http.ResponseWriter,
	err error,
	code int,
) {
	msg := fmt.Sprintf("%s %s %d - %s",
		r.Method,
		r.RequestURI,
		code,
		err.Error())

	log.Println(msg)
	http.Error((*w), msg, code)

	return
}

func _logRequest(
	r *http.Request,
) {

	msg := fmt.Sprintf("%s %s",
		r.Method,
		r.RequestURI)

	log.Println(msg)

}
