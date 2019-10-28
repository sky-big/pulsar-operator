# Build binary and image.
#
# Example:
#   make
#   make all
all: build
.PHONY: all

# Build the binaries
#
# Example:
#   make build
build:
	cd hack/build && sh build.sh
.PHONY: build

# Build the docker image
#
# Example:
#   make image
image:
	docker/build-image.sh
.PHONY: image

# Push the docker image
#
# Example:
#   make push
push:
	docker/push-image.sh
.PHONY: push

# generate code(zz_generated*)
# generate go mod list to vendor
# Example:
#   make generate
generate:
	operator-sdk generate k8s
	operator-sdk generate openapi
	go mod vendor
.PHONY: generate

# install to kubernetes
# Example:
#   make install
install:
	deploy/install.sh
.PHONY: install

# uninstall from kubernetes
# Example:
#   make uninstall
uninstall:
	deploy/uninstall.sh
.PHONY: uninstall

# start local test
# Example:
#   make start-local
start-local:
	test/local/install.sh
.PHONY: start-local

# stop local test
# Example:
#   make stop-local
stop-local:
	test/local/uninstall.sh
.PHONY: stop-local

# start pvc test
# Example:
#   make start-pvc
start-pvc:
	test/pvc/install.sh
.PHONY: start-pvc

# stop pvc test
# Example:
#   make stop-pvc
stop-pvc:
	test/pvc/uninstall.sh
.PHONY: stop-pvc

# clean all binaries
#
# Example:
#   make clean
clean:
	rm -rf ./bin
.PHONY: clean
