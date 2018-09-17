# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=bose
PREFIX=/usr/local

mkdir=mkdir -p

all: build
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./src
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v src/
	./src/$(BINARY_NAME)
install:
	@$(mkdir) $(DESTDIR)$(PREFIX)/bin
	@cp bin/$(BINARY_NAME) $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)
uninstall:
	@rm -f $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)