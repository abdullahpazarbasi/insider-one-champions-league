#!/usr/bin/env bash

set -euo pipefail

if [ -z "${ROOT_DIR:-}" ]; then
    echo "🛑  ROOT_DIR is undefined" >&2
    exit 1
fi

docker build \
    -t "ioclservice:latest" \
    -f "${ROOT_DIR}/back/service/Dockerfile" \
    "${ROOT_DIR}/back/service"

docker build \
    -t "ioclbff:latest" \
    -f "${ROOT_DIR}/back/bff/Dockerfile" \
    "${ROOT_DIR}/back/bff"

docker build \
    -t "ioclspa:latest" \
    -f "${ROOT_DIR}/front/spa/Dockerfile" \
    "${ROOT_DIR}/front/spa"

docker tag "ioclservice:latest" "localhost:5000/ioclservice:latest"
docker tag "ioclservice:latest" "ghcr.io/abdullahpazarbasi/ioclservice:latest"

docker push "localhost:5000/ioclservice:latest"

docker tag "ioclbff:latest" "localhost:5000/ioclbff:latest"
docker tag "ioclbff:latest" "ghcr.io/abdullahpazarbasi/ioclbff:latest"

docker push "localhost:5000/ioclbff:latest"

docker tag "ioclspa:latest" "localhost:5000/ioclspa:latest"
docker tag "ioclspa:latest" "ghcr.io/abdullahpazarbasi/ioclspa:latest"

docker push "localhost:5000/ioclspa:latest"

if [[ -n "${GHCR_TOKEN:-}" ]]; then
    echo "${GHCR_TOKEN}" | docker login ghcr.io -u "abdullahpazarbasi" --password-stdin
    docker push "ghcr.io/abdullahpazarbasi/ioclservice:latest"
    docker push "ghcr.io/abdullahpazarbasi/ioclbff:latest"
    docker push "ghcr.io/abdullahpazarbasi/ioclspa:latest"
fi
