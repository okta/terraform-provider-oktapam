HOSTNAME=okta.com
NAMESPACE=pam
NAME=oktapam
BINARY=terraform-provider-${NAME}

VERSION=0.6.0

OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)
PLUGIN_DIR=~/.terraform.d/plugins
DOCGEN_RESOURCES_DIR=docgen-resources
GOPRIVATE="github.com/atko-pam,github.com/okta"
SET_VERSION=-ldflags "-X github.com/okta/terraform-provider-oktapam/oktapam/version.Version=${VERSION}"

ifneq ($(DEBUG), )
  GOFLAGS :=${GOFLAGS} -gcflags=all="-N -l"
else
  GOFLAGS :=${GOFLAGS} -trimpath
endif

.DEFAULT_GOAL := install

build:
	go build ${GOFLAGS} -ldflags "-X github.com/okta/terraform-provider-oktapam/oktapam/version.Version=${VERSION}dev" -o ${BINARY}

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
	go test ./... -v $(TESTARGS)

testacc: 
# TESTARGS here can be used to pass arbitrary flags to go test, e.g. '-run TestMyTest'
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m   

testaccpam:
# TESTARGS here can be used to pass arbitrary flags to go test, e.g. '-run TestMyTest'
	TF_ACC=1 TF_ACC_PAM=1 go test ./... -v $(TESTARGS) -timeout 120m   


generate:
	go generate ./...

check-generate:
	go generate ./...
	git diff --compact-summary --exit-code || \
	  (echo; echo "Unexpected difference in directories after code generation. Run 'go generate ./...' command and commit."; exit 1)
updatedep:
	ifndef dep
		$(error you must specify dep=<your dependency>)
	endif
	GOPRIVATE=${GOPRIVATE} go mod tidy
	GOPRIVATE=${GOPRIVATE} go get $(dep)@latest
	GOPRIVATE=${GOPRIVATE} go mod tidy
	GOPRIVATE=${GOPRIVATE} go mod vendor