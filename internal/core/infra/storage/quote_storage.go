package storage

import (
	"context"
	"database/sql"
	"fmt"
	"quotes/internal/core/domain/quote"
)

type storageImpl struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) quote.Storage {
	return &storageImpl{
		db: db,
	}
}

type QuoteDTO struct {
	ID     int
	Author string
	Quote  string
}

func (dto *QuoteDTO) ToModel() quote.Quote {
	return quote.Quote{
		ID:     dto.ID,
		Author: dto.Author,
		Quote:  dto.Quote,
	}
}

func (storage *storageImpl) Create(ctx context.Context, newQuote quote.NewQuote) error {
	query := `INSERT INTO quotes (author, quote) VALUES ($1, $2)`

	_, err := storage.db.ExecContext(ctx, query, newQuote.Author, newQuote.Quote)
	if err != nil {
		return fmt.Errorf("failed to add new quote %w", err)
	}
	return nil
}

func (storage *storageImpl) Random(ctx context.Context) (quote.Quote, error) {
	var dto QuoteDTO
	query := `SELECT * FROM quotes ORDER BY RANDOM() LIMIT 1`

	row := storage.db.QueryRowContext(ctx, query)
	err := row.Scan(&dto.ID, &dto.Author, &dto.Quote)
	if err != nil {
		return dto.ToModel(), fmt.Errorf("failed to retrieve random quote %w", err)
	}

	return dto.ToModel(), nil
}

func (storage *storageImpl) Get(ctx context.Context, filter quote.Filter) ([]quote.Quote, error) {
	var quotes []quote.Quote
	var result *sql.Rows
	if filter.Author != "" {
		query := `SELECT id, author, quote FROM quotes WHERE author = $1`
		rows, err := storage.db.QueryContext(ctx, query, filter.Author)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve quotes by author: %w", err)
		}
		result = rows
	} else {
		query := `SELECT id, author, quote FROM quotes`
		rows, err := storage.db.QueryContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve all quotes: %w", err)
		}
		result = rows
	}
	defer result.Close()

	for result.Next() {
		var dto QuoteDTO
		if err := result.Scan(&dto.ID, &dto.Author, &dto.Quote); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		quotes = append(quotes, dto.ToModel())
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return quotes, nil
}

func (storage *storageImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM quotes WHERE id = $1`
	_, err := storage.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete by id %d: %w", id, err)
	}
	return nil
}
