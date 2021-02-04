.PHONY: deploy-webhook
deploy-webhook: ## Setup webhook
	@helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com
	@helm repo update
	@helm upgrade --namespace basic-setup --install default banzaicloud-stable/vault-secrets-webhook
