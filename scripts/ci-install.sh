#!/bin/bash

set -eo pipefail

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi

cd "${DIR}/.."

# create artifact directory
mkdir artifacts
mkdir more_artifacts
# remove container first, if it happens to exist
make -f Makefile.ci ci-remove-container || echo "Container doesn't exist. Continue."
# now do install and copy artifact to host
make -f Makefile.ci ci-artifact-copy
# remove container
make -f Makefile.ci ci-remove-container
# display contents of artifact dir
ls -al artifacts
# upload artifact
buildkite-agent artifact upload ./artifacts/* "s3://${BUILDKITE_S3_ARTIFACT_EXCHANGE_BUCKET}/${BUILDKITE_PIPELINE_NAME}/"
export BUILDKITE_ARTIFACT_UPLOAD_DESTINATION="s3://${BUILDKITE_S3_ARTIFACT_EXCHANGE_BUCKET}/${BUILDKITE_PIPELINE_NAME}/"
buildkite-agent artifact download "artifacts/*" ./more_artifacts
ls -al more_artifacts
