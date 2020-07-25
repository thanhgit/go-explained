# Helm

## Fix error: configmaps is forbidden: User "system:serviceaccount:kube-system:default" cannot list resource "configmaps"
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'      
helm init --service-account tiller --upgrade

## Set up helm to kubernetes
```text
export KUBECONFIG='/root/.kube/config'
# disable plugin with
export HELM_NO_PLUGINS=1
```
## Tab completion
```text
$source <(helm completion $SHELL)
```
## Common command
```text
$helm version
$helm init --upgrade
$helm ls
$helm search wordpress
```

## Interaction charts
* ### Install
```text
$helm install stable/wordpress --name wp-v1
helm install stable/vault-operator --name vault-server --set serviceType=LoadBalancer
helm install redis bitnami/redis --namespace= redis
```
* ### Delete
```text
$helm delete wp-v1
$helm delete wp-v1 --purge
```
* ### Upgrade
```text
$helm upgrade wp-v1 --version 9.0.3
```
* ### Rollback
```text
$helm rollback 
```

## Plugin
* ### Install
```text
$helm plugin install https://github.com/databus23/helm-diff
```
* ### Others
| Plugin | Description |
| -- | -- |
| helm diff | diff between a deployed release and a proposed helm upgrade |
| helm secrets | Used to help conceal secrets from Helm charts |
| helm monitor | Used to monitor a release and perform a rollback if certain events occur |
| helm unittest | Used to perform unit testing on a Helm chart |