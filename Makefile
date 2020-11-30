  
###########################################################################
## Make file
## @author: Irfan Andriansyah <irfan@99.co>
## @since: 2020.04.10
###########################################################################

EXECUTABLES=git docker
K := $(foreach exec,$(EXECUTABLES),\
    $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

###########################################################################
## Start Development
###########################################################################

init-dev:
	@echo "############################"
	@echo "######### Init Dev #########"
	@echo "############################"
	@make setup-git-hooks

###########################################################################
## Unit Testing
###########################################################################

test: setup-config-test
	@echo "########################################"
	@echo "######### Running Unit Testing #########"
	@echo "########################################"
	CORE_DIR=$(PWD) go test -cover -v -race -coverprofile=coverage.txt -covermode=atomic `go list ./... | grep -v build | grep -v api/pb_generated`

###########################################################################
## Linter Code
###########################################################################


clear-before-lint:
	@rm -f .lint.txt

lint: clear-before-lint
	@echo "##################################"
	@echo "######### Running Golint #########"
	@echo "##################################"
	@go list ./... | grep -v /vendor/ | xargs -L1 "$(GOPATH)/bin/golint" | tee .lint.txt
	@if [ -s .lint.txt ]; then echo "Lint is not successfull"; exit 1; fi

gofmt:
	@echo "###############################"
	@echo "######### Running FMT #########"
	@echo "###############################"
	@go list ./... | grep -v /vendor/ | grep -v api/pb_generated | xargs -L1 go fmt

govet:
	@echo "##################################"
	@echo "######### Running Go Vet #########"
	@echo "##################################"
	@go list ./... | grep -v /vendor/ | grep -v api/pb_generated | xargs -L1 go vet

###########################################################################
## Config File
###########################################################################

setup-git-hooks:
	@cp -b build/.githooks/pre-commit .git/hooks
	@chmod +775 .git/hooks/pre-commit

setup-config-development:
	@rm -f core.ini
	@cp build/config/dev/core.ini core.ini
