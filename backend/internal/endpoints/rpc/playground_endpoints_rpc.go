package rpc

import (
	"backend/api/playground"
	playgroundsvc "backend/internal/services/playground"
	"context"
)

type PlaygroundEndpoint struct {
	pgService *playgroundsvc.Service
}

func (p *PlaygroundEndpoint) Ping(ctx context.Context, _ *playground.EmptyRequest) (*playground.PingResponse, error) {
	out, err := p.pgService.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &playground.PingResponse{
		Message: out.Message,
	}, nil
}

func (p *PlaygroundEndpoint) RunCode(ctx context.Context, req *playground.CodeRequest) (*playground.RunResponse, error) {
	out, err := p.pgService.RunCode(ctx, &playgroundsvc.CommonCodeRequest{Body: req.Body})
	if err != nil {
		return nil, err
	}

	return &playground.RunResponse{
		Status:  out.Status,
		Output:  out.Output,
		Error:   out.Error,
		RunTime: out.RunTime,
	}, nil
}

func (p *PlaygroundEndpoint) FormatCode(ctx context.Context, req *playground.CodeRequest) (*playground.FormatCodeResponse, error) {
	out, err := p.pgService.FormatCode(ctx, &playgroundsvc.CommonCodeRequest{Body: req.Body})
	if err != nil {
		return nil, err
	}

	return &playground.FormatCodeResponse{
		FormattedCode: out.FormattedCode,
	}, nil
}

func (p *PlaygroundEndpoint) ShareCode(ctx context.Context, req *playground.CodeRequest) (*playground.ShareCodeResponse, error) {
	out, err := p.pgService.ShareCode(ctx, &playgroundsvc.CommonCodeRequest{Body: req.Body})
	if err != nil {
		return nil, err
	}

	return &playground.ShareCodeResponse{
		Code: out.Code,
	}, nil
}

func (p *PlaygroundEndpoint) GetCodeByShare(context.Context, *playground.CommonRequest) (*playground.GetCodeByShareResponse, error) {
	panic("implement me")
}

func NewPlaygroundEndpoints(pgService *playgroundsvc.Service) *PlaygroundEndpoint {
	return &PlaygroundEndpoint{pgService: pgService}
}
