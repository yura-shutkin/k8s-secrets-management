.PHONY: create-namespace
create-namespace: ## Create namespace basic-setup
	@kubectl --namespace basic-setup apply -f namespace.yml
	@kubectl --namespace basic-setup create serviceaccount vault-auth

.PHONY: create-secret
create-secret: ## Create secret in k8s
	@kubectl --namespace basic-setup apply -f secrets/db-creds.yml

.PHONY: apply-web-app
apply-web-app: ## Deploy basic service
	@kubectl --namespace basic-setup apply -f web-app
	@minikube --namespace basic-setup service --url web-app

.PHONY: apply-vault
apply-vault: ## Deploy vault
	@kubectl --namespace basic-setup apply -f vault
	@minikube --namespace basic-setup service --url vault

.PHONY: apply-bank-vaults
apply-bank-vaults: ## Deploy bank-vaults (Vault secrets webhook)
	@cd banzaicloud/bank-vaults/ && make deploy-webhook
