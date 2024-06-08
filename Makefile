SHELL=/bin/bash
PACKAGE_NAME=godating-dealls-service
BUILD_DIR=./build/compile

# Targets
build/service:
	@echo "compile service $(PACKAGE_NAME)"
	go build -o $(BUILD_DIR)/$(PACKAGE_NAME) cmd/main.go

run/service:
	@echo "run service $(PACKAGE_NAME)"
	./$(BUILD_DIR)/$(PACKAGE_NAME) main.go

brun/service: build/service run/service