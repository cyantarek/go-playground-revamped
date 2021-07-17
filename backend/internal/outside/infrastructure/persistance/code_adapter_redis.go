package persistance

import (
	"backend/internal/core/domain"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
)

type CodeDBRedis struct {
	client *redis.Client
}

func (cr *CodeDBRedis) Save(ctx context.Context, c domain.Code) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = cr.client.Set(c.ShortCode().String(), b, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cr *CodeDBRedis) Get() error {
	panic("implement me")
}

func (cr *CodeDBRedis) Update() error {
	panic("implement me")
}

func (cr *CodeDBRedis) GetCode(shortCode string) *domain.Code {
	panic("implement me")
}

func NewCodeDBRedis(client *redis.Client) *CodeDBRedis {
	return &CodeDBRedis{client: client}
}
