# Red Hat Developer Hub 1.9 End-to-End Scaffold

This repo now includes a baseline RHDH 1.9 delivery flow:

- `catalog-info.yaml`: Backstage entities for service + API.
- `mkdocs.yml` + `docs/*.md`: TechDocs pages rendered in RHDH.
- `.tekton/pipeline.yaml`: two pipelines:
  - `demo-webapi-dev`: build, test, push, deploy to `dev`
  - `demo-webapi-promote`: controlled promotion `dev->stage` and `stage->prod`
- `gitops/overlays/*`: Kustomize overlays for `dev`, `stage`, `prod`.
- `gitops/argocd/applicationset.yaml`: Argo CD ApplicationSet for all environments.
- `gitops/overlays/*/connectivity-link.yaml`: Red Hat Connectivity Link exposure for each environment.

## Prerequisites

1. OpenShift Pipelines installed (Tekton), with namespace `Task` definitions for `git-clone` and `buildah`.
2. OpenShift GitOps (Argo CD) installed.
3. Red Hat Connectivity Link Operator installed (Gateway API + Kuadrant CRDs available).
4. A valid `GatewayClass` available (default scaffold uses `istio`).
5. A `ClusterIssuer` for TLS (default scaffold uses `letsencrypt-prod`).
6. DNS provider credentials secret (default scaffold uses `aws-credentials`).
7. Registry credentials configured for the pipeline service account.
8. Git credentials secret named `git-credentials` for push access.

## Configure placeholders

Adjust repository/image placeholders in:

- `.tekton/pipelinerun-dev.yaml`
- `gitops/base/kustomization.yaml`
- `gitops/overlays/*/kustomization.yaml`
- `gitops/argocd/applicationset.yaml`

Replace environment hostnames and infrastructure-specific values in:

- `gitops/overlays/dev/connectivity-link.yaml`
- `gitops/overlays/stage/connectivity-link.yaml`
- `gitops/overlays/prod/connectivity-link.yaml`

Specifically check:

- `spec.gatewayClassName`
- Gateway listener `hostname`
- `TLSPolicy.spec.issuerRef.name`
- `DNSPolicy.spec.providerRefs[].name`

## Run CI/CD pipeline

Apply the pipeline and run example:

```bash
oc apply -n builder -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/git-clone/0.9/git-clone.yaml
oc apply -n builder -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/buildah/0.9/buildah.yaml
oc apply -f .tekton/pipeline.yaml
oc create -f .tekton/pipelinerun-dev.yaml
```

For stage promotion, apply:

```bash
oc create -f .tekton/pipelinerun-stage.yaml
```

For prod promotion, apply:

```bash
oc create -f .tekton/pipelinerun-prod.yaml
```

Use the same immutable `imageTag` through dev, stage, and prod.

## Enable Argo CD deployment

Apply ApplicationSet:

```bash
oc apply -f gitops/argocd/applicationset.yaml
```

Argo CD will continuously sync:

- `gitops/overlays/dev`
- `gitops/overlays/stage`
- `gitops/overlays/prod`

## Register in Red Hat Developer Hub 1.9

1. In RHDH, open **Create** > **Register Existing Component**.
2. Use the repo URL containing `catalog-info.yaml`.
3. Confirm the component appears in catalog with Tekton and Argo CD plugin cards.
4. Open the `Docs` tab on `demo-webapi` to verify TechDocs rendering.
5. Open `demo-webapi-api` entity and confirm OpenAPI is rendered from `docs/swagger.yaml`.
