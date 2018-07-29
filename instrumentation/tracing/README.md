# Tracing

Distributed tracing helps understand system's architecture, critical paths for a request across services, allows to pinpoint and understand latencies.

Instrumentation usually starts by adding tracing to common pieces in the code like the RPC framework or routing middleware. Within an application, two types of tracing can be distinguished, depending on their focus:

  - Internal tracing: Operations happening inside the service. The most common one is a request being served by the application.
  - Outbound tracing: Calls to external resources like RPC calls, communication with messaging applications or database operations.


## Zipkin
Zipkin is an open source distributed tracing tool created by twitter.

## Installation

### Zipkin server
a) running locally
docker run -d -p 9411:9411 openzipkin/zipkin

b) running within Kubernetes cluster

  ba) using k8s manifests

    https://github.com/fabric8io/kubernetes-zipkin

    baa)
    kubectl create -f http://repo1.maven.org/maven2/io/fabric8/zipkin/zipkin-starter/0.1.5/zipkin-starter-0.1.5-kubernetes.yml

    bab)
    kubectl create -f http://repo1.maven.org/maven2/io/fabric8/zipkin/zipkin-starter-minimal/0.1.5/zipkin-starter-minimal-0.1.5-kubernetes.yml

  bb) as a Helm Chart

    helm repo add zipkin-helm https://financial-times.github.io/zipkin-helm/docs

    helm install -f my-cassandra-config.yaml https://financial-times.github.io/zipkin-helm/docs/zipkin-helm-0.1.1.tgz

  my-cassandra-config.yaml ... for override of default values

### Zipkin libraries
github.com/openzipkin/zipkin-go/model

### Notes
Only the library was implemented. Currently not used within the microservices.

## Credits
https://medium.com/devthoughts/instrumenting-a-go-application-with-zipkin-b79cc858ac3e

## Useful links
https://github.com/alextanhongpin/go-jaeger-trace/blob/master/main.go

https://github.com/helm/charts/tree/master/incubator/jaeger

https://github.com/Financial-Times/zipkin-helm

https://github.com/fabric8io/kubernetes-zipkin#minikube

https://github.com/openzipkin/zipkin-go/blob/master/example_httpserver_test.go
