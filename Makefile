#! /usr/bin/make

BUILDPATH=$(CURDIR)/../../../../
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=main

export GOPATH=$(BUILDPATH)

api_utils:
	@echo "make";
	@echo $(BUILDPATH)

	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

	@$(GOGET) github.com/gorilla/mux
	@$(GOGET) github.com/onsi/gomega
