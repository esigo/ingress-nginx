helm repo add jetstack https://charts.jetstack.io
helm install cert --create-namespace --namespace cert-manager --version v1.12.3 jetstack/cert-manager

helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm upgrade --install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --create-namespace --namespace kubernetes-dashboard
kubectl apply -f dashboard.yml
kubectl get secret admin-user -n kubernetes-dashboard -o jsonpath={".data.token"} | base64 -d

kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-nginx-controller 8443:443
# make create-kind-cluster
# https://localhost:8443/#/pod/gateway-system/
