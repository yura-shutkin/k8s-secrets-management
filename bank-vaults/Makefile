.PHONY: create-namespace
create-namespace: ## Create namespace basic-setup
	@kubectl --namespace bank-vaults apply -f namespace.yml

.PHONY: deploy-webhook
namespace=bank-vaults
deploy-webhook: ## Setup webhook
	@helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com
	@helm repo update
	@helm upgrade --namespace $(namespace) --install default banzaicloud-stable/vault-secrets-webhook

.PHONY: setup-rbac
setup-rbac: ## Setup rbac
	@kubectl --namespace default apply -f rbac/
