package service

import (
	"context"
	"fmt"
	"study/internal/models"
	"study/internal/storage"
)

type Service struct {
	storage *storage.Storage
}

func NewService(s *storage.Storage) *Service {
	return &Service{storage: s}
}

func (s Service) CreatUser(ctx context.Context, user models.User) error {
	if user.Name == "" {
		return fmt.Errorf("user.Name is empty")
	}
	if user.Age < 0 {
		return fmt.Errorf("user.Age is negative")
	}

	err := s.storage.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("storage.CreateUser: %w", err)
	}

	return err
}

func (s Service) GetUsers(ctx context.Context) ([]models.User, error) {
	rows, err := s.storage.DB.Query(
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

func (s Service) DeleteUser(ctx context.Context, name string) error {
	_, err := s.storage.DB.Exec(ctx,
		"DELETE FROM users WHERE name=$1",
		name,
	)
	return err
}

func (s Service) UpdateUser(ctx context.Context, user models.User) error {
	_, err := s.storage.DB.Exec(ctx,
		"UPDATE users SET age=$1,WHERE name=$2",
		user.Age,
		user.Name,
	)

	return err
}
