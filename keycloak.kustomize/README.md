# keycloak kustomizations
keycloak is an authentication server by jboss.

To install / remove:
```sh
kustomize build overlays/prod | kubectl apply -f -

kustomize build overlays/prod | kubectl delete -f -
```
Subtitute prod for dev to deploy a dev server.

*IMPORTANT*: Do not deploy both dev and prod.
This causes a conflict as the docker images have
built-in clustering and clustering has not (yet)
been properly configured.
