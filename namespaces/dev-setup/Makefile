NAMESPACE=dev-setup

.PHONY: namespace-create
namespace-create: ## Create namespace $(NAMESPACE)
	@kubectl --namespace $(NAMESPACE) apply -f namespace.yml
	@$(MAKE) secret-create
	@$(MAKE) project-1-pgsql-sa-create
	@$(MAKE) project-2-dev-sa-create

.PHONY: project-1-pgsql-sa-create
project-1-pgsql-sa-create: ## Create sa for project-1/pgsql
	@kubectl --namespace $(NAMESPACE) apply -f project-1-pgsql-sa.yml

.PHONY: project-2-dev-sa-create
project-2-dev-sa-create: ## Create sa for project-1/pgsql
	@kubectl --namespace $(NAMESPACE) apply -f project-2-dev-sa.yml

.PHONY: namespace-delete
namespace-delete: ## Delete namespace $(NAMESPACE)
	@kubectl --namespace $(NAMESPACE) delete -f namespace.yml

.PHONY: namespace-recreate
namespace-recreate: ## Recreate namespace $(NAMESPACE)
	@$(MAKE) namespace-delete
	@$(MAKE) namespace-create

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
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-1__dev\",\"jwt\":\"$(shell make default-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'

########################################################################################################################
#### project-1 #### stage ##############################################################################################
########################################################################################################################

.PHONY: project-1-stage-app-deploy
project-1-stage-app-deploy: ## Deploy basic stage-app
	@kubectl --namespace $(NAMESPACE) apply -f project-1__stage-app

.PHONY: project-1-stage-app-rollout
project-1-stage-app-rollout: ## Rollout restart stage-app
	@kubectl --namespace $(NAMESPACE) rollout restart deployments/project-1-stage-app

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
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-1__stage\",\"jwt\":\"$(shell make default-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'

########################################################################################################################
#### project-1 #### pgsql ################################################################################################
########################################################################################################################

.PHONY: project-1-pgsql-deploy
project-1-pgsql-deploy: ## Deploy basic pgsql
	@kubectl --namespace $(NAMESPACE) apply -f project-1__pgsql

.PHONY: project-1-pgsql-rollout
project-1-pgsql-rollout: ## Rollout restart pgsql
	@kubectl --namespace $(NAMESPACE) rollout restart deployments/project-1-pgsql

.PHONY: project-1-pgsql-delete
project-1-pgsql-delete: ## Delete pgsql
	@kubectl --namespace $(NAMESPACE) delete -f project-1__pgsql

.PHONY: project-1-pgsql-pod-get
project-1-pgsql-pod-get: ## Get id of vault pod
	@kubectl --namespace $(NAMESPACE) get pods --selector app=project-1-pgsql -o json | jq -r '.items[0].metadata.name'

.PHONY: project-1-pgsql-port-forward
PROJECT_1_PGSQL_POD=$(shell make project-1-pgsql-pod-get)
project-1-pgsql-port-forward: ## Forward port from vault to localhost
	@kubectl --namespace $(NAMESPACE) port-forward pods/$(PROJECT_1_PGSQL_POD) 15432:5432

.PHONY: project-1-pgsql-sa-secret-name-get
project-1-pgsql-sa-secret-name-get: ## Show project-1-pgsql sa secret name
	@kubectl --namespace $(NAMESPACE) get serviceaccounts/project-1-pgsql -o json | jq -r '.secrets[0].name'

.PHONY: project-1-pgsql-sa-token-get
project-1-pgsql-sa-token-get: ## Get token of project-1-pgsql SA
	@kubectl --namespace $(NAMESPACE) get secret $(shell make project-1-pgsql-sa-secret-name-get) -o json | jq -r '.data.token' | base64 --decode

.PHONY: project-1-pgsql-check-login
project-1-pgsql-check-login: ## Check if we can log into vault
# TODO: check on windows and mac os
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-1__pgsql\",\"jwt\":\"$(shell make project-1-pgsql-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'

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
	@curl -s -D /dev/stderr -X POST -d "{\"role\":\"project-2__dev\",\"jwt\":\"$(shell make project-2-dev-sa-token-get)\"}" http://localhost:8200/v1/auth/kubernetes/login | jq '.auth'
