# Monitoring

Prometheus monitoring and alerting system is used to provide useful insight into the
performance and possible issues of the api service API.

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

/metrics endpoint of the api service

### Prometheus dashboard
kubectl port-forward -n monitoring prometheus-kube-prometheus-0 9090

### Grafana dashboard
kubectl port-forward $(kubectl get  pods --selector=app=kube-prometheus-grafana -n  monitoring --output=jsonpath="{.items..metadata.name}") -n monitoring  3000

## Useful links
https://github.com/zbindenren/negroni-prometheus/blob/master/middleware.go

https://itnext.io/kubernetes-monitoring-with-prometheus-in-15-minutes-8e54d1de2e13

https://rsmitty.github.io/Prometheus-Exporters/

https://alex.dzyoba.com/blog/go-prometheus-service/

https://blog.alexellis.io/prometheus-monitoring/

https://sysdig.com/blog/prometheus-metrics/

https://github.com/prometheus/prometheus/issues/4052

http://www.ru-rocker.com/2017/04/02/micro-services-using-go-kit-monitoring-services

https://coreos.com/operators/prometheus/docs/latest/user-guides/running-exporters.html