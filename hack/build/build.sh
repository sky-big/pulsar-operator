#!/usr/bin/env bash

set -o errexit
set -o nounset
# set -o pipefail

# dir
bin_dir="$(pwd)/../../bin"
mkdir -p ${bin_dir} || true

# build function
go_build_operator () {
	echo "[START] building "pulsar operator"..."
	# We’re disabling cgo which gives us a static binary.
	# This is needed for building minimal container based on alpine image.
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v $GO_BUILD_FLAGS -o ${bin_dir}/pulsar-operator -installsuffix cgo -ldflags "$go_ldflags" ../../cmd/manager/
	echo "[END] building "pulsar operator"..."
}

# check golang
if ! which go > /dev/null; then
	echo "golang needs to be installed"
	exit 1
fi

GIT_SHA=`git rev-parse --short HEAD || echo "GitNotFound"`

gitHash="github.com/sky-big/pulsar-operator/version.GitSHA=${GIT_SHA}"

go_ldflags="-X ${gitHash}"

GO_BUILD_FLAGS="$@" go_build_operator
