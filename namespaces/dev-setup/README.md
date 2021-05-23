# Dev setup namespace

## How to set

- Run `cd ../dev-setup`
- Run `make namespace-create`
- Run `make secret-create`
- To check authentication run `make project-1-dev-app-check-login`
- To check authentication run `make project-1-stage-app-check-login`
- To check authentication run `make project-1-pgsql-check-login`
- To check authentication run `make project-2-dev-app-check-login`
- Run `make project-1-dev-app-deploy`
- Run `make project-1-dev-app-port-forward`
- Visit http://localhost:11080
- Run `make project-1-stage-app-deploy`
- Run `make project-1-stage-app-port-forward`
- Visit http://localhost:11081
- Run `make project-2-dev-app-deploy`
- Run `make project-2-dev-app-port-forward`
- Visit http://localhost:11180
- Run `make project-1-pgsql-deploy`
- Run `make project-1-pgsql-port-forward`
<!-- Check on mac os and windows -->
- Run `docker run --rm -ti -e 'PGPASSWORD=S3cr3t' --network=host postgres:12-alpine psql -h 0.0.0.0 -U project -p 15432 -d project_db -c 'SELECT NOW();'` 

---

- Next step: [prod namespace](../prod-setup)

### Documentation
<!-- - https://learn.hashicorp.com/tutorials/vault/agent-kubernetes -->

- [Vault Secrets Webhook annotations](https://banzaicloud.com/docs/bank-vaults/mutating-webhook/annotations/)