package outgoing

import (
	"backend/internal/core/application/dto"
	"backend/internal/core/domain"
	"context"
)

type SandboxRepository interface {
	FormatCode(ctx context.Context, code domain.Code) (domain.Code, error)
	CompileAndRun(ctx context.Context, code domain.Code) (dto.RunCodeResult, error)
}
