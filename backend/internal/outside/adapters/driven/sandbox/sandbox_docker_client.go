package sandbox

import (
	"backend/internal/core/domain"
	"context"
)

type DockerBasedSandbox struct {

}

func (d DockerBasedSandbox) FormatCode(ctx context.Context, code domain.Code) (domain.Code, error) {
	panic("implement me")
}

func (d DockerBasedSandbox) CompileAndRun() error {
	panic("implement me")
}
