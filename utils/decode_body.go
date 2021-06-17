package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func DecodeBody(r *http.Request) io.ReadCloser {
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		return nil
	}

	return ioutil.NopCloser(bytes.NewBuffer(buf))
}
