###############################################################################
# Tooling configuration. 
#
# All executable names should be defined as variables so that they can be
# overloaded.
###############################################################################

GO            ?= go
DOCKER        ?= docker
GOLANGCI_LINT ?= golangci-lint

###############################################################################
# Build configuration
###############################################################################

GO_MODULE_NAME      ?= $(shell head -1 go.mod | awk '{print $$2}')

.PHONY: help
help:
	@echo
	@echo Usage: make [target...]
	@echo
	$(eval doc_expanded := $(shell grep -E -h '^(.PHONY:|# TARGETDOC:) .* #' $(MAKEFILE_LIST) \
		| sed -E -n 's/(\.PHONY|# TARGETDOC): (.*) # (.*)/\\033[36m\2\\033[m:  \3\\n/'p | column -c2 -t -s :))
	@printf ' $(doc_expanded)'

.PHONY: run-frontend # Start the frontend server
run-frontend:
	@go run cmd/frontend/*.go
