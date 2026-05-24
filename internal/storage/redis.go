package storage

import (
	"context"
	"encoding/json"
	"errors"
	"study/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RDB *redis.Client
}

func NewRedis(connStr string) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr: connStr,
	})

	return &Redis{RDB: rdb}
}

func (r *Redis) SaveUser(user models.User) error {
	data, err := json.Marshal(user)

	if err != nil {
		return err
	}

	return r.RDB.Set(context.Background(), user.Name, data, time.Minute*5).Err()
}

func (r *Redis) GetUser(name string) (models.User, error) {
	var user models.User

	data, err := r.RDB.Get(context.Background(), name).Result()

	if err != nil {
		return user, errors.New("user not found in redis")
	}

	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *Redis) DeleteUser(name string) error {

	return r.RDB.Del(context.Background(), name).Err()
}
