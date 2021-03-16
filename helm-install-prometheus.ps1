helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade --install prometheus prometheus-community/kube-prometheus-stack -n gomicro --create-namespace -f .\prometheus-values.yaml
kubectl get secret prometheus-grafana -o jsonpath="{.data.admin-user}" | gobase64.exe -d
kubectl get secret prometheus-grafana -o jsonpath="{.data.admin-password}" | gobase64.exe -d