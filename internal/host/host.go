package host

import "net/url"

// ホスト名
type Host string

func From(url url.URL) Host {
	return Host(url.Host)
}

func (d Host) Bytes() []byte {
	return []byte(d)
}

func (d Host) Match(other Host) bool {
	return d == other
}

func (d Host) String() string {
	return string(d)
}
