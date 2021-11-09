# Prod setup namespace

## How to set

- Run `cd ../prod-setup`
- Run `make namespace-create`
- To check authentication run `make project-1-prod-app-check-login`
- Run `make project-1-prod-app-deploy`
- Run `make project-1-prod-app-port-forward`
- Visit http://localhost:12080

---

- Next step: [vault namespace](../vault)

### Documentation
<!-- - https://learn.hashicorp.com/tutorials/vault/agent-kubernetes -->

- [Vault Secrets Webhook annotations](https://banzaicloud.com/docs/bank-vaults/mutating-webhook/annotations/)