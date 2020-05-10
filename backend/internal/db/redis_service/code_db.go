package redis_service

import (
	"backend/internal/domains"
	"encoding/json"
	"github.com/go-redis/redis"
)

type codeRedis struct {
	rdCl *redis.Client
}

func (cr *codeRedis) SaveCode(shortCode string, c domains.Code) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = cr.rdCl.Set(shortCode, b, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cr *codeRedis) GetCode(shortCode string) *domains.Code {
	panic("implement me")
}

func New(rdCl *redis.Client) *codeRedis {
	return &codeRedis{rdCl: rdCl}
}
