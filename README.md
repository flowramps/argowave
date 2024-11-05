https://artifacthub.io/packages/helm/argo/argo-cd

helm repo add argo https://argoproj.github.io/argo-helm

helm install argo-cd argo/argo-cd --version 7.6.12 -f values-of.yaml

kubectl port-forward service/argo-cd-argocd-server -n argocd 8080:443
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d



argocd repo add https://github.com/flowramps/workshop-argo.git --name "flow-ramps"  --username nonexistant --password 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 --grpc-web --upsert
argocd repo add https://github.com/flowramps/workshop-argo.git --name "flow-ramps"  --username nonexistant --password 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 --grpc-web --upsert


argocd repo add https://github.com/flowramps/workshop-argo.git --name "flow-ramps" --project flow-ramps --username nonexistant --password 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 --insecure-skip-server-verification --grpc-web --upsert

argocd login localhost:8080 --username admin --password jE9c8uRvaPxdkCQn

argocd app get flow-ramps

argocd repo list

Validar projetos 
kubectl get appprojects -n argocd 
kubectl get app -n argocd


kubectl patch configmap argocd-cm -n argocd --type='merge' -p '{"data": {"globalProjects": "- labelSelector:\n    matchExpressions:\n    - key: fake.label\n      operator: DoesNotExist\n  projectName: default"}}'

kubectl rollout restart deployment argo-cd-argocd-server -n argocd