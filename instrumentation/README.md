# Instrumentation

The application, a set of microservices, is running in Kubernetes.
Prometheus monitoring and alerting system is used to provide useful insight into the
performance and possible issues.

## Prometheus installation

### Installing prometheus operator through a helm chart

  - install Helm (https://github.com/kubernetes/helm/blob/master/docs/install.md):
    - curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh
    - chmod 700 get_helm.sh
    - ./get_helm.sh

  - install Kubernetes Prometheus operator
    - helm repo add coreos https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/
    - helm install coreos/prometheus-operator --name prometheus-operator --namespace monitoring
    - helm install coreos/kube-prometheus --name kube-prometheus --set global.rbacEnable=true --namespace monitoring

### Prometheus configuration

Configuration which services to scrape is done through Kubernetes Custom Resources. Prometheus operator CRs are of Kind ServiceMonitor and are stored in deployment/prometheus directory:

kubectl apply -g deployment/prometheus/

## Display

### Prometheus dashboard
kubectl port-forward -n monitoring prometheus-kube-prometheus-0 9090

### Grafana dashboard
kubectl port-forward $(kubectl get  pods --selector=app=kube-prometheus-grafana -n  monitoring --output=jsonpath="{.items..metadata.name}") -n monitoring  3000