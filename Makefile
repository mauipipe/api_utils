BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=main

export GOPATH=$(CURDIR)

myname:
	@echo "api_utils"

makedir:
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

get:
	@$(GOGET) github.com/gorilla/mux
	@$(GOGET) github.com/onsi/gomega

build:
	@echo "start building..."
	$(GOINSTALL) $(EXENAME)
	@echo "Completed"
