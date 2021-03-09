.PHONY: start-minikube
start-minikube: ## Start minikube
	@minikube start --nodes=3 --insecure-registry "10.0.0.0/24"

.PHONY: start-dashboard
start-dashboard: ## Start dashboard
	@minikube dashboard &

.PHONY: stop-minikube
stop-minikube: ## Stop minikube
	@minikube stop

.PHONY: delete-minikube
delete-minikube: ## Stop minikube
	@minikube delete
