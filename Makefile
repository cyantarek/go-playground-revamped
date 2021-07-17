.PHONY: proto

APP_NAME?=playground
RANDOM?=$(shell bash -c 'echo $$RANDOM')

generate-mock:
	mockgen -source=internal/services/authservice/service.go -destination=internal/mocks/auth_service_mock.go -package=mocks

test:
	go test -v ./...

run:
	docker-compose up -d

SERVICES := playground
# SERVICES := auth payment content upload contact

proto:
	for service in $(SERVICES); do \
  		cd proto && \
  		protoc --go_out=.. \
		--go-grpc_out=.. \
		--grpc-gateway_out=logtostderr=true:.. \
		--js_out=import_style=commonjs:../frontend/src/api/playground \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:../frontend/src/api/playground \
  		"$$service"_*.proto && cd ..; \
    done

docker-run:
	docker-compose up -d

docker-down:
	docker-compose down
