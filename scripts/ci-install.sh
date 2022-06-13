#!/bin/bash

set -eo pipefail

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi

cd "${DIR}/.."

# create artifact directory
mkdir artifacts
# remove container first, if it happens to exist
make -f Makefile.ci ci-remove-container || echo "Container doesn't exist. Continue."
# now do install and copy artifact to host
make -f Makefile.ci ci-artifact-copy
# remove container
make -f Makefile.ci ci-remove-container
# display contents of artifact dir
ls -al artifacts
