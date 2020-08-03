package card

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

type Service struct {
	db *sql.DB
}

type Card struct {
	Id int64
	Number string
	Balance int64
}

type DbError struct {
	Err error
}

func NewDbError(err error) *DbError {
	return &DbError{Err: err}
}

func (e DbError) Error() string {
	return fmt.Sprintf("db error: %s", e.Err.Error())
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) All(ctx context.Context) ([]*Card, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT id, number, balance FROM cards
		WHERE status = 'ACTIVE'
		LIMIT 50
	`)
	if err != nil {
		if err != pgx.ErrNoRows {
			return nil, NewDbError(err)
		}
		return nil, nil
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Println(cerr)
		}
	}()

	var result []*Card
	for rows.Next() {
		card := &Card{}
		err = rows.Scan(&card.Id, &card.Number, &card.Balance)
		if err != nil {
			return nil, NewDbError(err)
		}
		result = append(result, card)
	}
	err = rows.Err()
	if err != nil {
		return nil, NewDbError(err)
	}
	return result, nil
}

