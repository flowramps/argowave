# ArgoCD
![image](https://github.com/user-attachments/assets/0473d207-c76d-4a25-991b-ff81f0b37a95)


### Objetivo

Este repositório foi criado para demonstrar a instalação e configuração do ArgoCD em um ambiente Kubernetes, junto com uma aplicação em Go. Ele é organizado em passos detalhados, cobrindo desde a instalação do ArgoCD usando Helm até a configuração de sincronização automática e controle de versões no GitOps.

1. **Instalação do ArgoCD:** Guiar na instalação e configuração básica do ArgoCD no cluster.
2. **Aplicação Go:** Introduzir a aplicação em Go, fornecendo instruções para execução local e integração com o ArgoCD.
3. **GitOps e Primeira Aplicação:** Configurar o ArgoCD para gerenciar a aplicação via repositório GitOps.
4. **Sincronização e Monitoramento:** Demonstrar o uso do "auto sync" para facilitar o processo de atualização contínua e gerenciamento de mudanças.

Esse repositório serve como uma referência para implementar práticas de GitOps, utilizando ArgoCD para automação e gestão de deploys em Kubernetes.



- [ArgoCD](#argocd)
  - [Passo 01 - Instalação do ArgoCD](#passo-01---instalação-do-argocd)
  - [Passo 02 - Conhecendo nossa aplicação em Go](#passo-02---conhecendo-nossa-aplicação-em-go)
  - [Passo 03 - Conhecendo nosso repositório do GitOps](#passo-03---conhecendo-nosso-repositório-do-gitops)
  - [Passo 04 - Configurando nossa primeira App no ArgoCD](#passo-04---configurando-nossa-primeira-app-no-argocd)
    - [Configurando uma nova App](#configurando-uma-nova-app)
  - [Passo 05 - Simular alterações com o `auto sync` desabilitado.](#passo-05---simular-alterações-com-o-auto-sync-desabilitado)
  - [Passo 06 - Habilitando `auto sync`](#passo-06---habilitando-auto-sync)



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

#### ArgoCD via CLI

- Instuções de instalação doc oficial, necessário para intereção via terminal
```
https://kostis-argo-cd.readthedocs.io/en/refresh-docs/getting_started/install_cli/
```
- Login via cli expondo a porta
```
argocd login localhost:8080 --username admin --password changeme
```

- Exemplos possiveis para login, doc oficial
```
https://argo-cd.readthedocs.io/en/stable/user-guide/commands/argocd_login/
```

- Adicionar repositório git no argocd via cli
```
argocd repo add https://github.com/flowramps/argowave.git --name "argowave" --project testeds --username nonexistant --password changeme --insecure-skip-server-verification --grpc-web --upsert
```

Listar apps existentes
```
kubectl get app -n argocd
```
Observar informações de uma app especifica 
```
argocd app get flow-ramps
```

Listar repositórios 
```
argocd repo list
```

Listar projetos 
```
kubectl get appprojects -n argocd 
```


Anotações para testes, label fake
```
kubectl patch configmap argocd-cm -n argocd --type='merge' -p '{"data": {"globalProjects": "- labelSelector:\n    matchExpressions:\n    - key: fake.label\n      operator: DoesNotExist\n  projectName: default"}}'
```
```
kubectl rollout restart deployment argo-cd-argocd-server -n argocd
```

## Passo 02 - Conhecendo nossa aplicação em GO

- Repositório: https://github.com/flowramps/argowave/blob/main/main.go

Instale as dependências
```
go mod download
```

Compile e execute a aplicação Go localmente
```
go run main.go
```

## Passo 03 - Conhecendo nosso repositório do GitOps

- Repositório: https://github.com/flowramps/argowave/tree/main/.github/workflows

## Passo 04 - Configurando nossa primeira App no ArgoCD

### Configurando uma nova App

New App
![image](https://github.com/user-attachments/assets/da3b8c98-df4f-4749-8f80-134e5406e265)


General
![image](https://github.com/user-attachments/assets/7b6f4ad8-664f-48b7-b857-6edb4c32344b)


Source
![image](https://github.com/user-attachments/assets/d4cca8fc-bf48-4e06-ad0b-36d68d7338fe)


Destination
![image](https://github.com/user-attachments/assets/bb4f40a6-7453-4ada-89cc-377802956524)


Kustomize
![image](https://github.com/user-attachments/assets/1d0a8574-42fe-48cf-bdbc-b704e74f6083)

> Para o nosso exemplo acima vou utilizar o `Kustomize`, mas pode ser variado dependendo do fluxo do deploy que você estiver criando, podendo ser helm, kustomize, plugin ...

Clique em **Create**

Agora faça o sync, para que o deploy conclua com sucesso 
![image](https://github.com/user-attachments/assets/a8bfa874-1bb1-4d2c-ada3-3e8c55f9e5d8)



## Passo 05 - Simular alterações com o `auto sync` desabilitado.

- Aumente o número de replicas do deployment

```
kubectl -n flowramps scale deployment goapp --replicas 5
```
> OBS, se estiver utilizando `HPA` como no meu caso, edite o `HPA` para surtir efeito das alterações.



- Veja o argo OutOfSync
![image](https://github.com/user-attachments/assets/bd80aa2c-f349-4702-b361-7ae84d425b54)


- App Diff
![image](https://github.com/user-attachments/assets/0417ade8-adaa-4324-9495-08994e1461ae)


## Passo 06 - Habilitando `auto sync`

![image](https://github.com/user-attachments/assets/6ed45db3-948f-4876-9bac-d67675853230)


![image](https://github.com/user-attachments/assets/580df01d-a635-4a91-bfc1-abbf42ece7b1)


![image](https://github.com/user-attachments/assets/32046485-fefa-453e-be61-3a6dbf5b763f)

Role a barra até o final para localizar o `SYN POLICY` e habilite
![image](https://github.com/user-attachments/assets/e8494bf0-f08f-40b9-89e9-8fdf2d6b1365)

Ao clicar, irá pedir um ok dizendo, Tem certeza de que deseja habilitar a sincronização automatizada de aplicativos?
![image](https://github.com/user-attachments/assets/ca88d5e5-53d3-41e6-9fa4-1b193685e42e)


Feito isso, o Auto sync estará ativo e você pode esperar ele fazer o processo automatico ou forçar ele da seguinte forma caso esteja ansioso.
![image](https://github.com/user-attachments/assets/c62e08f3-c63e-4c4c-92f0-dd3431041153)










