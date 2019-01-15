#!/bin/bash

# Copyright 2019 The Jetstack cert-manager contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

_SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
REPO_ROOT="${_SCRIPT_ROOT}/../../.."

# This file contains common definitions that are re-used in other scripts

KIND_TARGET="${KIND_TARGET:-//test/e2e:e2e-1.11}"
KIND_CLUSTER_NAME="${KIND_TARGET##*:}"
KIND_CONTAINER_NAME="kind-${KIND_CLUSTER_NAME}-control-plane"

# DOCKER_REPO is the docker repo to use for cert-manager images, either when
# building or deploying cert-manager using these scripts.
export DOCKER_REPO="quay.io/jetstack"

# DOCKER_TAG is the docker tag to use for the cert-manager images.
# This defaults to 'build' so it doesn't conflict with images built for any
# other purpose
export DOCKER_TAG="build"
export APP_VERSION="${DOCKER_TAG}"

if [ ! "${CM_DEPS_LOADED:-}" = "1" ]; then
    # Build all e2e test dependencies
    bazel build \
        //hack/bin:helm \
        //test/e2e:e2e.test \
        //test/e2e/charts:images \
        "${KIND_TARGET}"

    genfiles="$(bazel info bazel-genfiles)"
    # build 'kind'
    export KIND="${genfiles}/hack/bin/kind"
    export KUBECONFIG="${genfiles}/test/e2e/${KIND_CLUSTER_NAME}.kubeconfig"

    export CM_DEPS_LOADED="1"
fi