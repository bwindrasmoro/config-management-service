OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

BINARY_NAME := config-management-service
RUN_CMD := ./$(BINARY_NAME)

ifeq ($(OS),windows)
	BINARY_NAME := config-management-service.exe
	RUN_CMD := $(BINARY_NAME)
	CLEAN_CMD := if exist $(BINARY_NAME) (echo Binary found, deleting... && del $(BINARY_NAME))

	CHECK_N_RUN_CMD := if exist $(BINARY_NAME) ($(RUN_CMD)) else (echo Binary not found, please execute 'make build' to build the binary.)
else
	CLEAN_CMD := [ -f $(BINARY_NAME) ] && echo Binary found, deleting... && rm $(BINARY_NAME)

	CHECK_N_RUN_CMD := [ -f $(BINARY_NAME) ] && $(RUN_CMD) || echo Binary not found, please execute 'make build' to build the binary.
endif

.PHONY: build run test

build:
	@echo Building in $(OS) env
	@$(CLEAN_CMD)
	@echo Building binary
	@go build -o $(BINARY_NAME) main.go
	@echo Build completed: $(BINARY_NAME)

run:	
	@echo Running in $(OS) env
	@$(CHECK_N_RUN_CMD)

test:
	go test ./test
	