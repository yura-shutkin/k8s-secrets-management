# Vault

## How to deploy

- Run `cd ../vault`
- Run `make namespace-create`
- Run `make vault-deploy`
- Run `make vault-port-forward`
- Run `make vault-setup`

---

- Now Vault set, go into [dev-setup](../dev-setup/README.md)

## Hashicorp Vault injector

- Run `make vault-injector-install`

## Useful locations

| Location | Purpose |
| --- | --- |
| [/ui/vault/access/kubernetes/item/role](http://localhost:8200/ui/vault/access/kubernetes/item/role) | List of configured kubernetes roles |

## Useful links

- [chart values](https://github.com/hashicorp/vault-helm/blob/master/values.yaml)
- [sidecar injector](https://learn.hashicorp.com/tutorials/vault/kubernetes-sidecar?in=vault/kubernetes)
- [csi provider](https://learn.hashicorp.com/tutorials/vault/kubernetes-secret-store-driver?in=vault/kubernetes)
- [vault agent](https://learn.hashicorp.com/tutorials/vault/agent-kubernetes?in=vault/kubernetes)
- [vault injector examples](https://www.vaultproject.io/docs/platform/k8s/injector/examples)
