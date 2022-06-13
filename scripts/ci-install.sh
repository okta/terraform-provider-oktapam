#!/bin/bash

set -eo pipefail

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi

cd "${DIR}/.."

# remove container first, if it happens to exist
make -f Makefile.ci ci-remove-container || echo "Container doesn't exist. Continue."
# now do install
make -f Makefile.ci ci-install
# copy artifact to host
xet -x
docker cp "${INSTALL_CONTAINER_NAME}:${INSTALL_TARGET}/${PROVIDER_BINARY}" /tmp
# remove container
make -f Makefile.ci ci-remove-container
# display contents of /tmp
ls -al /tmp
