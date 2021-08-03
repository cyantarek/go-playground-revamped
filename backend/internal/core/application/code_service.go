package application

import (
	"backend/internal/core/application/dto"
	"backend/internal/core/domain"
	"backend/internal/core/port/outgoing"
	"context"
	"time"
)

type CodeService struct {
	codeRepo    outgoing.CodeRepository
	sandboxRepo outgoing.SandboxRepository
}

func NewCodeService(codeRepo outgoing.CodeRepository, sandboxRepo outgoing.SandboxRepository) CodeService {
	return CodeService{codeRepo: codeRepo, sandboxRepo: sandboxRepo}
}

func (c CodeService) Format(ctx context.Context, data dto.FormatCode) (dto.FormatCodeResult, error) {
	id := domain.NewCodeID(c.codeRepo.NextID())
	code := domain.NewCode(id, data.Code)

	// timeout for long running toxic malicious codes
	ctxCancel, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	formattedCode, err := c.sandboxRepo.FormatCode(ctxCancel, code)
	if err != nil {
		return dto.FormatCodeResult{}, err
	}

	return dto.FormatCodeResult{
		Code: formattedCode.Code(),
	}, nil
}

func (c CodeService) Run(ctx context.Context, data dto.RunCode) (dto.RunCodeResult, error) {
	id := domain.NewCodeID(c.codeRepo.NextID())
	code := domain.NewCode(id, data.Code)

	ctxCancel, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	runResult, err := c.sandboxRepo.CompileAndRun(ctxCancel, code)
	if err != nil {
		return dto.RunCodeResult{}, err
	}

	return dto.RunCodeResult{
		RunID:   runResult.RunID,
		Output:  runResult.Output,
		RunTime: runResult.RunTime,
	}, nil
}

func (c CodeService) Share(ctx context.Context, data dto.ShareCode) (dto.ShareCodeResult, error) {
	id := domain.NewCodeID(c.codeRepo.NextID())
	code := domain.NewCode(id, data.Code)

	err := c.codeRepo.Save(ctx, code)
	if err != nil {
		return dto.ShareCodeResult{}, err
	}

	return dto.ShareCodeResult{
		ShortCode: id.String(),
	}, nil
}

func (c CodeService) GetByID(ctx context.Context, id string) (dto.GetCodeByID, error) {
	code, err := c.codeRepo.Get(ctx, domain.NewShortCode(id))
	if err != nil {
		return dto.GetCodeByID{}, err
	}

	code.MarkVisit()
	err = c.codeRepo.Save(ctx, code)
	if err != nil {
		return dto.GetCodeByID{}, err
	}

	return dto.GetCodeByID{
		ID:        code.ID().String(),
		Code:      code.Code(),
		CreatedAt: code.WhenCreated(),
		UpdatedAt: code.WhenLastUpdated(),
	}, nil
}
