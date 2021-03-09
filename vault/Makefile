.PHONY: create-namespace
create-namespace: ## Create namespace basic-setup
	@kubectl --namespace vault apply -f namespace.yml

.PHONY: deploy-vault
deploy-vault: ## Deploy vault
	@kubectl --namespace vault apply -f deployment.yml -f service.yml
	@minikube --namespace vault service --url vault

.PHONY: setup-vault
VAULT_TOKEN=12345
VAULT_URL=$(shell minikube -n vault service --url vault)
CA_CERT=$(shell cat ~/.minikube/ca.crt | tr '\n' ',' | sed 's|,|\\n|g')
setup-vault: ## Create necessary resources in vault
	# Enable audit device
	-@curl -X PUT -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/audit-device.json "$(VAULT_URL)/v1/sys/audit/stdout"
	# Mount project kv
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/kv.json "$(VAULT_URL)/v1/sys/mounts/project/kv"
	# Mount transit backend
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/transit.json "$(VAULT_URL)/v1/sys/mounts/project/transit"
	# Enable kubernetes auth backend
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/k8s-auth-backend.json "$(VAULT_URL)/v1/sys/auth/kubernetes"
	# Configure k8s auth backend
	# TODO: jwt reviewer token is missed in this configuration
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d "{\"kubernetes_host\": \"https://10.96.0.1\", \"kubernetes_ca_cert\":\"$(CA_CERT)\", \"disable_local_ca_jwt\": \"true\"}" "$(VAULT_URL)/v1/auth/kubernetes/config"
	# Create or update k8s auth roles
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/k8s-dev-role.json "$(VAULT_URL)/v1/auth/kubernetes/role/dev-ro"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/k8s-stage-role.json "$(VAULT_URL)/v1/auth/kubernetes/role/stage-ro"
	# Create or update policies
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/dev-ro-policy.json "$(VAULT_URL)/v1/sys/policies/acl/dev_ro"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/dev-key-decrypt.json "$(VAULT_URL)/v1/sys/policies/acl/dev_key_decrypt"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/stage-ro-policy.json "$(VAULT_URL)/v1/sys/policies/acl/stage_ro"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/stage-key-decrypt.json "$(VAULT_URL)/v1/sys/policies/acl/stage_key_decrypt"
	# Put secret into project/dev/token
	@ # Prevent rising of index. So we have 1 and 2 versions every time
	@ # web-app
	@curl -X DELETE -H "X-Vault-Token: $(VAULT_TOKEN)" "$(VAULT_URL)/v1/project/kv/metadata/dev/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v1.json "$(VAULT_URL)/v1/project/kv/data/dev/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v2.json "$(VAULT_URL)/v1/project/kv/data/dev/token"
	@ # Second app
	@curl -X DELETE -H "X-Vault-Token: $(VAULT_TOKEN)" "$(VAULT_URL)/v1/project/kv/metadata/stage/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v2.json "$(VAULT_URL)/v1/project/kv/data/stage/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v1.json "$(VAULT_URL)/v1/project/kv/data/stage/token"
	# Create dev key
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/dev-key.json "$(VAULT_URL)/v1/project/transit/keys/dev"
	# Create stage key
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/stage-key.json "$(VAULT_URL)/v1/project/transit/keys/stage"
	# Done
