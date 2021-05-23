NAMESPACE=kubernetes-dashboard

.PHONY: minikube-start
NODES=1
minikube-start: ## Start minikube
	@minikube start --nodes=$(NODES)

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
	@echo -e '\n\n'
	@echo "http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login"

.PHONY: proxy
proxy: ## Proxy k8s api on localhost
	@kubectl proxy --port=8001
