# Go related variables.
GOBASE=$(shell pwd)
GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Cross compilation variables
GOOS_linux=linux
GOARCH_linux=amd64
GOOS_windows=windows
GOARCH_windows=amd64

all: image

# Native build (current platform)
test-iowrite: $(GOFILES)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o ./bin/test_iowrite $(GOFILES)

# Cross compile for Linux x86_64 from Windows
test-iowrite-linux: $(GOFILES)
	@set GOOS=$(GOOS_linux) && set GOARCH=$(GOARCH_linux) && \
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o ./bin/test_iowrite_linux_amd64 $(GOFILES)
	@echo "Cross compilation complete: ./bin/test_iowrite_linux_amd64"

# Cross compile for Linux x86_64 (Unix style command for Git Bash/WSL)
test-iowrite-linux-unix: $(GOFILES)
	GOOS=$(GOOS_linux) GOARCH=$(GOARCH_linux) \
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o ./bin/test_iowrite_linux_amd64 $(GOFILES)
	@echo "Cross compilation complete: ./bin/test_iowrite_linux_amd64"

image:
	docker build -t registry/iowrite:v1 .

clean:
	rm -f ./bin/*
	rm -f *.o
	docker rmi registry/iowrite:v1 2>/dev/null || true
