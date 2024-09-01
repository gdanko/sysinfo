GOPATH := $(shell go env GOPATH)
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
VERSION := 0.3.3

.PHONY: all build install

all: build install

.PHONY: mod-tidy
mod-tidy:
	@echo "${GOPATH}"
	go mod tidy

.PHONY: build OS ARCH
build: mod-tidy clean
	@echo "================================================="
	@echo "Building SysInfo"
	@echo "=================================================\n"

	GOOS=${GOOS} GOARCH=${GOARCH} go build -o sysinfo
	sleep 1
	tar czvf sysinfo_${VERSION}_${GOOS}_${GOARCH}.tgz sysinfo

.PHONY: clean
clean:
	@echo "================================================="
	@echo "Cleaning SysInfo"
	@echo "=================================================\n"
	@if [ -f sysinfo ]; then \
		rm -f sysinfo; \
	fi

.PHONY: install
install:
	@echo "================================================="
	@echo "Installing SysInfo in ${GOPATH}/bin"
	@echo "=================================================\n"

	go install -race
