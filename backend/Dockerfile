FROM golang:latest
ARG BIN_NAME=playground
COPY .. /$BIN_NAME
WORKDIR /$BIN_NAME
RUN CGO_ENABLED=0 GOOS=linux GIT_COMMIT=$(git rev-list -1 HEAD) go build -ldflags "-X main.GitCommit=$GIT_COMMIT" -o $BIN_NAME ./cmd/$BIN_NAME
ENTRYPOINT [ "./playground" ]

EXPOSE 5000 5001
