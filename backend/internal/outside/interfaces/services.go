package interfaces

import (
	"backend/internal/core/application/dtos"
	"context"
)

type CodeService interface {
	Format(ctx context.Context, data dtos.FormatCode) (dtos.FormatCodeResult, error)
	Run(ctx context.Context, data dtos.RunCode) (dtos.RunCodeResult, error)
	Shorten() error
	Expand() error
}
