package sandbox

import (
	"backend/internal/core/domain"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type DockerBasedSandbox struct {
	client *client.Client
}

func NewDockerBasedSandbox(client *client.Client) *DockerBasedSandbox {
	return &DockerBasedSandbox{client: client}
}

func (d DockerBasedSandbox) FormatCode(ctx context.Context, code domain.Code) (domain.Code, error) {
	containerCreated, err := d.client.ContainerCreate(
		ctx,
		&container.Config{},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		&v1.Platform{},
		"",
	)
	if err != nil {
		return domain.Code{}, err
	}

	_ = d.client.ContainerStart(ctx, containerCreated.ID, types.ContainerStartOptions{})

	return domain.Code{}, nil
}

func (d DockerBasedSandbox) CompileAndRun() error {
	panic("implement me")
}
