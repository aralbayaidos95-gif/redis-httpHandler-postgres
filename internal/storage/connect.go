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

func (s *Storage) GetUsers(ctx context.Context) ([]models.User, error) {
	rows, err := s.DB.Query(
		ctx,
		"SELECT name, age FROM users",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		rows.Scan(&user.Name, &user.Age)
		users = append(users, user)
	}

	return users, nil
}

func (s *Storage) DeleteUser(ctx context.Context, name string) error {
	_, err := s.DB.Exec(ctx,
		"DELETE FROM users WHERE name=$1",
		name,
	)
	return err
}

func (s *Storage) UpdateUser(ctx context.Context, user models.User) error {
	_, err := s.DB.Exec(ctx,
		"UPDATE users SET age=$1,WHERE name=$2",
		user.Age,
		user.Name,
	)

	return err
}
