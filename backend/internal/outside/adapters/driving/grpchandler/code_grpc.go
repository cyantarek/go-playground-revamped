package grpchandler

import (
	"backend/internal/core/application/dto"
	"backend/internal/core/port/incoming"
	playgroundproto "backend/pkg/proto/playground"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CodeGRPC struct {
	codeService incoming.CodeService
	playgroundproto.UnimplementedPlaygroundServer
}

func NewCodeGRPC(codeService incoming.CodeService) *CodeGRPC {
	return &CodeGRPC{codeService: codeService}
}

func (c CodeGRPC) Ping(ctx context.Context, request *playgroundproto.EmptyRequest) (*playgroundproto.PingResponse, error) {
	return &playgroundproto.PingResponse{
		Message: "Alive :)",
	}, nil
}

func (c CodeGRPC) FormatCode(ctx context.Context, request *playgroundproto.FormatCodeRequest) (*playgroundproto.FormatCodeResponse, error) {
	formattedCode, err := c.codeService.Format(ctx, dto.FormatCode{
		Code:     request.GetCode(),
		Language: request.GetLanguage(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &playgroundproto.FormatCodeResponse{FormattedCode: formattedCode.Code}, nil
}

func (c CodeGRPC) RunCode(ctx context.Context, request *playgroundproto.RunCodeRequest) (*playgroundproto.RunCodeResponse, error) {
	result, err := c.codeService.Run(ctx, dto.RunCode{
		Code:     request.GetCode(),
		Language: request.GetLanguage(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &playgroundproto.RunCodeResponse{
		RunId:   result.RunID,
		Output:  result.Output,
		RunTime: result.RunTime,
	}, nil
}

func (c CodeGRPC) ShareCode(ctx context.Context, request *playgroundproto.ShareCodeRequest) (*playgroundproto.ShareCodeResponse, error) {
	shareResult, err := c.codeService.Share(ctx, dto.ShareCode{
		Code:     request.GetCode(),
		Language: request.GetLanguage(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &playgroundproto.ShareCodeResponse{ShortCode: shareResult.ShortCode}, nil
}

func (c CodeGRPC) GetCodeByShare(ctx context.Context, request *playgroundproto.CodeByIDRequest) (*playgroundproto.GetCodeByShareResponse, error) {
	result, err := c.codeService.Expand(ctx, request.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &playgroundproto.GetCodeByShareResponse{
		Id:        result.ID,
		Code:      result.Code,
		CreatedAt: result.CreatedAt.String(),
		UpdatedAt: result.UpdatedAt.String(),
	}, nil
}

func (c CodeGRPC) Wire(grpcServer *grpc.Server, httpServer *runtime.ServeMux) {
	playgroundproto.RegisterPlaygroundServer(grpcServer, c)
	_ = playgroundproto.RegisterPlaygroundHandlerServer(context.Background(), httpServer, c)
}
