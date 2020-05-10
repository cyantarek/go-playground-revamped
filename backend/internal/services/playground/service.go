package playground

import (
	"backend/internal/domains"
)

// Server represents the actual grpc server implementation
type Service struct {
	codeService domains.CodeService
}

// New constructs the server and returns it
func New(codeService domains.CodeService) (*Service, error) {
	srv := Service{
		codeService: codeService,
	}

	return &srv, nil
}
