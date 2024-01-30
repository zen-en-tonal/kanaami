package http

import (
	"fmt"
	"net/http"
)

type ContentSrc func(r *http.Request) string

func (c ContentSrc) Handle(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(buildIframe(c(req))))
}

func buildIframe(src string) string {
	return fmt.Sprintf("<iframe src='%s' />", src)
}
