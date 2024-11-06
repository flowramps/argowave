
#### Documentação do repositório de instalação
```
https://artifacthub.io/packages/helm/argo/argo-cd
```
#### Instalação e acesso

Add repo via helm argocd
```
helm repo add argo https://argoproj.github.io/argo-helm
```

Instalação via helm
```
helm install argo-cd argo/argo-cd --version 7.6.12 -f values-of.yaml
```

Caso não esteja usando ingress expor a porta
```
kubectl port-forward service/argo-cd-argocd-server -n argocd 8080:443
```

Recuperar secret
```
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

#### ArgoCD

Login via cli expondo a porta
```
argocd login localhost:8080 --username admin --password changeme
```

Adicionar repositório git no argocd via cli
```
argocd repo add https://github.com/flowramps/argowave.git --name "argowave" --project testeds --username nonexistant --password changeme --insecure-skip-server-verification --grpc-web --upsert
```

```
argocd app get flow-ramps
```
```
argocd repo list
```
Validar projetos 
```
kubectl get appprojects -n argocd 
```
```
kubectl get app -n argocd
```

Teste fake label
```
kubectl patch configmap argocd-cm -n argocd --type='merge' -p '{"data": {"globalProjects": "- labelSelector:\n    matchExpressions:\n    - key: fake.label\n      operator: DoesNotExist\n  projectName: default"}}'
```
```
kubectl rollout restart deployment argo-cd-argocd-server -n argocd
```

#### Go App

Instale as dependências
```
go mod download

```

Compile e execute a aplicação Go
```
go run main.go
```