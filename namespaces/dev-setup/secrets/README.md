### Documentation
- https://kubernetes.io/docs/tasks/configmap-secret/managing-secret-using-config-file/

### How to create secret in K8S

- first, you need to encode string in base64

    ```bash
    $ echo -n 'admin' | base64
    YWRtaW4=
    ```

- next, you can update [db-creds.yml](./db-creds.yml) or create a new file
