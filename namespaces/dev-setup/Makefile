NAMESPACE=dev-setup

.PHONY: namespace-create
namespace-create: ## Create namespace basic-setup
	@kubectl --namespace $(NAMESPACE) apply -f namespace.yml
	@kubectl --namespace $(NAMESPACE) apply -f project-2-dev-sa.yml

########################################################################################################################
#### common ############################################################################################################
########################################################################################################################

.PHONY: default-sa-secret-name-get
default-sa-secret-name-get: ## Show default sa secret name
	@kubectl --namespace $(NAMESPACE) get serviceaccounts/default -o json | jq -r '.secrets[0].name'

.PHONY: default-sa-token-get
default-sa-token-get: ## Get token of default SA
	@kubectl --namespace $(NAMESPACE) get secret $(shell make default-sa-secret-name-get) -o json | jq -r '.data.token' | base64 --decode

.PHONY: secret-create
secret-create: ## Create secret in k8s
	@kubectl --namespace $(NAMESPACE) apply -f secrets/db-creds.yml

########################################################################################################################
#### project-1 #### dev ################################################################################################
########################################################################################################################

.PHONY: project-1-dev-app-deploy
project-1-dev-app-deploy: ## Deploy basic dev-app
	@kubectl --namespace $(NAMESPACE) apply -f project-1__dev-app

.PHONY: project-1-dev-app-rollout
project-1-dev-app-rollout: ## Rollout restart dev-app
	@kubectl --namespace $(NAMESPACE) rollout restart deployments/project-1-dev-app

.PHONY: project-1-dev-app-delete
project-1-dev-app-delete: ## Delete dev-app
	@kubectl --namespace $(NAMESPACE) delete -f project-1__dev-app

.PHONY: project-1-dev-app-pod-get
project-1-dev-app-pod-get: ## Get id of vault pod
	@kubectl --namespace $(NAMESPACE) get pods --selector app=project-1-dev-app -o json | jq -r '.items[0].metadata.name'

.PHONY: project-1-dev-app-port-forward
PROJECT_1_DEV_APP_POD=$(shell make project-1-dev-app-pod-get)
project-1-dev-app-port-forward: ## Forward port from vault to localhost
	@kubectl --namespace $(NAMESPACE) port-forward pods/$(PROJECT_1_DEV_APP_POD) 11080:8080

.PHONY: project-1-dev-app-check-login
project-1-dev-app-check-login: ## Check if we can log into vault
# TODO: check on windows and mac os
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-1__dev-ro\",\"jwt\":\"$(shell make default-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'

########################################################################################################################
#### project-1 #### stage ##############################################################################################
########################################################################################################################

.PHONY: project-1-stage-app-deploy
project-1-stage-app-deploy: ## Deploy basic stage-app
	@kubectl --namespace $(NAMESPACE) apply -f project-1__stage-app

.PHONY: project-1-stage-app-rollout
project-1-stage-app-rollout: ## Rollout restart stage-app
	@kubectl --namespace $(NAMESPACE) rollout restart deployments/project-1stage-app

.PHONY: project-1-stage-app-delete
project-1-stage-app-delete: ## Delete stage-app
	@kubectl --namespace $(NAMESPACE) delete -f project-1__stage-app

.PHONY: project-1-stage-app-pod-get
project-1-stage-app-pod-get: ## Get id of vault pod
	@kubectl --namespace $(NAMESPACE) get pods --selector app=project-1-stage-app -o json | jq -r '.items[0].metadata.name'

.PHONY: project-1-stage-app-port-forward
PROJECT_1_STAGE_APP_POD=$(shell make project-1-stage-app-pod-get)
project-1-stage-app-port-forward: ## Forward port from vault to localhost
	@kubectl --namespace $(NAMESPACE) port-forward pods/$(PROJECT_1_STAGE_APP_POD) 11081:8080

.PHONY: project-1-stage-app-check-login
project-1-stage-app-check-login: ## Check if we can log into vault
# TODO: check on windows and mac os
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-1__stage-ro\",\"jwt\":\"$(shell make default-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'

########################################################################################################################
#### project-2 #### dev ################################################################################################
########################################################################################################################

.PHONY: project-2-dev-app-deploy
project-2-dev-app-deploy: ## Deploy project-2/dev-app
	@kubectl --namespace $(NAMESPACE) apply -f project-2__dev-app

.PHONY: project-2-dev-app-rollout
project-2-dev-app-rollout: ## Rollout restart project-2/dev-app
	@kubectl --namespace $(NAMESPACE) rollout restart deployments/project-2-dev-app

.PHONY: project-2-dev-app-delete
project-2-dev-app-delete: ## Delete dev-app
	@kubectl --namespace $(NAMESPACE) delete -f project-2__dev-app

.PHONY: project-2-dev-sa-secret-name-get
project-2-dev-sa-secret-name-get: ## Show project-2-dev sa secret name
	@kubectl --namespace $(NAMESPACE) get serviceaccounts/project-2-dev -o json | jq -r '.secrets[0].name'

.PHONY: project-2-dev-sa-token-get
project-2-dev-sa-token-get: ## Get token of project-2-dev SA
	@kubectl --namespace $(NAMESPACE) get secret $(shell make project-2-dev-sa-secret-name-get) -o json | jq -r '.data.token' | base64 --decode

.PHONY: project-2-dev-app-pod-get
project-2-dev-app-pod-get: ## Get id of project-2/dev-app pod
	@kubectl --namespace $(NAMESPACE) get pods --selector app=project-2-dev-app -o json | jq -r '.items[0].metadata.name'

.PHONY: project-2-dev-app-port-forward
PROJECT_2_DEV_APP_POD=$(shell make project-2-dev-app-pod-get)
project-2-dev-app-port-forward: ## Forward port from project-2/dev-app to localhost
	@kubectl --namespace $(NAMESPACE) port-forward pods/$(PROJECT_2_DEV_APP_POD) 11180:8080

.PHONY: project-2-dev-app-check-login
project-2-dev-app-check-login: ## Check if we can log into vault
# TODO: check on windows and mac os
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-2__dev-ro\",\"jwt\":\"$(shell make project-2-dev-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'
