package quote

import (
	"context"
	"fmt"
	"log/slog"
)

type Service interface {
	Create(ctx context.Context, newQuote NewQuote) error
	Get(ctx context.Context, filter Filter) ([]Quote, error)
	Delete(ctx context.Context, filter Filter) error
	Random(ctx context.Context) (Quote, error)
}

type serviceImpl struct {
	log     *slog.Logger
	storage Storage
}

func NewService(log *slog.Logger, storage Storage) *serviceImpl {
	return &serviceImpl{
		log:     log,
		storage: storage,
	}
}

func (s *serviceImpl) Create(ctx context.Context, newQuote NewQuote) error {
	err := s.storage.Create(ctx, newQuote)
	if err != nil {
		s.log.Error("failed to create quote: ", "err", err)
		return fmt.Errorf("failed to create %w", err)
	}
	return nil
}

func (s *serviceImpl) Random(ctx context.Context) (Quote, error) {
	randomQuote, err := s.storage.Random(ctx)
	if err != nil {
		s.log.Error("failed to get random quote", "err", err)
		return Quote{}, fmt.Errorf("failed to retrieve random quote %w", err)
	}
	return randomQuote, nil
}

func (s *serviceImpl) Get(ctx context.Context, filter Filter) ([]Quote, error) {
	getQuote, err := s.storage.Get(ctx, filter)
	if err != nil {
		s.log.Error("failed to get random quote", "err", err)
		return nil, fmt.Errorf("failed to retrieve random quote %w", err)
	}
	return getQuote, nil

}

func (s *serviceImpl) Delete(ctx context.Context, filter Filter) error {
	err := s.storage.Delete(ctx, filter.ID)
	if err != nil {
		s.log.Error("failed to delete quote", "err", err)
		return fmt.Errorf("failed to delete quote %w", err)
	}
	return nil
}
