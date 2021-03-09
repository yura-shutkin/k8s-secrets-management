# Secrets for apps deployed into K8S

This repo is set of examples and prototypes for secrets management

## Repo status

---
    Development in progress
---

## Requirements

- [minikube](https://minikube.sigs.k8s.io/docs/start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [helm](https://helm.sh/docs/intro/install/)
- [vault](https://vaultproject.io)
- [docker](https://docs.docker.com/get-docker/)

## Documentation

* [errors and how to fix them](docs/errors.md)
* [TODO](TODO.md)

## How to start

- Run `make start-minikube`
- To start k8s dashboard exec `make start-dashboard`
- Deploy vault
  - `cd vault`
  - `make create namespace`
  - `make deploy-vault`
  - `make setup-vault`
  - `cd ../`
- Deploy Vault Secrets Webhook
  - `cd bank-vaults`
  - `make create-namespace`
  - `make deploy-webhook`
  - `make setup-rbac`
  - `cd ../`
- Deploy application
  - `cd basic-setup`
  - `make create-namespace`
  - `make create-secret`
  - `make update-vault-auth`
  - `make deploy-web-app`
  - `make deploy-second-app`
  - `cd ../`
