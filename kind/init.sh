go mod init k8s.io/ingress-nginx-gateway
kubebuilder init --domain k8s.io --repo k8s.io/ingress-nginx/gateway
kubectl api-resources | grep gate
kubebuilder create api --controller --group gateway.networking --version v1beta1 --kind GatewayClass
kubebuilder create api --controller --group gateway.networking --version v1beta1 --kind Gateway
kubebuilder create api --controller --group gateway.networking --version v1beta1 --kind HTTPRoute

# sigs.k8s.io/gateway-api/apis/v1beta1