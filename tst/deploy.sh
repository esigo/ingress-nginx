#!/bin/bash
set -e

# docker build -t otel-dev .
# docker save --output otel-dev.tar otel-dev
# sshpass -p 'root' rsync otel-dev.tar -e 'ssh -p 2222' --progress root@127.0.0.1:/root/otel-dev.tar
# sshpass  -p 'root' ssh -o StrictHostKeychecking=no -p 2222 root@127.0.0.1 "sudo k3s ctr images import /root/otel-dev.tar"

cd ..

make build
make image

docker save --output controller_v1.3.0.tar gcr.io/k8s-staging-ingress-nginx/controller:v1.3.0
sshpass -p 'root' rsync controller_v1.3.0.tar -e 'ssh -p 2222' --progress root@127.0.0.1:/root/controller_v1.3.0.tar
sshpass  -p 'root' ssh -o StrictHostKeychecking=no -p 2222 root@127.0.0.1 "sudo k3s ctr images import /root/controller_v1.3.0.tar"

helm uninstall ingress-nginx --namespace ingress-nginx
helm upgrade --install ingress-nginx \
  /workspace/ingress-nginx/charts/ingress-nginx \
  --values /workspace/ingress-nginx/charts/ingress-nginx/values.yaml \
  --namespace ingress-nginx --create-namespace

# sshpass  -p 'root' ssh -o StrictHostKeychecking=no -p 2222 root@127.0.0.1 "sudo k3s ctr images ls"

# kubectl apply -f myapp.yaml
