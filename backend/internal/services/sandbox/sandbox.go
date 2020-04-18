package sandbox

import (
	"backend/pkg/generator"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Reference Command:
// docker run --rm -v $(pwd)/files/:/files golang bash -c 'go run /files/print.go'

const (
	workingDir = "tmp"
)

func CompileAndRun(body []byte) (string, error, float64) {
	randUid := generator.RandomStringGenerator(5)
	
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s.go", workingDir, randUid), body, 0644)
	if err != nil {
		log.Fatal(err)
	}
	
	start := time.Now()
	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker run --rm --network none -v $(pwd)/%s/:/files golang bash -c 'go run /files/%s.go'", workingDir, randUid))
	var stdOut, stdErr bytes.Buffer
	
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	
	defer func() {
		err = os.Remove(fmt.Sprintf("%s/%s.go", workingDir, randUid))
		if err != nil {
			log.Println(err)
		}
	}()
	
	err = cmd.Run()
	if err != nil {
		return "", errors.New(stdErr.String()), math.Round(time.Since(start).Seconds())
	}
	
	return stdOut.String(), nil, math.Round(time.Since(start).Seconds())
}

func FormatCode(body []byte) (string, error) {
	randUid := RandStringRunes(5)
	
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s.go", workingDir, randUid), body, 0644)
	if err != nil {
		log.Fatal(err)
	}
	
	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker run --rm -v $(pwd)/%s/:/files golang bash -c 'go fmt /files/%s.go'", workingDir, randUid))
	var stdErr bytes.Buffer
	
	cmd.Stderr = &stdErr
	
	defer func() {
		err = os.Remove(fmt.Sprintf("%s/%s.go", workingDir, randUid))
		if err != nil {
			log.Println(err)
		}
	}()
	
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return "", errors.New(stdErr.String())
	}
	
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.go", workingDir, randUid))
	if err != nil {
		log.Fatal(err)
	}
	
	return string(b), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
