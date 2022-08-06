registry_addr="localhost:5000"
arch=ubuntu
go_ver="1.18.2"
version=$(shell cat version)
platform="linux/amd64"

.PHONY: build
build: ## Build app image
	@docker build --rm -t "$(registry_addr)/web-app:$(version)-$(arch)" --build-arg=$(go_ver) -f docker/Dockerfile-$(arch) .

.PHONY: buildx
buildx: ## Build app image
	@docker buildx build --platform $(platform) --rm -t "$(registry_addr)/web-app:$(version)-$(arch)" --build-arg=$(go_ver) -f docker/Dockerfile-$(arch) .

.PHONY: push
push: ## Push image into registry
	@docker push "$(registry_addr)/web-app:$(version)-$(arch)"

.PHONY: build-n-push
build-n-push: ## Build app image and push into registry
	@$(MAKE) build
	@$(MAKE) push
