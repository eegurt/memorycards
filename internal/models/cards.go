package models

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Card struct {
	CardSetID int
	Question  string
	Answer    string
}

type CardModel struct {
	DB *pgxpool.Pool
}

func (m *CardModel) Insert(ctx context.Context, cardSetID int, question string, answer string) error {
	query := `INSERT INTO cards (card_set_id, question, answer)
	VALUES($1, $2, $3)`

	err := m.DB.QueryRow(ctx, query, cardSetID, question, answer).Scan()
	if err != nil {
		return err
	}

	return nil
}
