package patterns

import (
	"regexp"

	"github.com/zen-en-tonal/kanaami/internal/host"
)

type Regexp struct {
	inner regexp.Regexp
}

func NewRegexp(r regexp.Regexp) Regexp {
	return Regexp{r}
}

func (r Regexp) Match(d host.Host) bool {
	return r.inner.Match(d.Bytes())
}

func (r Regexp) String() string {
	return r.inner.String()
}
