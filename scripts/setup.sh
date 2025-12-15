#!/usr/bin/env bash

# shellcheck source=/dev/null
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/lib/assert-running-in-bash.sh"

set -euo pipefail

cd "$(dirname "$0")/.."
ROOT_DIR="$(pwd)"

# shellcheck source=/dev/null
source "${ROOT_DIR}/scripts/lib/dotenv.sh"

if [ -z "${PROFILE_NAME:-}" ]; then
    echo "🛑  PROFILE_NAME is undefined" >&2
    exit 1
fi

echo ""
echo "--------------------------------------------------------------------------------"
echo " 🏗️  Setup"
echo "--------------------------------------------------------------------------------"

K8S_BASE_DIR="${ROOT_DIR}/k8s/base"
K8S_LOCAL_DIR="${ROOT_DIR}/k8s/overlays/local"

if [ ! -f "${ROOT_DIR}/.env.local" ]; then
    cp "${ROOT_DIR}/.env.local.dist" "${ROOT_DIR}/.env.local"
fi

set +e
bash "${ROOT_DIR}/scripts/create-image-registry.sh"
exit_code=$?
set -e
if [ $exit_code -eq 0 ]; then
    echo "✅  The image registry is ready"
else
    echo "🛑  The image registry could not be created (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
# shellcheck disable=SC2097,SC2098
ROOT_DIR="${ROOT_DIR}" bash "${ROOT_DIR}/scripts/lib/ensure-cluster-tls-certificates-exist.sh" "${BFF_HOST}"
exit_code=$?
set -e
if [ $exit_code -ne 0 ]; then
    echo "🛑  TLS certificates could not be created (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
# shellcheck disable=SC2097,SC2098
ROOT_DIR="${ROOT_DIR}" bash "${ROOT_DIR}/scripts/lib/ensure-cluster-tls-certificates-exist.sh" "${SPA_HOST}"
exit_code=$?
set -e
if [ $exit_code -ne 0 ]; then
    echo "🛑  TLS certificates could not be created (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
# shellcheck disable=SC2097,SC2098
ROOT_DIR="${ROOT_DIR}" bash "${ROOT_DIR}/scripts/lib/ensure-docker-exists.sh"
exit_code=$?
set -e
if [ $exit_code -ne 0 ]; then
    echo "🛑  Docker is not available (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
# shellcheck disable=SC2097,SC2098
ROOT_DIR="${ROOT_DIR}" bash "${ROOT_DIR}/scripts/lib/build-and-register-images.sh"
exit_code=$?
set -e
if [ $exit_code -ne 0 ]; then
    echo "🛑  image building and registering failed (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
# shellcheck disable=SC2097,SC2098
ROOT_DIR="${ROOT_DIR}" bash "${ROOT_DIR}/scripts/lib/ensure-minikube-exists.sh"
exit_code=$?
set -e
if [ $exit_code -ne 0 ]; then
    echo "🛑  Minikube is not available (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
bash "${ROOT_DIR}/scripts/lib/assert-minikube-host-running.sh"
exit_code=$?
set -e
if [ $exit_code -ne 0 ]; then
    echo "⏳  Minikube '${PROFILE_NAME}' is being started (exit code: ${exit_code})..."
    set +e
    bash "${ROOT_DIR}/scripts/lib/start-minikube-cluster.sh"
    exit_code=$?
    set -e
    if [ $exit_code -ne 0 ]; then
        echo "🛑  Minikube '${PROFILE_NAME}' could not be started (exit code: ${exit_code})" >&2
        exit 1
    fi
fi

set +e
# shellcheck disable=SC2097,SC2098
host_IP="$( ROOT_DIR="${ROOT_DIR}" bash "${ROOT_DIR}/scripts/lib/resolve-host-ip.sh" )"
exit_code=$?
set -e
if [ $exit_code -eq 0 ]; then
    echo "✅  Host IP resolved: ${host_IP}"
else
    echo "🛑  Host IP could not be resolved (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
bash "${ROOT_DIR}/scripts/add-image-registry-host-entry-in-minikube.sh" "${host_IP}"
exit_code=$?
set -e
if [ $exit_code -eq 0 ]; then
    echo "✅  The image registry hostname is registered"
else
    echo "🛑  The image registry hostname could not be registered (exit code: ${exit_code})" >&2
    exit 1
fi

set +e
# shellcheck disable=SC2097,SC2098
ROOT_DIR="${ROOT_DIR}" bash "${ROOT_DIR}/scripts/lib/trust-in-ca-for-image-registry.sh"
exit_code=$?
set -e
if [ $exit_code -eq 0 ]; then
    echo "✅  Minikube '${PROFILE_NAME}' now trusts the root CA"
else
    echo "🛑  Minikube '${PROFILE_NAME}' could not trust in root CA (exit code: ${exit_code})" >&2
    exit 1
fi

minikube -p "${PROFILE_NAME}" kubectl -- apply -k "${K8S_LOCAL_DIR}"

set +e
bash "${ROOT_DIR}/scripts/lib/wait-for-ingress-webhook-to-become-healthy.sh"
exit_code=$?
set -e
if [ $exit_code -eq 0 ]; then
    echo "⏳  Ingress manifests is being applied..."
    minikube -p "${PROFILE_NAME}" kubectl -- apply -f "${K8S_BASE_DIR}/ingress/bff-ingress.yaml"
    minikube -p "${PROFILE_NAME}" kubectl -- apply -f "${K8S_BASE_DIR}/ingress/spa-ingress.yaml"
else
    echo ""
    echo "⚠️   Apply ingress manifest manually by making run the command below:"
    echo ""
    echo "minikube -p \"${PROFILE_NAME}\" kubectl -- apply -f \"${K8S_BASE_DIR}/ingress/bff-ingress.yaml\""
    echo "minikube -p \"${PROFILE_NAME}\" kubectl -- apply -f \"${K8S_BASE_DIR}/ingress/spa-ingress.yaml\""
    echo ""
fi

minikube -p "${PROFILE_NAME}" kubectl -- rollout status deploy/service --timeout=300s
minikube -p "${PROFILE_NAME}" kubectl -- rollout status deploy/bff --timeout=300s
minikube -p "${PROFILE_NAME}" kubectl -- rollout status deploy/spa --timeout=300s

set +e
bash "${ROOT_DIR}/scripts/lib/wait-for-cluster-to-become-healthy.sh"
exit_code=$?
set -e
if [ $exit_code -eq 0 ]; then
    echo "👮  Minikube '$PROFILE_NAME' is ready"
else
    echo "🛑  Minikube '${PROFILE_NAME}' is not ready (exit code: ${exit_code})" >&2
    exit 1
fi

echo ""
echo "🙏  A couple of host entries will be added into the hosts file of the system:"
echo ""

sudo bash "${ROOT_DIR}/scripts/add-host-entry-in-host.sh" "${BFF_HOST}" "$(minikube -p "${PROFILE_NAME}" ip)"
sudo bash "${ROOT_DIR}/scripts/add-host-entry-in-host.sh" "${SPA_HOST}" "$(minikube -p "${PROFILE_NAME}" ip)"

echo ""

bash "${ROOT_DIR}/scripts/lib/view-urls.sh"
echo "👌  Setup completed."
