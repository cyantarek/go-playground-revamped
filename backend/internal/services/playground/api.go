package playground

import (
	"backend/internal/domains"
	"backend/internal/services/sandbox"
	"backend/pkg/generator"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

func (p *Service) Ping(_ context.Context) (*PingResponse, error) {
	return &PingResponse{
		Message: "Ok",
	}, nil
}

func (p *Service) GetCodeByShare(_ context.Context, req *CodeByShareRequest) (*CodeByShareResponse, error) {
	panic("implement me")
}

func (p *Service) ShareCode(_ context.Context, req *CommonCodeRequest) (*ShaCodeResponse, error) {
	if len(req.Body) == 0 {
		return nil, status.Error(1000, "empty body is not allowed")
	}
	
	randUid := generator.RandomStringGenerator(5)
	
	newCode := domains.Code{
		Body:      req.Body,
		ShortCode: randUid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	err := p.codeService.SaveCode(randUid, newCode)
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(1001, "internal server error")
	}
	
	return &ShaCodeResponse{
		Code: randUid,
	}, nil
}

func (p *Service) FormatCode(_ context.Context, req *CommonCodeRequest) (*FormatCodeResponse, error) {
	if len(req.Body) == 0 {
		return nil, status.Error(1000, "empty body is not allowed")
	}
	
	output, err := sandbox.FormatCode([]byte(req.Body))
	if err != nil {
		return nil, status.Error(1002, err.Error())
	}
	
	return &FormatCodeResponse{
		FormattedCode: output,
	}, nil
}

func (p *Service) RunCode(_ context.Context, req *CommonCodeRequest) (*CodeRunResponse, error) {
	if len(req.Body) == 0 {
		return nil, status.Error(1000, "empty body is not allowed")
	}
	
	if strings.Contains(req.Body, "net/http") || strings.Contains(req.Body, "os") {
		return nil, status.Error(1000, "malicious code is not allowed. If you wonder what does it mean, check about")
	}
	
	output, err, runTime := sandbox.CompileAndRun([]byte(req.Body))
	
	response := CodeRunResponse{
		Output:  output,
		RunTime: runTime,
	}
	
	if err != nil {
		response.Status = "failed"
		response.Error = err.Error()
	} else {
		response.Status = "ok"
	}
	
	return &response, nil
}
