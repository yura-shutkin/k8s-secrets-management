NAMESPACE=bank-vaults

.PHONY: namespace-create
namespace-create: ## Create namespace basic-setup
	@kubectl --namespace $(NAMESPACE) apply -f namespace.yml

.PHONY: webhook-install
namespace=bank-vaults
webhook-install: ## Install webhook
	@helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com
	@helm repo update
	@helm upgrade --namespace $(namespace) --install -f values.yml default banzaicloud-stable/vault-secrets-webhook --version=1.13.0

.PHONY: webhook-uninstall
namespace=bank-vaults
webhook-uninstall: ## Uninstall webhook
	@helm uninstall --namespace $(namespace) default

.PHONY: webhook-sa-secret-name-get
webhook-sa-secret-name-get: ## Show default-vault-secrets-webhook sa secret name
	@kubectl --namespace $(NAMESPACE) get serviceaccounts/default-vault-secrets-webhook -o json | jq -r '.secrets[0].name'

.PHONY: webhook-sa-token-get
webhook-sa-token-get: ## Get token of default SA
	kubectl --namespace $(NAMESPACE) get secret $(shell make get-sa-secret-name) -o json | jq -r '.data.token' | base64 --decode; echo
