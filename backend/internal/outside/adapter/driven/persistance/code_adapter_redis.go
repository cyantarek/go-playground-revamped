package persistance

import (
	"backend/internal/core/domain"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"time"
)

type CodeDBRedis struct {
	client *redis.Client
}

func (cr *CodeDBRedis) NextID() string {
	id, _ := uuid.NewUUID()

	return id.String()
}

func (cr *CodeDBRedis) Save(ctx context.Context, c domain.Code) error {
	code := codeRedis{
		ID:        c.ID().String(),
		Code:      c.Code(),
		CreatedAt: c.WhenCreated(),
		UpdatedAt: c.WhenLastUpdated(),
	}

	b, err := json.Marshal(code)
	if err != nil {
		return err
	}

	err = cr.client.Set(c.ID().String(), b, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

type codeRedis struct {
	ID        string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cr *CodeDBRedis) Get(ctx context.Context, code domain.ShortCode) (domain.Code, error) {
	result, err := cr.client.Get(code.String()).Bytes()
	if err != nil {
		return domain.Code{}, err
	}

	var out codeRedis

	err = json.Unmarshal(result, &out)
	if err != nil {
		return domain.Code{}, err
	}

	domainCode := domain.NewCode(domain.NewCodeID(out.ID), out.Code)
	domainCode.TimeWise(out.CreatedAt, out.UpdatedAt)

	return domainCode, nil
}

func (cr *CodeDBRedis) GetCode(shortCode string) *domain.Code {
	panic("implement me")
}

func NewCodeDBRedis(client *redis.Client) *CodeDBRedis {
	return &CodeDBRedis{client: client}
}
