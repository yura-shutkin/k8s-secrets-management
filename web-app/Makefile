registry_addr="localhost:5000"
arch=ubuntu
version=$(shell cat version)

.PHONY: build-n-push
build-n-push: ## Build app image and push into registry
	@docker build --rm -t "$(registry_addr)/web-app:$(version)-$(arch)" -f $(arch)/Dockerfile .
	@docker push "$(registry_addr)/web-app:$(version)-$(arch)"
