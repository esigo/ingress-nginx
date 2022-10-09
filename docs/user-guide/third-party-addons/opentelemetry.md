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
