<!---
The rendered version of this file can be found in "./README.md".
Please only edit the template "./README.template.md".
-->
# cert-manager

cert-manager is a Kubernetes addon to automate the management and issuance of
TLS certificates from various issuing sources.

It will ensure certificates are valid and up to date periodically, and attempt
to renew certificates at an appropriate time before expiry.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

Full installation instructions, including details on how to configure extra
functionality in cert-manager can be found in the [installation docs](https://cert-manager.io/docs/installation/kubernetes/).

Before installing the chart, you must first install the cert-manager CustomResourceDefinition resources.
This is performed in a separate step to allow you to easily uninstall and reinstall cert-manager without deleting your installed custom resources.

```bash
$ kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.0.1/cert-manager.crds.yaml
```

To install the chart with the release name `my-cert-manager`:

```console
# Add the jetstack Helm repository
$ helm repo add jetstack https://charts.jetstack.io

# Install the cert-manager helm chart
$ helm install my-cert-manager jetstack/cert-manager -n cert-manager --version=v0.0.1
```

In order to begin issuing certificates, you will need to set up a ClusterIssuer
or Issuer resource (for example, by creating a 'letsencrypt-staging' issuer).

More information on the different types of issuers and how to configure them
can be found in [our documentation](https://cert-manager.io/docs/configuration/).

For information on how to configure cert-manager to automatically provision
Certificates for Ingress resources, take a look at the
[Securing Ingresses documentation](https://cert-manager.io/docs/usage/ingress/).

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Special considerations may be required when upgrading the Helm chart, and these
are documented in our full [upgrading guide](https://cert-manager.io/docs/installation/upgrading/).

**Please check here before performing upgrades!**

## Uninstalling the Chart

To uninstall/delete the `my-cert-manager` deployment:

```console
$ helm delete my-cert-manager
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

If you want to completely uninstall cert-manager from your cluster, you will also need to
delete the previously installed CustomResourceDefinition resources:

```console
$ kubectl delete -f https://github.com/jetstack/cert-manager/releases/download/v0.0.1/cert-manager.crds.yaml
```

## Configuration

The following table lists the configurable parameters of the cert-manager chart and their default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imagePullSecrets` | # Reference to one or more secrets to be used when pulling images # ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/ # | `[]` |
| `global.priorityClassName` | Optional priority class to be used for the cert-manager pods | `""` |
| `global.rbac.create` |  | `true` |
| `global.podSecurityPolicy.enabled` |  | `false` |
| `global.podSecurityPolicy.useAppArmor` |  | `true` |
| `global.logLevel` | Set the verbosity of cert-manager. Range of 0 - 6 with 6 being the most verbose. | `2` |
| `global.leaderElection.namespace` | Override the namespace used to store the ConfigMap for leader election | `"kube-system"` |
| `installCRDs` | DEPRECATED: use components instead! Setting this value to true is the same as adding "crd" to the components list. CRDs will be rendered if "crd" is added to components OR installCRDs is set to true. | `false` |
| `replicaCount` |  | `1` |
| `strategy` |  | `{}` |
| `featureGates` | Comma separated list of feature gates that should be enabled on the controller pod. | `""` |
| `image.repository` |  | `quay.io/jetstack/cert-manager-controller` |
| `image.pullPolicy` | Override the image tag to deploy by setting this variable. If no value is set, the chart's appVersion will be used. tag: canary  Setting a digest will override any tag digest: sha256:0e072dddd1f7f8fc8909a2ca6f65e76c5f0d2fcfb8be47935ae3457e8bbceb20 | `IfNotPresent` |
| `clusterResourceNamespace` | Override the namespace used to store DNS provider credentials etc. for ClusterIssuer resources. By default, the same namespace as cert-manager is deployed within is used. This namespace will not be automatically created by the Helm chart. | `""` |
| `serviceAccount.create` | Specifies whether a service account should be created | `true` |
| `serviceAccount.automountServiceAccountToken` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template name: "" Optional additional annotations to add to the controller's ServiceAccount annotations: {} Automount API credentials for a Service Account. | `true` |
| `extraArgs` | Optional additional arguments | `[]` |
| `extraEnv` |  | `[]` |
| `resources` |  | `{}` |
| `securityContext.runAsNonRoot` |  | `true` |
| `containerSecurityContext` | Container Security Context to be set on the controller component container ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{}` |
| `volumes` |  | `[]` |
| `volumeMounts` |  | `[]` |
| `podLabels` | Optional additional annotations to add to the controller Deployment deploymentAnnotations: {}  Optional additional annotations to add to the controller Pods podAnnotations: {} | `{}` |
| `nodeSelector` | Optional additional labels to add to the controller Service serviceLabels: {}  Optional DNS settings, useful if you have a public and private DNS zone for the same domain on Route 53. What follows is an example of ensuring cert-manager can access an ingress or DNS TXT records at all times. NOTE: This requires Kubernetes 1.10 or `CustomPodDNS` feature gate enabled for the cluster to work. podDnsPolicy: "None" podDnsConfig:   nameservers:     - "1.1.1.1"     - "8.8.8.8" | `{}` |
| `ingressShim` |  | `{}` |
| `prometheus.enabled` |  | `true` |
| `prometheus.servicemonitor.enabled` |  | `false` |
| `prometheus.servicemonitor.prometheusInstance` |  | `default` |
| `prometheus.servicemonitor.targetPort` |  | `9402` |
| `prometheus.servicemonitor.path` |  | `/metrics` |
| `prometheus.servicemonitor.interval` |  | `60s` |
| `prometheus.servicemonitor.scrapeTimeout` |  | `30s` |
| `prometheus.servicemonitor.labels` |  | `{}` |
| `affinity` | Use these variables to configure the HTTP_PROXY environment variables http_proxy: "http://proxy:8080" https_proxy: "https://proxy:8080" no_proxy: 127.0.0.1,localhost  expects input structure as per specification https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/#affinity-v1-core for example:   affinity:     nodeAffinity:      requiredDuringSchedulingIgnoredDuringExecution:        nodeSelectorTerms:        - matchExpressions:          - key: foo.bar.com/role            operator: In            values:            - master | `{}` |
| `tolerations` | expects input structure as per specification https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/#toleration-v1-core for example:   tolerations:   - key: foo.bar.com/role     operator: Equal     value: master     effect: NoSchedule | `[]` |
| `webhook.replicaCount` |  | `1` |
| `webhook.timeoutSeconds` |  | `10` |
| `webhook.strategy` |  | `{}` |
| `webhook.securityContext.runAsNonRoot` |  | `true` |
| `webhook.containerSecurityContext` | Container Security Context to be set on the webhook component container ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{}` |
| `webhook.extraArgs` | Optional additional annotations to add to the webhook Deployment deploymentAnnotations: {}  Optional additional annotations to add to the webhook Pods podAnnotations: {}  Optional additional annotations to add to the webhook MutatingWebhookConfiguration mutatingWebhookConfigurationAnnotations: {}  Optional additional annotations to add to the webhook ValidatingWebhookConfiguration validatingWebhookConfigurationAnnotations: {}  Optional additional arguments for webhook | `[]` |
| `webhook.resources` |  | `{}` |
| `webhook.livenessProbe.failureThreshold` |  | `3` |
| `webhook.livenessProbe.initialDelaySeconds` |  | `60` |
| `webhook.livenessProbe.periodSeconds` |  | `10` |
| `webhook.livenessProbe.successThreshold` |  | `1` |
| `webhook.livenessProbe.timeoutSeconds` |  | `1` |
| `webhook.readinessProbe.failureThreshold` |  | `3` |
| `webhook.readinessProbe.initialDelaySeconds` |  | `5` |
| `webhook.readinessProbe.periodSeconds` |  | `5` |
| `webhook.readinessProbe.successThreshold` |  | `1` |
| `webhook.readinessProbe.timeoutSeconds` |  | `1` |
| `webhook.nodeSelector` |  | `{}` |
| `webhook.affinity` |  | `{}` |
| `webhook.tolerations` |  | `[]` |
| `webhook.podLabels` | Optional additional labels to add to the Webhook Pods | `{}` |
| `webhook.image.repository` |  | `quay.io/jetstack/cert-manager-webhook` |
| `webhook.image.pullPolicy` | Override the image tag to deploy by setting this variable. If no value is set, the chart's appVersion will be used. tag: canary  Setting a digest will override any tag digest: sha256:0e072dddd1f7f8fc8909a2ca6f65e76c5f0d2fcfb8be47935ae3457e8bbceb20 | `IfNotPresent` |
| `webhook.serviceAccount.create` | Specifies whether a service account should be created | `true` |
| `webhook.serviceAccount.automountServiceAccountToken` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template name: "" Optional additional annotations to add to the controller's ServiceAccount annotations: {} Automount API credentials for a Service Account. | `true` |
| `webhook.securePort` | The port that the webhook should listen on for requests. In GKE private clusters, by default kubernetes apiservers are allowed to talk to the cluster nodes only on 443 and 10250. so configuring securePort: 10250, will work out of the box without needing to add firewall rules or requiring NET_BIND_SERVICE capabilities to bind port numbers <1000 | `10250` |
| `webhook.hostNetwork` | Specifies if the webhook should be started in hostNetwork mode.  Required for use in some managed kubernetes clusters (such as AWS EKS) with custom CNI (such as calico), because control-plane managed by AWS cannot communicate with pods' IP CIDR and admission webhooks are not working  Since the default port for the webhook conflicts with kubelet on the host network, `webhook.securePort` should be changed to an available port if running in hostNetwork mode. | `false` |
| `webhook.serviceType` | Specifies how the service should be handled. Useful if you want to expose the webhook to outside of the cluster. In some cases, the control plane cannot reach internal services. | `ClusterIP` |
| `webhook.url` | Overrides the mutating webhook and validating webhook so they reach the webhook service using the `url` field instead of a service. | `{}` |
| `cainjector.enabled` |  | `true` |
| `cainjector.replicaCount` |  | `1` |
| `cainjector.strategy` |  | `{}` |
| `cainjector.securityContext.runAsNonRoot` |  | `true` |
| `cainjector.containerSecurityContext` | Container Security Context to be set on the cainjector component container ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{}` |
| `cainjector.extraArgs` | Optional additional annotations to add to the cainjector Deployment deploymentAnnotations: {}  Optional additional annotations to add to the cainjector Pods podAnnotations: {}  Optional additional arguments for cainjector | `[]` |
| `cainjector.resources` |  | `{}` |
| `cainjector.nodeSelector` |  | `{}` |
| `cainjector.affinity` |  | `{}` |
| `cainjector.tolerations` |  | `[]` |
| `cainjector.podLabels` | Optional additional labels to add to the CA Injector Pods | `{}` |
| `cainjector.image.repository` |  | `quay.io/jetstack/cert-manager-cainjector` |
| `cainjector.image.pullPolicy` | Override the image tag to deploy by setting this variable. If no value is set, the chart's appVersion will be used. tag: canary  Setting a digest will override any tag digest: sha256:0e072dddd1f7f8fc8909a2ca6f65e76c5f0d2fcfb8be47935ae3457e8bbceb20 | `IfNotPresent` |
| `cainjector.serviceAccount.create` | Specifies whether a service account should be created | `true` |
| `cainjector.serviceAccount.automountServiceAccountToken` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template name: "" Optional additional annotations to add to the controller's ServiceAccount annotations: {} Automount API credentials for a Service Account. | `true` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. For example,

```console
$ helm install my-cert-manager jetstack/cert-manager -n cert-manager --version=v0.0.1 --values values.yaml
```
> **Tip**: You can use the default [values.yaml](https://github.com/jetstack/cert-manager/blob/master/deploy/charts/cert-manager/values.yaml)

## Contributing

This chart is maintained at [github.com/jetstack/cert-manager](https://github.com/jetstack/cert-manager/tree/master/deploy/charts/cert-manager).
