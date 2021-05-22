# Errors

## Service account name is not authorized

* To read logs execute `kubectl --namespace vault logs --follow=true -l app=vault --all-containers=true`

  ```json
  {
    "time": "2021-03-08T16:04:53.226175278Z",
    "type": "response",
    "auth": {
      "token_type": "default"
    },
    "request": {
      "id": "289e0a1b-9842-6dd3-b168-f6ed48c01277",
      "operation": "update",
      "mount_type": "kubernetes",
      "namespace": {
        "id": "root"
      },
      "path": "auth/kubernetes/login",
      "data": {
        "jwt": "hmac-sha256:a0663610fca2fc30040edd4025a47328b2722b01cb08591bc8bd11c5c0233203",
        "role": "hmac-sha256:58d9129ebbee1f530b48fefa991413958c89f631a85ce82f092ef763315b0c02"
      },
      "remote_address": "10.244.2.4"
    },
    "response": {
      "mount_type": "kubernetes"
    },
    "error": "service account name not authorized"
  }
  ```

* To see which JWT and role are used in queries execute 

  ```bash
  curl -sX POST -H "X-Vault-Token: 12345" -d '{"audit_non_hmac_request_keys": ["role", "jwt"]}' $(minikube --namespace vault service --url vault)/v1/sys/auth/kubernetes/tune
  ```

  * Then you can see in vault logs

  ```json
  {
    "time": "2021-03-08T16:22:46.429096507Z",
    "type": "response",
    "auth": {
      "token_type": "default"
    },
    "request": {
      "id": "a95b4032-b2ca-294f-7073-24b46c3a5990",
      "operation": "update",
      "mount_type": "kubernetes",
      "namespace": {
        "id": "root"
      },
      "path": "auth/kubernetes/login",
      "data": {
        "jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6IkcyNFdLeXVSUndsSnNnaDZKOEVLTFEtbHFJeEFxeVg3dnk0ZndybDY4cUkifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJiYXNpYy1zZXR1cCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0LXRva2VuLXRqZ2hnIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJhYmIxYzM5Ni1kNWYzLTRiZDEtODViNC0zZDNmMDZlMjNhM2YiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6YmFzaWMtc2V0dXA6ZGVmYXVsdCJ9.ihpOR4UeGPo-w1_RNTWL29wCDF_j330eOnc0wKPA5ndlwRIznoL1kJ_9YXIU8R74wQMjX7hGxIHlca45IhWUJLx98LLrkk6Ppdiz7UrteZWRV72vv43j0iBBcUtIvqFvpFAAwl2UR9RyRh5eqk3MjoJXHUDtQGjawbKt6ayXE_Dn8-4mudxFiT69mEnMC3R60b_A5o26AbRfrb7rFqWYiy_NVjQ6kX8dYl5tDE8RrX8kfjFZVfdMXu7PIlk9QnIof9Nvy-EUBk6xkZxkBP4Q5tVPdxuDNr_RqQ5u6XoEwGydZ_HQASR2MZ4nP-Xfq2CHNVte3ZiN5Xs54mBEb3QUuQ",
        "role": "dev-ro"
      },
      "remote_address": "10.244.2.6"
    },
    "response": {
      "mount_type": "kubernetes"
    },
    "error": "service account name not authorized"
  }
  ```

## Permission denied

```json
{
  "time": "2021-03-09T05:02:16.552571791Z",
  "type": "response",
  "auth": {
    "token_type": "default"
  },
  "request": {
    "id": "9f01629b-8472-0591-7542-ef47e1f0aa05",
    "operation": "update",
    "mount_type": "kubernetes",
    "namespace": {
      "id": "root"
    },
    "path": "auth/kubernetes/login",
    "data": {
      "jwt": "hmac-sha256:a0663610fca2fc30040edd4025a47328b2722b01cb08591bc8bd11c5c0233203",
      "role": "hmac-sha256:58d9129ebbee1f530b48fefa991413958c89f631a85ce82f092ef763315b0c02"
    },
    "remote_address": "10.244.2.12"
  },
  "response": {
    "mount_type": "kubernetes"
  },
  "error": "permission denied"
}
```
This error may appear when role binding is incorrect

```
2021-03-09T05:02:16.552Z [ERROR] auth.kubernetes.auth_kubernetes_c964001c: login unauthorized due to: {"kind":"Status","apiVersion":"v1","metadata":{},"status":"Failure","message":"tokenreviews.authentication.k8s.io is forbidden: User \"system:serviceaccount:basic-setup:default\" cannot create resource \"tokenreviews\" in API group \"authentication.k8s.io\" at the cluster scope","reason":"Forbidden","details":{"group":"authentication.k8s.io","kind":"tokenreviews"},"code":403}
```

Check [clusterRoleBinding.yml](/bank-vaults/rbac/clusterRoleBinding.yml)

## Message authentication failed

This message can be frustrating, but this message means you have issue with decrypting message described in [basic-setup/web-app/deployment.yml](/basic-setup/web-app/deployment.yml)

```json
{
  "time": "2021-03-09T05:07:10.363616083Z",
  "type": "response",
  "auth": {
    "client_token": "hmac-sha256:77b88698eed718b2637d248436f2db62e1aae13aaa650c3d2f1a5c7bd4352d9c",
    "accessor": "hmac-sha256:67f482f745ddb131ad8f9779766abde20a243657e80742227709db1c9429f968",
    "display_name": "kubernetes-basic-setup-default",
    "policies": [
      "default",
      "dev_key_decrypt",
      "dev_ro"
    ],
    "token_policies": [
      "default",
      "dev_key_decrypt",
      "dev_ro"
    ],
    "metadata": {
      "role": "dev-ro",
      "service_account_name": "default",
      "service_account_namespace": "basic-setup",
      "service_account_secret_name": "default-token-tjghg",
      "service_account_uid": "abb1c396-d5f3-4bd1-85b4-3d3f06e23a3f"
    },
    "entity_id": "6a50f740-09be-555c-e810-04bcff62bc90",
    "token_type": "service",
    "token_ttl": 60,
    "token_issue_time": "2021-03-09T05:07:06Z"
  },
  "request": {
    "id": "f190ad9b-8020-5853-a3bb-8140f2355151",
    "operation": "update",
    "mount_type": "transit",
    "client_token": "hmac-sha256:77b88698eed718b2637d248436f2db62e1aae13aaa650c3d2f1a5c7bd4352d9c",
    "client_token_accessor": "hmac-sha256:67f482f745ddb131ad8f9779766abde20a243657e80742227709db1c9429f968",
    "namespace": {
      "id": "root"
    },
    "path": "project/transit/decrypt/dev",
    "data": {
      "ciphertext": "hmac-sha256:6f46e8eb12bb606ab1322623effcf3d4ec8fd3cf49f4d969028de5abfe5c1235"
    },
    "remote_address": "10.244.1.5"
  },
  "response": {
    "mount_type": "transit"
  },
  "error": "1 error occurred:\n\t* cipher: message authentication failed\n\n"
}
```

## Namespace not found

Check if namespace exist
```
Error from server (NotFound): namespaces "basic-setup" not found
HTTP/1.1 500 Internal Server Error
Cache-Control: no-store
Content-Type: application/json
Date: Thu, 20 May 2021 05:15:51 GMT
Content-Length: 52

{"errors":["could not load backend configuration"]}
```

### Could not load backend configuration

Check if kubernetes auth backend setup properly
```
HTTP/1.1 500 Internal Server Error
Cache-Control: no-store
Content-Type: application/json
Date: Thu, 20 May 2021 05:16:07 GMT
Content-Length: 52

{"errors":["could not load backend configuration"]}
```
