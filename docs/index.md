# Demo WebAPI

`demo-webapi` is a Go HTTP service exposing ping and auth-related endpoints.

## Runtime

- Language: Go
- Default port: `8080`
- Container image: `quay.io/rh-ee-egevorky/demo-webapi`

## Endpoints

- `GET /api/ping`
- `GET /realms/ext/npwl-user/logout`
- `GET /realms/ext/protocol/cas/connect`

## RHDH integration

This component is registered with `catalog-info.yaml` and exposes:

- Component docs through TechDocs
- API entity docs from `docs/swagger.yaml`
- CI/CD and deployment visibility via Tekton and Argo CD annotations
