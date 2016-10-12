#! /usr/bin/make

BUILDPATH=$(CURDIR)/../../../../
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get
COVERAGE_FILE = coverage.out

EXENAME=main

export GOPATH=$(BUILDPATH)

DEPENDENCIES=\
	github.com/gorilla/mux \
	github.com/onsi/gomega \
	github.com/fzipp/gocyclo \

all: dependencies test clean

dependencies:
	@echo "Retrieving dependencies"
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi
	@$(GOGET) $(DEPENDENCIES)
	@echo "Dependencies Ok"

test:
	@echo "Running test"
	go test -v ./ -covermode=count -coverprofile=$(COVERAGE_FILE) $HOME/gopath/bin/goveralls -service=travis-ci -repotoken $COVERALLS_TOKEN
	@echo "Cyclomatic complexity"
	@echo $(BUILDPATH)bin/gocyclo
	$(BUILDPATH)bin/gocyclo -over 12 .

clean:
	@echo "Deleting "
	rm -rf $(COVERAGE_FILE)

