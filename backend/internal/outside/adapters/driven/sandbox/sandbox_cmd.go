package sandbox

import (
	"backend/internal/core/application/dto"
	"backend/internal/core/domain"
	"backend/pkg/generator"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

type CMDBasedSandbox struct{}

func NewCMDBasedSandbox() CMDBasedSandbox {
	return CMDBasedSandbox{}
}

const (
	workingDir = "tmp"
)

func (c CMDBasedSandbox) FormatCode(ctx context.Context, code domain.Code) (domain.Code, error) {
	var resultChan = make(chan string)
	var errChan = make(chan error)

	runID := generator.RandomStringGenerator(5)

	go runSandboxContainer(resultChan, errChan, code, "fmt", runID)

	select {
	case <-ctx.Done():
		return domain.Code{}, ctx.Err()
	case result := <-resultChan:
		return domain.NewCode(code.CodeID(), result), nil
	case err := <-errChan:
		return domain.Code{}, err
	}
}

func (c CMDBasedSandbox) CompileAndRun(ctx context.Context, code domain.Code) (dto.RunCodeResult, error) {
	var resultChan = make(chan string)
	var errChan = make(chan error)

	var starTime = time.Now()

	runID := generator.RandomStringGenerator(5)

	go runSandboxContainer(resultChan, errChan, code, "run", runID)

	select {
	case <-ctx.Done():
		return dto.RunCodeResult{}, ctx.Err()
	case result := <-resultChan:
		executionTime := time.Since(starTime).Seconds()

		return dto.RunCodeResult{
			RunID:   runID,
			Output:  result,
			RunTime: executionTime,
		}, nil
	case err := <-errChan:
		return dto.RunCodeResult{}, err
	}
}

// Reference Command:
// docker run --rm -v $(pwd)/files/:/files golang bash -c 'go run /files/print.go'
// docker run --rm -it golang timeout 3 sh -c '/bin/sleep 10s'

func runSandboxContainer(resultChan chan string, errChan chan error, code domain.Code, command, runID string) {
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s.go", workingDir, runID), []byte(code.Code()), 0644)
	if err != nil {
		errChan <- err
		return
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker run --rm -v $(pwd)/%s/:/files golang timeout %d bash -c 'go %s /files/%s.go'", workingDir, 2, command, runID))

	var stdErr, stdOut bytes.Buffer

	cmd.Stderr = &stdErr
	cmd.Stdout = &stdOut

	defer func() {
		err = os.Remove(fmt.Sprintf("%s/%s.go", workingDir, runID))
		if err != nil {
			errChan <- err
			return
		}
	}()

	err = cmd.Run()
	if err != nil {
		errChan <- err
		return
	}

	if command == "run" {
		resultChan <- stdOut.String()
	}

	if command == "fmt" {
		b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.go", workingDir, runID))
		if err != nil {
			errChan <- err
			return
		}

		if len(b) == 0 {
			errChan <- errors.New("timeout")
			return
		}

		resultChan <- string(b)
	}
}
