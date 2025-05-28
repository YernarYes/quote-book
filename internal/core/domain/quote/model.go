package quote

type Quote struct {
	ID     int
	Author string
	Quote  string
}

type NewQuote struct {
	Author string
	Quote  string
}

type Filter struct {
	ID     int
	Author string
}
