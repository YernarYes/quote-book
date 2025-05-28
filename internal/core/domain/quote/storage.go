package quote

import "context"

type Storage interface {
	Create(ctx context.Context, newQuote NewQuote) error
	Random(ctx context.Context) (Quote, error)
	Get(ctx context.Context, filter Filter) ([]Quote, error)
	Delete(ctx context.Context, id int) error
}
