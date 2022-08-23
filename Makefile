HOSTNAME=okta.com
NAMESPACE=pam
NAME=oktapam
BINARY=terraform-provider-${NAME}
# On verion changes, update tag-checks.yml
VERSION=0.2.1
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)
PLUGIN_DIR=~/.terraform.d/plugins
# Must update binary at TERRAFORM_PATH when TERRAFORM_VERSION is changed
TERRAFORM_PATH=testacc/terraform-cli/terraform
TERRAFORM_VERSION=1.2.7

SET_VERSION=-ldflags "-X github.com/okta/terraform-provider-oktapam/oktapam/version.Version=${VERSION}"

.DEFAULT_GOAL := install

build:
	go build -ldflags "-X github.com/okta/terraform-provider-oktapam/oktapam/version.Version=${VERSION}dev" -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_darwin_arm64	
	GOOS=freebsd GOARCH=386 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build ${SET_VERSION} -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ${PLUGIN_DIR}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ${PLUGIN_DIR}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

link_legacy:
	mkdir -p ${PLUGIN_DIR}
	ln -s ${PLUGIN_DIR}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY} ${PLUGIN_DIR}/${BINARY}

test: 
# TESTARGS here can be used to pass arbitrary flags to go test, e.g. '-run TestMyTest'
	TF_ACC_TERRAFORM_PATH=TERRAFORM_PATH TF_ACC_TERRAFORM_VERSION=TERRAFORM_VERSION go test ./... -v $(TESTARGS)

testacc: 
# TESTARGS here can be used to pass arbitrary flags to go test, e.g. '-run TestMyTest'
	TF_ACC=1 TF_ACC_TERRAFORM_PATH=TERRAFORM_PATH TF_ACC_TERRAFORM_VERSION=TERRAFORM_VERSION go test ./... -v $(TESTARGS) -timeout 120m

generate:
	go generate ./...

check-generate:
	go generate ./...
	git diff --compact-summary --exit-code || \
	  (echo; echo "Unexpected difference in directories after code generation. Run 'go generate ./...' command and commit."; exit 1)
