#!/bin/bash
set -e

pushd A
docker build -t service-a:last .
popd

pushd B
docker build -t service-b:last .
popd

docker save --output service-a.tar service-a:last
docker save --output service-b.tar service-b:last

#sshpass -p 'root' scp --progress -o StrictHostKeychecking=no -P 2222 root@127.0.0.1:/etc/rancher/k3s/k3s.yaml ~/.kube/config

sshpass -p 'root' rsync service-a.tar -e 'ssh -p 2222' --progress root@127.0.0.1:/root/service-a.tar
sshpass -p 'root' rsync service-b.tar -e 'ssh -p 2222' --progress root@127.0.0.1:/root/service-b.tar

sshpass  -p 'root' ssh -o StrictHostKeychecking=no -p 2222 root@127.0.0.1 "sudo k3s ctr images import /root/service-a.tar"
sshpass  -p 'root' ssh -o StrictHostKeychecking=no -p 2222 root@127.0.0.1 "sudo k3s ctr images import /root/service-b.tar"

sshpass  -p 'root' ssh -o StrictHostKeychecking=no -p 2222 root@127.0.0.1 "sudo k3s ctr images ls"

kubectl delete namespace myapp
kubectl create namespace myapp
kubectl apply -f myapp.yaml