package http_utils

import (
	"io/ioutil"
	"net/http"
)

func GetRequestBody(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return body
}
