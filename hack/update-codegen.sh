#!/usr/bin/env bash

# Modified from https://github.com/kubernetes-sigs/kueue/blob/065451d907fa27a0647436505b3cac38718ef136/hack/update-codegen.sh
# Apache-2.0, Copyright 2023 The Kubernetes Authors

set -o errexit
set -o nounset
set -o pipefail

GO_CMD=${1:-go}
PKG_ROOT=$(realpath "$(dirname ${BASH_SOURCE[0]})/..")
CODEGEN_PKG=$($GO_CMD list -m -f "{{.Dir}}" k8s.io/code-generator)

cd $PKG_ROOT

source "${CODEGEN_PKG}/kube_codegen.sh"

ln -s . github.com
ln -s .. w-h-a
trap "rm github.com && rm w-h-a" EXIT

kube::codegen::gen_helpers \
  --input-pkg-root github.com/w-h-a/crd/pkg/apis \
  --boilerplate /dev/null \
  --output-base "${PKG_ROOT}"

kube::codegen::gen_client \
  --input-pkg-root github.com/w-h-a/crd/pkg/apis \
  --output-pkg-root github.com/w-h-a/crd/pkg/client \
  --boilerplate /dev/null \
  --output-base "${PKG_ROOT}" \
  --with-watch \
  --with-applyconfig

# clean up temporary libraries added in go.mod by code-generator
"${GO_CMD}" mod tidy