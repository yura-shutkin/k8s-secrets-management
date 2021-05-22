NAMESPACE="prod-setup"

.PHONY: namespace-create
namespace-create: ## Create namespace prod-setup
	@kubectl --namespace $(NAMESPACE) apply -f namespace.yml

.PHONY: project-1-prod-app-deploy
project-1-prod-app-deploy: ## Deploy basic project-1/prod-app
	@kubectl --namespace $(NAMESPACE) apply -f project-1__prod-app

.PHONY: project-1-prod-app-rollout
project-1-prod-app-rollout: ## Rollout restart project-1/prod-app
	@kubectl --namespace $(NAMESPACE) rollout restart deployments/project-1-prod-app

.PHONY: project-1-prod-app-delete
project-1-prod-app-delete: ## Delete project-1/prod-app
	@kubectl --namespace $(NAMESPACE) delete -f project-1__prod-app

.PHONY: default-sa-secret-name-get
default-sa-secret-name-get: ## Show default sa secret name
	@kubectl --namespace $(NAMESPACE) get serviceaccounts/default -o json | jq -r '.secrets[0].name'

.PHONY: default-sa-token-get
default-sa-token-get: ## Get token of default SA
	@kubectl --namespace $(NAMESPACE) get secret $(shell make default-sa-secret-name-get) -o json | jq -r '.data.token' | base64 --decode

.PHONY: project-1-prod-app-pod-get
project-1-prod-app-pod-get: ## Get id of vault pod
	@kubectl --namespace $(NAMESPACE) get pods --selector app=project-1-prod-app -o json | jq -r '.items[0].metadata.name'

.PHONY: project-1-prod-app-port-forward
WEB_APP_POD=$(shell make project-1-prod-app-pod-get)
project-1-prod-app-port-forward: ## Forward port from vault to localhost
	@kubectl --namespace $(NAMESPACE) port-forward pods/$(WEB_APP_POD) 12080:8080

.PHONY: project-1-prod-app-check-login
project-1-prod-app-check-login: ## Check if we can log into vault
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-1__prod-ro\",\"jwt\":\"$(shell make default-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'
