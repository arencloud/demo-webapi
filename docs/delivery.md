# Delivery and Deployment

## CI/CD pipeline

Tekton pipeline definition:

- `.tekton/pipeline.yaml`

Pipeline responsibilities:

- clone source
- run unit tests
- build and push container image
- promote by updating the target GitOps overlay image tag

## GitOps and Argo CD

GitOps resources:

- `gitops/base/*`
- `gitops/overlays/dev|stage|prod/*`
- `gitops/argocd/applicationset.yaml`

Argo CD syncs environment overlays continuously.

## Connectivity Link exposure

Per-environment Connectivity Link resources are managed in:

- `gitops/overlays/*/connectivity-link.yaml`
