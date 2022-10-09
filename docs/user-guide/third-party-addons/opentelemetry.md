# OpenTelemetry

Enables requests served by NGINX for distributed telemetry via The OpenTelemetry Project.

Using the third party module [opentelemetry-cpp-contrib/nginx](https://github.com/open-telemetry/opentelemetry-cpp-contrib/tree/main/instrumentation/nginx) the NGINX ingress controller can configure NGINX to enable [OpenTelemetry](http://opentelemetry.io) instrumentation.
By default this feature is disabled.

## Usage

To enable the instrumentation we must enable OpenTelemetry in the configuration ConfigMap:
```yaml
data:
  enable-opentelemetry: "true"
```

To enable or disable instrumentation for a single Ingress, use
the `enable-opentelemetry` annotation:
```yaml
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/enable-opentelemetry: "true"
```

We must also set the host to use when uploading traces:

```yaml
otlp-collector-host: "otel-coll-collector.otel.svc"
```
NOTE: While the option is called `otlp-collector-host`, you will need to point this to any backend that recieves otlp-grpc.

Next you will need to deploy a distributed telemetry system which uses OpenTelemetry.
[opentelemetry-collector](https://github.com/open-telemetry/opentelemetry-collector) and
[Tempo](https://github.com/grafana/tempo)
have been tested.

Other optional configuration options:
```yaml
# specifies the name to use for the server span
opentelemetry-operation-name

# sets whether or not to trust incoming telemetry spans
opentelemetry-trust-incoming-span

# specifies the port to use when uploading traces, Default: 9411
otlp-collector-port

# specifies the service name to use for any traces created, Default: nginx
otel-service-name

# The maximum queue size. After the size is reached data are dropped.
otel-max-queuesize

# The delay interval in milliseconds between two consecutive exports.
otel-schedule-delay-millis
        
# How long the export can run before it is cancelled.
otel-schedule-delay-millis

# The maximum batch size of every export. It must be smaller or equal to maxQueueSize.
otel-max-export-batch-size

# specifies sample rate for any traces created, Default: 0.01
otel-sampler-ratio

# specifies the port to use when uploading traces, Default: 4317
otlp-collector-port

# specifies the sampler to be used when sampling traces.
# The available samplers are: AlwaysOn,  AlwaysOff, TraceIdRatioBased, Default: AlwaysOff
otel-sampler

# Uses sampler implementation which by default will take a sample if parent Activity is sampled, Default: false
otel-sampler-parent-based
```

All these options (including host) allow environment variables, such as `$HOSTNAME` or `$HOST_IP`. In the case of Jaeger, if you have a Jaeger agent running on each machine in your cluster, you can use something like `$HOST_IP` (which can be 'mounted' with the `status.hostIP` fieldpath, as described [here](https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/#capabilities-of-the-downward-api)) to make sure traces will be sent to the local agent.


Note that you can also set whether to trust incoming spans (global default is true) per-location using annotations like the following:
```
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/opentelemetry-trust-incoming-span: "true"
```

## Examples

The following examples show how to deploy and test different distributed telemetry systems. These example can be performed using Minikube.

### Zipkin

In the [rnburn/zipkin-date-server](https://github.com/rnburn/zipkin-date-server)
GitHub repository is an example of a dockerized date service. To install the example and Zipkin collector run:

```
kubectl create -f https://raw.githubusercontent.com/rnburn/zipkin-date-server/master/kubernetes/zipkin.yaml
kubectl create -f https://raw.githubusercontent.com/rnburn/zipkin-date-server/master/kubernetes/deployment.yaml
```

Also we need to configure the Ingress-NGINX controller ConfigMap with the required values:

```
$ echo '
apiVersion: v1
kind: ConfigMap
data:
  enable-opentelemetry: "true"
  zipkin-collector-host: zipkin.default.svc.cluster.local
metadata:
  name: ingress-nginx-controller
  namespace: kube-system
' | kubectl replace -f -
```

In the Zipkin interface we can see the details:
![zipkin screenshot](../../images/zipkin-demo.png "zipkin collector screenshot")

### Jaeger

1. Enable Ingress addon in Minikube:
    ```
    $ minikube addons enable ingress
    ```

2. Add Minikube IP to /etc/hosts:
    ```
    $ echo "$(minikube ip) example.com" | sudo tee -a /etc/hosts
    ```

3. Apply a basic Service and Ingress Resource:
    ```
    # Create Echoheaders Deployment
    $ kubectl run echoheaders --image=registry.k8s.io/echoserver:1.4 --replicas=1 --port=8080

    # Expose as a Cluster-IP
    $ kubectl expose deployment echoheaders --port=80 --target-port=8080 --name=echoheaders-x

    # Apply the Ingress Resource
    $ echo '
      apiVersion: networking.k8s.io/v1
      kind: Ingress
      metadata:
        name: echo-ingress
      spec:
        ingressClassName: nginx
        rules:
        - host: example.com
          http:
            paths:
            - path: /echo
              pathType: Prefix
              backend:
                service:
                  name: echoheaders-x
                  port:
                    number: 80
      ' | kubectl apply -f -
    ```

4. Enable OpenTelemetry and set the jaeger-collector-host:
    ```
    $ echo '
      apiVersion: v1
      kind: ConfigMap
      data:
        enable-opentelemetry: "true"
        jaeger-collector-host: jaeger-agent.default.svc.cluster.local
      metadata:
        name: ingress-nginx-controller
        namespace: kube-system
      ' | kubectl replace -f -
    ```

5. Apply the Jaeger All-In-One Template:
    ```
    $ kubectl apply -f https://raw.githubusercontent.com/jaegertelemetry/jaeger-kubernetes/master/all-in-one/jaeger-all-in-one-template.yml
    ```

6. Make a few requests to the Service:
    ```
    $ curl example.com/echo -d "meow"

    CLIENT VALUES:
    client_address=172.17.0.5
    command=POST
    real path=/echo
    query=nil
    request_version=1.1
    request_uri=http://example.com:8080/echo

    SERVER VALUES:
    server_version=nginx: 1.10.0 - lua: 10001

    HEADERS RECEIVED:
    accept=*/*
    connection=close
    content-length=4
    content-type=application/x-www-form-urlencoded
    host=example.com
    user-agent=curl/7.54.0
    x-forwarded-for=192.168.99.1
    x-forwarded-host=example.com
    x-forwarded-port=80
    x-forwarded-proto=http
    x-original-uri=/echo
    x-real-ip=192.168.99.1
    x-scheme=http
    BODY:
    meow
    ```

7. View the Jaeger UI:
    ```
    $ minikube service jaeger-query --url

    http://192.168.99.100:30183
    ```

    In the Jaeger interface we can see the details:
    ![jaeger screenshot](../../images/jaeger-demo.png "jaeger collector screenshot")
