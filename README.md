- [ArgoCD](#argocd)
  - [Passo 01 - Instalação do ArgoCD](#passo-01---instalação-do-argocd)
  - [Passo 02 - Conhecendo nossa aplicação em Go](#passo-02---conhecendo-nossa-aplicação-em-go)
  - [Passo 03 - Conhecendo nosso repositório do GitOps](#passo-03---conhecendo-nosso-repositório-do-gitops)
  - [Passo 04 - Configurando nossa primeira App no ArgoCD](#passo-04---configurando-nossa-primeira-app-no-argocd)
    - [Configurando uma nova App](#configurando-uma-nova-app)
  - [Passo 05 - Simular alterações com o `auto sync` desabilitado.](#passo-05---simular-alterações-com-o-auto-sync-desabilitado)
  - [Passo 06 - Habilitando `auto sync`](#passo-06---habilitando-auto-sync)


#### Documentação do repositório de instalação
```
https://artifacthub.io/packages/helm/argo/argo-cd
```

## Passo 01 - Instalação do ArgoCD

- Documentação do repositório de instalação
```
https://artifacthub.io/packages/helm/argo/argo-cd
```

- Add repo via helm argocd
```
helm repo add argo https://argoproj.github.io/argo-helm
```

- Instalação via helm

```
helm install argo-cd argo/argo-cd --version 7.6.12 -f values-of.yaml --namespace argocd --create-namespace
```

- Caso não esteja utilizando ingress, expor a porta
```
kubectl port-forward service/argo-cd-argocd-server -n argocd 8080:443
```

- Recuperando a senha
```
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

> Considerando que você já tem o ingress funcionando no seu cluster

- Acesse via browser

http://dev.goapp.127.0.0.1.nip.io

> Usuário padrão é `admin`

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
