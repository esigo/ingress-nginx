# FROM alpine:latest
FROM golang:latest

ARG KIND_VERSION=0.16.3

RUN apt install -y curl make git openssl

RUN curl -fsSL https://get.docker.com -o get-docker.sh \
    && bash ./get-docker.sh
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin/kubectl

RUN curl -L https://github.com/kubernetes-sigs/kind/releases/download/v${KIND_VERSION}/kind-linux-amd64 -o /usr/local/bin/kind && \
    chmod +x /usr/local/bin/kind

RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 \
    && chmod 700 get_helm.sh \
    && ./get_helm.sh

RUN curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)" \
    && chmod +x kubebuilder \
    && mv kubebuilder /usr/local/bin/

# helm repo add jetstack https://charts.jetstack.io
# helm install cert --create-namespace --namespace cert-manager --version v1.12.3 jetstack/cert-manager
# helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/

# helm upgrade --install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --create-namespace --namespace kubernetes-dashboard
# kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-nginx-controller 8443:443
# kubectl create serviceaccount dashboard-admin -n kubernetes-dashboard
# kubectl create clusterrolebinding dashboard-admin-binding --clusterrole=cluster-admin --serviceaccount=kubernetes-dashboard:dashboard-admin

# make create-kind-cluster
# kubectl get secret admin-user -n kubernetes-dashboard -o jsonpath={".data.token"} | base64 -d

# kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v0.7.1/standard-install.yaml
# kubectl apply -f deploy/manifests/nginx-gateway.yaml

