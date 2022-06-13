#!/bin/bash

set -eo pipefail

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi

cd "${DIR}/.."

make -f Makefile.ci ci-install
docker cp "${PROVIDER_BINARY}" /tmp
make -f Makefile.ci ci-remove-container
ls -al /tmp
