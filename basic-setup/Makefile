.PHONY: create-namespace
create-namespace: ## Create namespace basic-setup
	@kubectl --namespace basic-setup apply -f namespace.yml
	@kubectl --namespace basic-setup create serviceaccount vault-auth

.PHONY: create-secret
create-secret: ## Create secret in k8s
	@kubectl --namespace basic-setup apply -f secrets/db-creds.yml

.PHONY: deploy-web-app
deploy-web-app: ## Deploy basic service
	@kubectl --namespace basic-setup apply -f web-app
	@minikube --namespace basic-setup service --url web-app

.PHONY: rollout-web-app
rollout-web-app: ## Rollout restart web-app
	@kubectl --namespace basic-setup rollout restart deployments/web-app

.PHONY: deploy-bank-vaults
deploy-bank-vaults: ## Deploy bank-vaults (Vault secrets webhook)
	@cd banzaicloud/bank-vaults/ && make deploy-webhook

.PHONY: setup-roles
setup-roles: ## Setup roles
	@kubectl --namespace basic-setup apply -f vault-auth
