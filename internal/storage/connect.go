package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB *pgx.Conn
}

func NewStorage(connStr string) (*Storage, error) {
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return nil, err
	}
	return &Storage{DB: conn}, nil
}
