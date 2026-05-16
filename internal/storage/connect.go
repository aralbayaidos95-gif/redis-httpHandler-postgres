package storage

import (
	"context"
	"fmt"
	"study/internal/models"

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

func (s *Storage) CreateUser(ctx context.Context, user models.User) error {
	_, err := s.DB.Exec(
		ctx,
		"INSERT INTO users (name, age) VALUES ($1,$2) ",
		user.Name,
		user.Age,
	)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}
