.PHONY: create-namespace
create-namespace: ## Create namespace basic-setup
	@kubectl --namespace basic-setup apply -f namespace.yml

.PHONY: create-secret
create-secret: ## Create secret in k8s
	@kubectl --namespace basic-setup apply -f secrets/db-creds.yml

.PHONY: deploy-web-app
deploy-web-app: ## Deploy basic web-app
	@kubectl --namespace basic-setup apply -f web-app
	@minikube --namespace basic-setup service --url web-app

.PHONY: rollout-web-app
rollout-web-app: ## Rollout restart web-app
	@kubectl --namespace basic-setup rollout restart deployments/web-app

.PHONY: delete-web-app
delete-web-app: ## Delete web-app
	@kubectl --namespace basic-setup delete -f web-app

.PHONY: deploy-second-app
deploy-second-app: ## Deploy basic second-app
	@kubectl --namespace basic-setup apply -f second-app
	@minikube --namespace basic-setup service --url second-app

.PHONY: rollout-second-app
rollout-second-app: ## Rollout restart second-app
	@kubectl --namespace basic-setup rollout restart deployments/second-app

.PHONY: delete-second-app
delete-second-app: ## Delete second-app
	@kubectl --namespace basic-setup delete -f second-app

.PHONY: start-ubuntu-pod
start-ubuntu-pod: ## Start pod with ubuntu
	@kubectl --namespace basic-setup run ubuntu --image=ubuntu:20.04 sleep infinity

.PHONY: stop-ubuntu-pod
stop-ubuntu-pod: ## Stop and delete pod with ubuntu
	@kubectl --namespace basic-setup delete ubuntu

.PHONY: attach-ubuntu-pod
attach-ubuntu-pod: ## Exec command in ubuntu pod
	@kubectl --namespace basic-setup exec -ti ubuntu -- bash

.PHONY: update-vault-auth
update-vault-auth: ## Setup CA and reviewer JWT
	curl -sX POST -H "X-Vault-Token: 12345" -d "{\"token_reviewer_jwt\": \"$(shell kubectl --namespace basic-setup get secret default-token-tjghg -o json | jq -r '.data.token' | base64 -d)\", \"kubernetes_host\": \"https://10.96.0.1\"}" $(shell minikube --namespace vault service --url vault)/v1/auth/kubernetes/config
