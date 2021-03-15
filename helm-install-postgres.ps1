helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm upgrade --install postgres bitnami/postgresql -n gomicro --create-namespace -f .\postgres-values.yaml