# 使用するコマンド
## 最初から立ち上げる
``` sh
# cluster作成
kind create cluster --name <name>
# goのビルド
docker build -o main
# imageをkindに知らせる
kind load docker-image go-postgres-app:0.0.1 --name my-cluster
# yamlをkubectlで適用
kubectl apply -f go-app-deployment.yaml
kubectl apply -f postgres-deployment.yaml
# pod確認
kubectl get pods
# node確認
kubectl get nodes
# pod詳細確認
kubectl describe pods <podのID>
```


## 詰まったポイント
- latesタグにするとkubernetesが自動でリモーとからプルしようとするらしい
  - 対策として、latest以外のタグにする＆yamlに`imagePullPolicy: Never`を入れる
`

