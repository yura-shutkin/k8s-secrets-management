## Useful commands

- `kubectl get events --watch --all-namespaces`
- `kubectl --namespace basic-setup get events --watch`
- `kubectl --namespace vault logs --follow=true -l app=vault --all-containers=true --since=1m`
- `kubectl --namespace dev-setup get events --watch`
- `minikube docker-env`
