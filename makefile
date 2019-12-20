BIN=go-ots
BIN_OUT=bin/$(BIN)
BIN_INSTALL=/usr/local/bin/$(BIN)
GOTEST=@go test ./...
GOBUILD=@go build -a -o $(BIN_OUT) main.go

all: banner test build

.PHONY: banner
banner:
	@echo "  ____  ___         ___ _____ ____  "
	@echo " / ___|/ _ \       / _ \_   _/ ___| "
	@echo "| |  _| | | |_____| | | || | \___ \ "
	@echo "| |_| | |_| |_____| |_| || |  ___) |"
	@echo " \____|\___/       \___/ |_| |____/ "                                     

.PHONY: test
test:
	@echo "\n==========================================================================="
	@echo "Running tests"
	@echo "===========================================================================\n"
	$(GOTEST)
	@echo "Done"

.PHONY: build
build:
	@echo "\n==========================================================================="
	@echo "Building Binaries"
	@echo "===========================================================================\n"
	$(GOBUILD)
	@echo "Done"

.PHONY: install
install:
	@echo "\n==========================================================================="
	@echo "Installing to $(BIN_INSTALL)"
	@echo "===========================================================================\n"
	@cp $(BIN_OUT) $(BIN_INSTALL)
	@chmod +x $(BIN_INSTALL)
	@echo "Done"