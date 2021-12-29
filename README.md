# DB8 [dɪˈbeɪt]

A [Kubernetes operator][website-operatorframework] and a collection of [Helm][website-helm] charts for databases.

## Motivation

Since the deprecation of the official Helm chart repository, many users started using the Bitnami Helm charts. The Bitnami Helm charts are of high quality, but have some shortcomings that make them difficult to use from the perspective of an open source enthusiast:

- **Limited compatibility with library containers**  
  Some Helm charts are compatible with the library containers, but additional friction exists due to the usage of non-standard default configuration options. An example for this is the `mariadb` helm chart, which uses the `/bitnami/mariadb` path to store its data rather than the default `/var/lib/mysql`.

- **Limited support for ARM**  
  As highlighted in [this issue][issue-bitnami-arm64], the Bitnami charts lack support for `arm64v8`.

- **Proprietary build process**  
  The Helm charts heavily rely on Bitnami's container images, which are built in a proprietary process. This limits the ability of open source contributors to help fixing the aforementioned issues. From a software lifecycle perspective this creates similar constraint as other proprietary software and solutions.

I am a major advocate for free and open source software, where **everything** is — well — free and open source. Are we paying or otherwise supporting open source contributors enough? Heck no! Is this fair towards the contributors? Of course not! Turning open source software into non-free or proprietary software, like it happened with [Inlets][website-inlets] or [KubeDB][website-kubedb] is in my humble opinion not the solution, although I acknowledge the following issues:

- **Exploitation of open source software**  
  Major software companies — not to name them specifically but you may know whom I am talking about — are exploiting open source software by providing very profitable services built on top of free software without contributing back either financially or intellectually. [MongoDB][website-mongodb], [Inlets][website-inlets] and [KubeDB][website-kubedb] fell victim to this and were thereby forced to take action by either adjusting their license or their business model to prevent this.

- **Lack of a well-defined business model**  
  Although the exploitation of these projects is a major issue. The lack of an appropriate business model for open source companies or individuals is also a reason why this happens. Some companies are better at commercializing open source, such as [Rancher][website-rancher], while others do not succeed.

Nice rant, Nicklas, so what can we as an open source community do?

- **Sponsor the contributors monthly**  
  Use GitHub sponsors to contribute financially to the projects you are using. If you are a company and you use open source projects, it **should be a no-brainer to allocate funding in your budget**. It's about ethics, but if you want a more commercial argument, it is about **sustainability**. Your software supply chain needs maintenance and maintenance is not free. You simply cannot believe that exploitation is sustainable. For the colonialization in the past we all paying the price today. We now have to deal with fascism, racism, child mortality, discrimination, polical instability, war, resource shortages and depressing knowledge that most if not all of the injustice and pain in the world is caused by exploitation. And also today we are experiencing the same things all over again with our environment — exploitation for which we will eventually pay the price.

- **Allocate time to contribute**  
  If you don't like the idea of supporting the coffee addictions of our valued open source heros, allocate a **fixed, non-negotiable** amount of time for contributions to open source projects. This does not only entail the development of features, which will benefit you, but can also be things like process improvements to make release pipelines more efficient or by triaging issues. This time is **not for you** but to take something of our open source heros' plates.

If you feel bad now, I get it. The truth hurts. Use this energy to do either of the outlined options above.

## Architecture

This repository contains a collection of Helm charts, where each Helm chart must be compatible with the library container images provided by the [Docker Hub][website-docker-hub]. These Helm charts are then used by the Operator to install the databases. As this could also be done by using a Helm chart operator, the following features are provided by this operator:

- **Credential-injection via annotations**  
  You can create annotations on `Deployments`, `DaemonSets` or `StatefulSets` to automatically generate credentials and inject them into the container

- **Quotas**  
  Manage access to databases via per-instance quotas that define which application can use which instance.

- **Automated backup and recovery via annotations**  
  Automatically back up databases and recover to previous snapshots.

## License

This project is licensed under the terms of the [MIT license][file-license].

[website-operatorframework]: https://operatorframework.io/about
[website-helm]: https://helm.sh
[website-docker-hub]: https://hub.docker.com
[issue-bitnami-arm64]: https://github.com/bitnami/charts/issues/7305
[website-mongodb]: https://mongodb.com
[website-inlets]: https://inlets.dev
[website-kubedb]: https://kubedb.com
[website-rancher]: https://rancher.com
[file-license]: ./LICENSE.md
