package host

type Pattern interface {
	Match(d Host) bool
	String() string
}

type Patterns []Pattern

func (a Patterns) IsPassed(d Host) bool {
	for _, x := range a {
		if x.Match(d) {
			return true
		}
	}
	return false
}
