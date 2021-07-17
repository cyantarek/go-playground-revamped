package sandbox

import (
	"backend/internal/core/application/dtos"
	"backend/internal/core/domain"
	"backend/pkg/generator"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

type CMDBasedSandbox struct {}

func NewCMDBasedSandbox() CMDBasedSandbox {
	return CMDBasedSandbox{}
}

const (
	workingDir = "tmp"
)

func (c CMDBasedSandbox) FormatCode(ctx context.Context, code domain.Code) (domain.Code, error) {
	var resultChan = make(chan string)
	var errChan = make(chan error)

	go runSandboxContainer(resultChan, errChan, code, "fmt")

	select {
	case <-ctx.Done():
		return domain.Code{}, ctx.Err()
	case result := <- resultChan:
		return domain.NewCode(result), nil
	case err := <-errChan:
		return domain.Code{}, err
	}
}

func (c CMDBasedSandbox) CompileAndRun(ctx context.Context, code domain.Code) (dtos.RunCodeResult, error) {
	//result, err := runSandboxContainer(ctx, code, "run")
	//if err != nil {
	//	return dtos.RunCodeResult{}, err
	//}
	//
	//return dtos.RunCodeResult{
	//	RunID:   "",
	//	Status:  "ok",
	//	Output:  result,
	//	RunTime: 0,
	//}, nil

	return dtos.RunCodeResult{}, nil
}

// Reference Command:
// docker run --rm -v $(pwd)/files/:/files golang bash -c 'go run /files/print.go'

func runSandboxContainer(resultChan chan string, errChan chan error, code domain.Code, command string) {
	time.Sleep(time.Second*10)

	randUid := generator.RandomStringGenerator(5)

	err := ioutil.WriteFile(fmt.Sprintf("%s/%s.go", workingDir, randUid), []byte(code.Body()), 0644)
	if err != nil {
		errChan <- err
		return
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker run --rm -v $(pwd)/%s/:/files golang bash -c 'go %s /files/%s.go'", workingDir, command, randUid))

	var stdErr bytes.Buffer

	cmd.Stderr = &stdErr

	defer func() {
		err = os.Remove(fmt.Sprintf("%s/%s.go", workingDir, randUid))
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

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.go", workingDir, randUid))
	if err != nil {
		errChan <- err
		return
	}

	resultChan <- string(b)
}
