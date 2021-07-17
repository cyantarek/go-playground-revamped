package application

import (
	"backend/internal/core/application/dtos"
	"backend/internal/core/domain"
	"context"
	"time"
)

type CodeService struct {
	codeRepo CodeRepository
	sandboxRepo SandboxRepository
}

func NewCodeService(codeRepo CodeRepository, sandboxRepo SandboxRepository) CodeService {
	return CodeService{codeRepo: codeRepo, sandboxRepo: sandboxRepo}
}

func (c CodeService) Format(ctx context.Context, data dtos.FormatCode) (dtos.FormatCodeResult, error) {
	code := domain.NewCode(data.Body)

	ctxCancel, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	formattedCode, err := c.sandboxRepo.FormatCode(ctxCancel, code)
	if err != nil {
		return dtos.FormatCodeResult{}, err
	}

	return dtos.FormatCodeResult{
		Body: formattedCode.Body(),
	}, nil
}

func (c CodeService) Run(ctx context.Context, data dtos.RunCode) (dtos.RunCodeResult, error) {
	code := domain.NewCode(data.Body)
	return c.sandboxRepo.CompileAndRun(ctx, code)
}

func (c CodeService) Shorten() error {
	panic("implement me")
}

func (c CodeService) Expand() error {
	panic("implement me")
}

