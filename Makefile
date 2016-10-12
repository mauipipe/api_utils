#! /usr/bin/make

BUILDPATH=$(CURDIR)/../../../../
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=main

export GOPATH=$(BUILDPATH)

DEPENDENCIES=\
	github.com/gorilla/mux \
	github.com/onsi/gomega \
    github.com/fzipp/gocyclo \

api_utils:
	@echo "make";
	@echo $(BUILDPATH)

	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi
	@$(GOGET) $(DEPENDENCIES)


    go test -v ./ -covermode=count -coverprofile=coverage.out $HOME/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
    gocyclo ./