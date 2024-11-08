# How to deploy manifest


### Monitoring metrics menggunakan prometheus dan grafana dashboard

`~$ kubectl create namespace monitoring`

`~$ helm upgrade --install prom prometheus-community/kube-prometheus-stack --namespace monitoring --values prom-values.yaml`


### Logging pada pods menggunakan Loki 

`~$ helm upgrade --install promtail grafana/promtail -f promtail-values.yaml -n monitoring`

`~$ helm upgrade --install loki grafana/loki-distributed -n monitoring`


### Ingress grafana

* How to expose grafana

`~$ kubectl apply -f ingress-grafana.yaml`