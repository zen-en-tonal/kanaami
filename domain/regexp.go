package domain

import "regexp"

type Regexp struct {
	inner regexp.Regexp
}

func NewRegexp(r regexp.Regexp) Regexp {
	return Regexp{r}
}

func (r Regexp) Match(d Domain) bool {
	return r.inner.Match(d.Bytes())
}
