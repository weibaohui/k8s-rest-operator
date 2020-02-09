all: build push deploy

build:
	docker build -t quay.io/weibaohui/k8s-rest-operator .
push:
	docker push quay.io/weibaohui/k8s-rest-operator
deploy:
	kubectl apply -f deployment.yaml