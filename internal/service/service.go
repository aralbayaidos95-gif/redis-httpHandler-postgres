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
	users, err := s.storage.GetUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("storage.GetUsers: %w", err)
	}

	return users, nil
}

func (s Service) DeleteUser(ctx context.Context, name string) error {

	if name == "" {
		return fmt.Errorf("user.Name is empty")
	}

	err := s.storage.DeleteUser(ctx, name)

	if err != nil {
		return fmt.Errorf("storage.DeleteUser: %w", err)
	}

	return err
}

func (s Service) UpdateUser(ctx context.Context, user models.User) error {
	if user.Name == "" {
		return fmt.Errorf("user.Name is empty")
	}
	if user.Age < 0 {
		return fmt.Errorf("user.Age is negative")
	}

	err := s.storage.UpdateUser(ctx, user)

	if err != nil {
		return fmt.Errorf("storage.UpdateUser: %w", err)
	}

	return err
}
