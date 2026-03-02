ARG RESTIC_IMAGE=ghcr.io/restic/restic
ARG RESTIC_TAG=0.18.1

FROM ghcr.io/bdd/runitor:v1.4.1-alpine AS runitor

FROM ghcr.io/clevyr/kubedb:1.19.0 AS kubedb

FROM $RESTIC_IMAGE:$RESTIC_TAG AS restic

RUN apk add --no-cache sqlite

COPY --from=runitor /usr/local/bin/runitor /usr/bin/runitor
COPY --from=kubedb /usr/local/bin/kubedb /usr/bin/kubedb

ENV KUBECONFIG=/.kube/config
