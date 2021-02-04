.PHONY: start
start: ## Launch registry
	@docker run --rm -d --network=host --name=k8s-vault-secrets_minikube_1 alpine:3.13.1 ash -c "apk add socat && socat TCP-LISTEN:5000,reuseaddr,fork TCP:$(shell minikube ip):5000"

.PHONY: stop
stop: ## Stop registry
	@docker stop k8s-vault-secrets_minikube_1
