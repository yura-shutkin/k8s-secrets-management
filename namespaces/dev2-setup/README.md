# Dev setup namespace

## How to set
- Run `cd ../dev2-setup`
- Run `make namespace-create`
- Run `make secret-create`
- To check authentication run `make project-1-dev-app-check-login`
- To check authentication run `make project-1-stage-app-check-login`
- To check authentication run `make project-1-pgsql-check-login`
- To check authentication run `make project-2-dev-app-check-login`
- Encrypt string for dev env [project-1/transit/dev](http://localhost:8200/ui/vault/secrets/project-1%2Ftransit/actions/dev?action=encrypt)
- Put encrypted string into [project-1__dev-app/deployment.yml]()
- Encrypt string for dev env [project-1/transit/stage](http://localhost:8200/ui/vault/secrets/project-1%2Ftransit/actions/stage?action=encrypt)
- Put encrypted string into [project-1__dev-stage/deployment.yml]()
- Run `make project-1-pgsql-deploy`
- Run `make project-1-pgsql-vault-setup`
- Run `make project-1-pgsql-port-forward`
  <!-- Docker can't connect on mac, may be on Windows too -->
- On linux run `docker run --rm -ti -e 'PGPASSWORD=S3cr3t' --network=host postgres:12-alpine psql -h 0.0.0.0 -U project -p 15432 -d project_db -c 'SELECT NOW();'`
- Run `make project-1-dev-app-deploy`
- Run `make project-1-dev-app-port-forward`
- Visit http://localhost:11080
- Run `make project-1-stage-app-deploy`
- Run `make project-1-stage-app-port-forward`
- Visit http://localhost:11081
- Run `make project-2-dev-app-deploy`
- Run `make project-2-dev-app-port-forward`
- Visit http://localhost:11180 

---

### Pluses and minuses

- Go to deployment page [project-1/dev-app](http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/deployment/dev2-setup/project-1-dev-app?namespace=dev2-setup)
- New Replica Set
- Pods
- Exec into pod
- Run `env`
- Run `/vault/vault-env env | grep FROM_VAULT_`
- Run `VAULT_TOKEN=vault:login /vault/vault-env env | grep VAULT_TOKEN`
- Run `FROM_MY_VAR='vault:project-1/kv/data/dev/db_creds#db_name#1' /vault/vault-env env | grep FROM_VAULT_`
- Run `FROM_MY_VAR='vault:project-1/kv/data/stage/token#value#1' /vault/vault-env env | grep FROM_VAULT_` 

---

- Next step: [prod namespace](../prod-setup)

### Documentation
<!-- - https://learn.hashicorp.com/tutorials/vault/agent-kubernetes -->

- [Vault Secrets Webhook annotations](https://banzaicloud.com/docs/bank-vaults/mutating-webhook/annotations/)