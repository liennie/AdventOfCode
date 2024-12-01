BINARIES := $(foreach bin,$(shell ls cmd),bin/$(bin))

TEST_SOURCES := $(shell find . -type f -name '*.go' -print) go.mod go.sum
TEST_DIRS := $(shell go list ./...)
PWD := $(shell pwd)

all: $(BINARIES)

bin/%: cmd/%/*
	@echo $@
	@GOBIN=$(PWD)/bin go install -tags purego ./cmd/$(notdir $@)/...

test: $(TEST_SOURCES)
	@go test -cover -race $(LDFLAGS) $(TEST_DIRS)

clean:
	@if [ -d ./bin ] ; then rm -r ./bin ; fi

.PHONY: test clean
