package handler

import "quotes/internal/core/domain/quote"

type View struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func (v *View) ToModel() quote.NewQuote {
	return quote.NewQuote{
		Author: v.Author,
		Quote:  v.Quote,
	}
}
