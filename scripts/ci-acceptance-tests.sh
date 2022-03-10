#!/bin/bash

set -eo pipefail

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi

cd "${DIR}/.."

make -f Makefile.ci ci-acceptance-test
