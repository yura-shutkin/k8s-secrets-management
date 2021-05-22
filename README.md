# Secrets for apps deployed into K8S

This repo is set of examples and prototypes for secrets management

## Requirements

- [minikube](https://minikube.sigs.k8s.io/docs/start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [helm](https://helm.sh/docs/intro/install/)
- [docker](https://docs.docker.com/get-docker/)
- [jq](https://stedolan.github.io/jq/download/)
- curl
- <details>
  <summary>make</summary>

  - linux:
    - Depends on distro
  - mac os:
    - `brew install make`
  - windows:
    <!-- TODO: check instruction :arrow_down: -->
    - [instruction](http://gnuwin32.sourceforge.net/packages/make.htm)
  </details>
- <details>
  <summary>sed</summary>
  
  - linux:
    - You are OK
  - mac os:
    - [instruction](https://medium.com/@bramblexu/install-gnu-sed-on-mac-os-and-set-it-as-default-7c17ef1b8f64)
  - windows:
    <!-- TODO: check instruction :arrow_down: -->
    - [instruction](http://gnuwin32.sourceforge.net/packages/sed.htm)
  </details>

## Documentation

- [errors and how to fix them](docs/errors.md)
- [TODO](TODO.md)
- [Mutating webhook configuration examples](https://banzaicloud.com/docs/bank-vaults/mutating-webhook/configuration/)
- [Mutating webhook annotations](https://banzaicloud.com/docs/bank-vaults/mutating-webhook/annotations/)

## How to start

- Run `make minikube-start`
- Run `make dashboard-deploy`
- Run `make admin-sa-token-get`
- Run `make proxy`
- Open [dashboard](http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login)
---
- [Deploy bank vault Vault Secrets Webhook](namespaces/bank-vaults/README.md)