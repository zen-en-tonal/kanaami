package patterns

import (
	"errors"
	"net/url"
	"regexp"

	"github.com/zen-en-tonal/kanaami/internal/host"
)

var (
	ErrParse = errors.New("")
)

func Parse(s string) (host.Pattern, error) {
	r, err := regexp.Compile(s)
	if err == nil {
		return NewRegexp(*r), nil
	}
	u, err := url.Parse(s)
	if err == nil {
		return host.From(*u), nil
	}
	return nil, ErrParse
}

// パースに失敗したstringは無視する
func ParsePatrialy(s []string) *host.Patterns {
	var ps host.Patterns
	for _, v := range s {
		p, err := Parse(v)
		if err != nil {
			continue
		}
		ps = append(ps, p)
	}
	return &ps
}
