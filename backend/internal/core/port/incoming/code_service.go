package incoming

import (
	"backend/internal/core/application/dto"
	"context"
)

type CodeService interface {
	Format(ctx context.Context, data dto.FormatCode) (dto.FormatCodeResult, error)
	Run(ctx context.Context, data dto.RunCode) (dto.RunCodeResult, error)
	Share(ctx context.Context, data dto.ShareCode) (dto.ShareCodeResult, error)
	Expand(ctx context.Context, id string) (dto.GetCodeByID, error)
}
