package http

import (
	"net/http"
	"net/url"

	"github.com/zen-en-tonal/kanaami/internal/host"
)

type GateWay struct {
	CreateAllows
}

type CreateAllows func(*http.Request) (*host.Patterns, error)

func (f GateWay) AsMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pattern, err := f.CreateAllows(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ref, err := url.Parse(r.Referer())
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if !pattern.IsPassed(host.From(*ref)) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
