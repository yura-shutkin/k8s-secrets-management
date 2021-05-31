NAMESPACE="vault"

.PHONY: namespace-create
namespace-create: ## Create namespace basic-setup
	@kubectl --namespace $(NAMESPACE) apply -f namespace.yml

.PHONY: namespace-delete
namespace-delete: ## Delete namespace $(NAMESPACE)
	@kubectl --namespace $(NAMESPACE) delete -f namespace.yml

.PHONY: namespace-recreate
namespace-recreate: ## Recreate namespace $(NAMESPACE)
	@$(MAKE) namespace-delete
	@$(MAKE) namespace-create

.PHONY: vault-deploy
vault-deploy: ## Deploy vault
	@kubectl --namespace $(NAMESPACE) apply -f deployment.yml -f service.yml

.PHONY: vault-setup
VAULT_TOKEN=12345
VAULT_URL=http://localhost:8200
JWT=$(shell make sa-token-get)
CA_CRT=$(shell make sa-cert-get)
vault-setup: ## Create necessary resources in vault
	# Enable audit device
	-@curl -X PUT -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/audit-device.json "$(VAULT_URL)/v1/sys/audit/stdout"
	# Mount project-1 kv
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__kv.json "$(VAULT_URL)/v1/sys/mounts/project-1/kv"
	# Mount project-2 kv
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-2__kv.json "$(VAULT_URL)/v1/sys/mounts/project-2/kv"
	# Mount project-1 transit backend
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__transit.json "$(VAULT_URL)/v1/sys/mounts/project-1/transit"
	# Enable kubernetes auth backend
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/k8s-auth-backend.json "$(VAULT_URL)/v1/sys/auth/kubernetes"
	# Configure k8s auth backend
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d "{\"kubernetes_host\": \"https://10.96.0.1\", \"kubernetes_ca_cert\":\"$(CA_CRT)\", \"disable_local_ca_jwt\": \"true\", \"token_reviewer_jwt\":\"$(JWT)\"}" "$(VAULT_URL)/v1/auth/kubernetes/config"
	# Create or update k8s auth roles
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__k8s-role__dev.json "$(VAULT_URL)/v1/auth/kubernetes/role/project-1__dev"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__k8s-role__stage.json "$(VAULT_URL)/v1/auth/kubernetes/role/project-1__stage"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__k8s-role__prod.json "$(VAULT_URL)/v1/auth/kubernetes/role/project-1__prod"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__k8s-role__pgsql.json "$(VAULT_URL)/v1/auth/kubernetes/role/project-1__pgsql"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-2__k8s-role__dev.json "$(VAULT_URL)/v1/auth/kubernetes/role/project-2__dev"
	# Create or update policies
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__policy__dev-ro.json "$(VAULT_URL)/v1/sys/policies/acl/project-1__dev-ro"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__policy__dev_db-creds_root.json "$(VAULT_URL)/v1/sys/policies/acl/project-1__dev_db-creds_root"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__policy__dev_db-creds-ro-gen.json "$(VAULT_URL)/v1/sys/policies/acl/project-1__dev_db-creds_ro-gen"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__policy__dev-key-decrypt.json "$(VAULT_URL)/v1/sys/policies/acl/project-1__dev-key-decrypt"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__policy__stage-ro.json "$(VAULT_URL)/v1/sys/policies/acl/project-1__stage-ro"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__policy__stage-key-decrypt.json "$(VAULT_URL)/v1/sys/policies/acl/project-1__stage-key-decrypt"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__policy__prod-ro.json "$(VAULT_URL)/v1/sys/policies/acl/project-1__prod-ro"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-2__policy__dev-ro.json "$(VAULT_URL)/v1/sys/policies/acl/project-2__dev-ro"
	# Put secret into project/dev/token
	@ # Prevent rising of index. So we have 1 and 2 versions every time
	@ # project-1/dev-app
	@curl -X DELETE -H "X-Vault-Token: $(VAULT_TOKEN)" "$(VAULT_URL)/v1/project-1/kv/metadata/dev/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v1.json "$(VAULT_URL)/v1/project-1/kv/data/dev/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v2.json "$(VAULT_URL)/v1/project-1/kv/data/dev/token"
	@ # project-1/stage-app
	@curl -X DELETE -H "X-Vault-Token: $(VAULT_TOKEN)" "$(VAULT_URL)/v1/project-1/kv/metadata/stage/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v2.json "$(VAULT_URL)/v1/project-1/kv/data/stage/token"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/token_v1.json "$(VAULT_URL)/v1/project-1/kv/data/stage/token"
	@curl -X DELETE -H "X-Vault-Token: $(VAULT_TOKEN)" "$(VAULT_URL)/v1/project-1/kv/metadata/dev/db_creds"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__db_creds_v1.json "$(VAULT_URL)/v1/project-1/kv/data/dev/db_creds"
	@ # project-1/prod-app
	@curl -X DELETE -H "X-Vault-Token: $(VAULT_TOKEN)" "$(VAULT_URL)/v1/project-1/kv/metadata/prod/db_creds"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__db_creds_v1.json "$(VAULT_URL)/v1/project-1/kv/data/prod/db_creds"
	@ # project-2/dev-app
	@curl -X DELETE -H "X-Vault-Token: $(VAULT_TOKEN)" "$(VAULT_URL)/v1/project-2/kv/metadata/dev/db_creds"
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-2__db_creds_v1.json "$(VAULT_URL)/v1/project-2/kv/data/dev/db_creds"
	# Create project-1/dev key
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__key__dev.json "$(VAULT_URL)/v1/project-1/transit/keys/dev"
	# Create project-1/stage key
	@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__key__stage.json "$(VAULT_URL)/v1/project-1/transit/keys/stage"
	# Enable database SE for project-1
	@ # Can be enabled on second attempt. So ignoring errors
	-@curl -X POST -H "X-Vault-Token: $(VAULT_TOKEN)" -d @payloads/project-1__database.json "$(VAULT_URL)/v1/sys/mounts/project-1/database"
	# Done

.PHONY: vault-pod-get
vault-pod-get: ## Get id of vault pod
	@kubectl --namespace $(NAMESPACE) get pods --selector app=vault -o json | jq -r '.items[0].metadata.name'

.PHONY: vault-port-forward
VAULT_POD=$(shell make vault-pod-get)
vault-port-forward: ## Forward port from vault to localhost
	@kubectl --namespace $(NAMESPACE) port-forward pods/$(VAULT_POD) 8200:8200 &

.PHONY: jwt-in-audit-logs-show
jwt-in-audit-logs-show: ## Show JWT in audit logs
	@curl -sX POST -H "X-Vault-Token: 12345" -d '{"audit_non_hmac_request_keys": ["role", "jwt"]}' http://127.0.0.1:8200/v1/sys/auth/kubernetes/tune

.PHONY: jwt-in-audit-logs-hide
jwt-in-audit-logs-hide: ## Hide JWT in audit logs
	@curl -sX POST -H "X-Vault-Token: 12345" -d '{"audit_non_hmac_request_keys": []}' http://127.0.0.1:8200/v1/sys/auth/kubernetes/tune

.PHONY: sa-secret-name-get
sa-secret-name-get: ## Show default sa secret name
	@kubectl --namespace bank-vaults get serviceaccounts/default-vault-secrets-webhook -o json | jq -r '.secrets[0].name'

.PHONY: sa-token-get
sa-token-get: ## Get token of default-vault-secrets-webhook SA
	@kubectl --namespace bank-vaults get secret $(shell make sa-secret-name-get) -o json | jq -r '.data.token' | base64 --decode

.PHONY: sa-cert-get
sa-cert-get: ## Get CA cert of default-vault-secrets-webhook SA
	@kubectl --namespace bank-vaults get secret $(shell make sa-secret-name-get) -o json | jq -r '.data."ca.crt"' | base64 --decode | sed -z 's/\n/\\n/g'
