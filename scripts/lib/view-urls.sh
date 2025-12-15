#!/usr/bin/env bash

set -euo pipefail

if [ -z "${PROFILE_NAME:-}" ]; then
    echo "🛑  PROFILE_NAME is undefined" >&2
    exit 1
fi

SPA_URL="$(minikube -p "${PROFILE_NAME}" service spa --url | head -n1 || true)"
echo "📌  SPA URL: https://${SPA_HOST}/"
echo "📌  SPA node URL: ${SPA_URL}"

BFF_URL="$(minikube -p "${PROFILE_NAME}" service bff --url | head -n1 || true)"
echo "📌  BFF URL: https://${BFF_HOST}/"
echo "📌  BFF node URL: ${BFF_URL}"
