package domain

type Domain string

func New(s string) Domain {
	return Domain(s)
}

func (d Domain) Bytes() []byte {
	return []byte(d)
}

func (d Domain) Match(other Domain) bool {
	return d == other
}

type Match interface {
	Match(d Domain) bool
}

type Allows []Match

func (a Allows) IsAccepted(d Domain) bool {
	for _, x := range a {
		if x.Match(d) {
			return true
		}
	}
	return false
}
