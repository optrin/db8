# DB8 [dÉªËbeÉªt] ð¬

A [Kubernetes operator][website-operatorframework] for databases.

`db8` aims to automate the deployment and operation of databases inside a Kubernetes cluster. It was created independently from [KubeDB][website-kubedb] as a free and open source alternative.

## Roadmap â¨

The following features are currently on the roadmap for this operator.

- **Credential injection via annotations** ð  
  You can create annotations on `Deployments`, `DaemonSets` or `StatefulSets` to automatically generate credentials and inject them into the container.

  ```yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: myapp
    annotations:
      mariadb.db8.optrin.io/name: mariadb-prod
  spec:
  # Add your deployment configuration below.
  ```

- **Access control** ð¡ï¸  
  Manage access to databases via rules that define which applications can use which instance.

- **Automated backup and recovery via annotations** ð½  
  Automatically back up databases and recover to previous snapshots.

## Databases ð¾

Below you may find a list of databases and supported features.

| Feature â¨              | MariaDB ð¦­             |
| ----------------------- | --------------------- |
| Deployment ð           | Under construction ð§ |
| Credential injection ð | On roadmap ð¯         |
| Access control ð¡ï¸       | On roadmap ð¯         |
| Automated backup ð½     | On roadmap ð¯         |

## License ð

This project is and will always be licensed under the terms of the [MIT license][file-license].

[website-operatorframework]: https://operatorframework.io/about
[website-kubernetes]: https://kubernetes.io
[website-kubedb]: https://kubedb.com
[file-license]: ./LICENSE.md
