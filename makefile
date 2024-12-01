SOURCES := $(shell find cmd/prepare -type f -print)
PWD := $(shell pwd)

bin/prepare: $(SOURCES) go.mod go.sum
	@GOBIN=$(PWD)/bin go install ./cmd/prepare/...

clean:
	@if [ -d ./bin ] ; then rm -r ./bin ; fi

.PHONY: clean
