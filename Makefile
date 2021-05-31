NAMESPACE=kubernetes-dashboard
PREP_NS=prepare

.PHONY: minikube-start
NODES=1
CPUS=2
MEMORY=2G
minikube-start: ## Start minikube
	@minikube start --nodes=$(NODES) --cpus=$(CPUS) --memory=$(MEMORY)

.PHONY: dashboard-deploy
dashboard-deploy: ## Start dashboard
	@ # Deploy dahsboard
#	@kubectl apply --namespace=$(NAMESPACE) -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.2.0/aio/deploy/recommended.yaml
	@kubectl apply --namespace=$(NAMESPACE) -f dashboard/dashboard.yml
	@ # Create admin SA for access into dashboard
	@kubectl apply --namespace=$(NAMESPACE) -f dashboard/admin-sa.yml
	@kubectl apply --namespace=$(NAMESPACE) -f dashboard/admin-cluster-role-binding.yml

.PHONY: minikube-stop
minikube-stop: ## Stop minikube
	@minikube stop

.PHONY: minikube-delete-profile
minikube-delete-profile: ## Stop minikube
	@minikube delete

.PHONY: minikube-recreate
minikube-recreate: ## Recreate minikube
	@$(MAKE) minikube-delete-profile
	@$(MAKE) minikube-start

.PHONY: admin-sa-secret-name-get
admin-sa-secret-name-get: ## Show default sa secret name
	@kubectl --namespace $(NAMESPACE) get serviceaccounts/admin-user -o json | jq -r '.secrets[0].name'

.PHONY: admin-sa-token-get
admin-sa-token-get: ## Get token of default-vault-secrets-webhook SA
	@kubectl --namespace $(NAMESPACE) get secret $(shell make admin-sa-secret-name-get) -o json | jq -r '.data.token' | base64 --decode
	@echo
	@echo
	@echo "http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login"

.PHONY: proxy
proxy: ## Proxy k8s api on localhost
	@kubectl proxy --port=8001

.PHONY: prepare
prepare: ## Download images that will be used in a future
	@kubectl --namespace $(PREP_NS) apply -f prepare/namespace.yml
	@kubectl --namespace $(PREP_NS) apply -f prepare/jobs.yml
	@$(MAKE) -C namespaces/bank-vaults webhook-update

.PHONY: prepare-delete
prepare-delete: ## Download images that will be used in a future
	@kubectl --namespace $(PREP_NS) delete -f prepare/namespace.yml

.PHONY: events-watch
VERBOSE=3
events-watch: ## Stream events for every namespece
	@kubectl get events --watch --all-namespaces --v=$(VERBOSE)

.PHONY: pods-watch
pods-watch: ## Show status of pods
	@watch -n 1 kubectl --all-namespaces=true get pods
