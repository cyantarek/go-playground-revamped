package application

import (
	"backend/internal/core/application/dtos"
	"backend/internal/core/domain"
	"context"
)

type SandboxRepository interface {
	FormatCode(ctx context.Context, code domain.Code) (domain.Code, error)
	CompileAndRun(ctx context.Context, code domain.Code) (dtos.RunCodeResult, error)
}

type CodeRepository interface {
	Save(ctx context.Context, c domain.Code) error
	Get() error
	Update() error
}
