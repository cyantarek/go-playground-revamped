package outgoing

import (
	"backend/internal/core/domain"
	"context"
)

type CodeRepository interface {
	Save(ctx context.Context, c domain.Code) error
	Get(ctx context.Context, code domain.ShortCode) (domain.Code, error)
	NextID() string
}
