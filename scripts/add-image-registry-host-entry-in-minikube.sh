#!/usr/bin/env bash

set -euo pipefail

if [ -z "${PROFILE_NAME:-}" ]; then
    echo "🛑  PROFILE_NAME is undefined" >&2
    exit 1
fi

if [ -z "${REGISTRY_HOST:-}" ]; then
    echo "🛑  REGISTRY_HOST is undefined" >&2
    exit 2
fi

if [ "$#" -lt 1 ]; then
    echo "⚠️   Usage: $0 <registry-host-ip>" >&2
    exit 3
fi

HOST_IP="$1"

minikube -p "${PROFILE_NAME}" ssh -- "sudo bash -c 'grep -qE \"^$HOST_IP\s[^#]*\b$REGISTRY_HOST\b\" /etc/hosts || echo \"$HOST_IP	$REGISTRY_HOST\" >> /etc/hosts'"
