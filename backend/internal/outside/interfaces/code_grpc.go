package interfaces

import (
	"backend/internal/core/application/dtos"
	playgroundproto "backend/pkg/proto/playground"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CodeGRPC struct {
	codeService CodeService
	playgroundproto.UnimplementedPlaygroundServer
}

func NewCodeGRPC(codeService CodeService) *CodeGRPC {
	return &CodeGRPC{codeService: codeService}
}

func (c CodeGRPC) Ping(ctx context.Context, request *playgroundproto.EmptyRequest) (*playgroundproto.PingResponse, error) {
	return &playgroundproto.PingResponse{
		Message: "Alive :)",
	}, nil
}

func (c CodeGRPC) FormatCode(ctx context.Context, request *playgroundproto.FormatCodeRequest) (*playgroundproto.FormatCodeResponse, error) {
	formattedCode, err := c.codeService.Format(ctx, dtos.FormatCode{
		Body:     request.GetCode(),
		Language: request.GetLanguage(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &playgroundproto.FormatCodeResponse{FormattedCode: formattedCode.Body}, nil
}

func (c CodeGRPC) RunCode(ctx context.Context, request *playgroundproto.RunCodeRequest) (*playgroundproto.RunResponse, error) {
	panic("implement me")
}

func (c CodeGRPC) ShareCode(ctx context.Context, request *playgroundproto.ShareCodeRequest) (*playgroundproto.ShareCodeResponse, error) {
	panic("implement me")
}

func (c CodeGRPC) GetCodeByShare(ctx context.Context, request *playgroundproto.CommonRequest) (*playgroundproto.GetCodeByShareResponse, error) {
	panic("implement me")
}

func (c CodeGRPC) Wire(grpcServer *grpc.Server, httpServer *runtime.ServeMux) {
	playgroundproto.RegisterPlaygroundServer(grpcServer, c)
	_ = playgroundproto.RegisterPlaygroundHandlerServer(context.Background(), httpServer, c)
}
