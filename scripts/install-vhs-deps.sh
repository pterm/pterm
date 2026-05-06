#!/usr/bin/env bash

set -Eeuo pipefail

source "$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)/animation-common.sh"

VHS_DOCKER_IMAGE="${VHS_DOCKER_IMAGE:-ghcr.io/charmbracelet/vhs:latest}"

group "Install VHS dependencies"
require_cmd docker
info "Pulling $VHS_DOCKER_IMAGE"
docker pull "$VHS_DOCKER_IMAGE"
end_group

success "VHS Docker image is ready"
