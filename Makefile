all: build push deploy

build:
	docker build -t quay.io/cmit/k8s-rest-operator .
push:
	docker push quay.io/cmit/k8s-rest-operator
deploy:
	kubectl apply -f deploy/deployment.yaml