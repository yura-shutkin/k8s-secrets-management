.PHONY: start-minikube
start-minikube: ## Start minikube
	@minikube start --nodes=3 --insecure-registry "10.0.0.0/24"
	@minikube addons enable registry

.PHONY: start-dashboard
start-dashboard: ## Start dashboard
	@minikube dashboard &
