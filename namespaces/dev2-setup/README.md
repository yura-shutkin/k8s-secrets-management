# Dev setup namespace

## How to set
- Run `cd ../dev2-setup`
- Run `make namespace-create`
- To check authentication run `make project-1-dev-app-check-login`
- To check authentication run `make project-1-stage-app-check-login`
- To check authentication run `make project-1-pgsql-check-login`
- Run `make project-1-pgsql-deploy`
- Run `make project-1-pgsql-vault-setup`
- Run `make project-1-pgsql-port-forward`
  <!-- Docker can't connect on mac, may be on Windows too -->
- On linux run `docker run --rm -ti -e 'PGPASSWORD=S3cr3t' --network=host postgres:12-alpine psql -h 0.0.0.0 -U project -p 15433 -d project_db -c 'SELECT NOW();'`
- Run `make project-1-dev-app-deploy`
- Run `make project-1-dev-app-port-forward`
- Visit http://localhost:11080
- Run `make project-1-stage-app-deploy`
- Run `make project-1-stage-app-port-forward`
- Visit http://localhost:11081

---

### Pluses and minuses

- No application restart on dynamic secret upgrade
- Easier to see secrets
- Sidecar container

---

### Documentation

- [Vault Agent](https://learn.hashicorp.com/tutorials/vault/agent-kubernetes)