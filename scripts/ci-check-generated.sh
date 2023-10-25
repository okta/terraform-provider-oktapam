#!/bin/bash

set -eo pipefail

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi

cd "${DIR}/.."

export TF_ACC_TERRAFORM_VERSION=""

TF_BIN=$(which terraform)

if [[ -x "${TF_BIN}" ]]; then
    echo "found terraform at ${TF_BIN}, setting TF_ACC_TERRAFORM_PATH=${TF_BIN}"
    # This path needs to be added to the Docker run command
    export TF_ACC_TERRAFORM_PATH="${TF_BIN}"
    export TF_ACC_TERRAFORM_VERSION="$(cat .terraform-version)"
fi

make -f Makefile.ci ci-check-generate
