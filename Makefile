.PHONY: proto kubernetes-redeploy

RANDOM?=$(shell bash -c 'echo $$RANDOM')

proto:
	protoc -I ./proto -I ./proto/third_party --go_out=plugins=grpc:backend/api/playground --grpc-gateway_out=logtostderr=true:backend/api/playground --js_out=import_style=commonjs:frontend/src/api/playground --grpc-web_out=import_style=commonjs,mode=grpcwebtext:frontend/src/api/playground proto/playground_*.proto

dockerize-all: dockerize-frontend dockerize-backend

dockerize-frontend:
	docker build -t tarek5/todo-grpc-frontend:latest --no-cache -f frontend/Dockerfile .
	docker image push tarek5/todo-grpc-frontend:latest

dockerize-envoy:
	docker build -t tarek5/todo-grpc-envoy:latest --no-cache -f envoy-proxy/Dockerfile .
	docker image push tarek5/todo-grpc-envoy:latest

dockerize-backend:
	docker build -t tarek5/todo-grpc-backend:latest --no-cache -f backend/build/docker/Dockerfile .
	docker image push tarek5/todo-grpc-backend:latest

kubernetes-deploy:
	kubectl apply -f frontend/frontend-deployment.yml
	kubectl apply -f backend/build/kubernetes/backend-deployment.yml

kubernetes-deploy-ingress-controller:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.30.0/deploy/static/mandatory.yaml
	# kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx

kubernetes-deploy-ingress-resource:
	kubectl apply -f ingress-grpc-web.yml
	# kubectl get ingress todo-ingress

kubernetes-port-foreward:
	kubectl port-forward svc/todo-backend-service 5200:5200

kubernetes-destroy:
	kubectl delete -f frontend/frontend-deployment.yml
	kubectl delete -f backend/build/kubernetes/backend-deployment.yml

istio-setup:
	minikube addons list
	minikube addons enable metrics-server
	istioctl manifest apply
	kubectl get svc -n istio-system

istion-uninstall:
	kubectl delete ns istio-system

kubernetes-redeploy:
	kubectl patch deployment todo-backend-deployment -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"build\":\"$(RANDOM)\"}}}}}"
	kubectl patch deployment todo-frontend-deployment -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"build\":\"$(RANDOM)\"}}}}}"
