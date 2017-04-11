.PHONY: clean help default build

VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
VCERTVERSION := $(shell grep "const VCertVersion " version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)

BUILD_DIR=$(shell pwd)

VCDIR=$(BUILD_DIR)/embedded

BINNAME=vcert

help:
	@echo 'Management commands for the govcert library:'
	@echo
	@echo 'Usage:'
	@echo '    make build     Use go-bindata to create embedable versions of the VCert binary.'
	@echo '    make clean     Clean build assets.'
	@echo

default: clean build

build: $(VCDIR)/vcert_darwin_amd64.go $(VCDIR)/vcert_darwin_386.go $(VCDIR)/vcert_linux_amd64.go $(VCDIR)/vcert_linux_386.go $(VCDIR)/vcert_windows_amd64.go $(VCDIR)/vcert_windows_386.go

bindata:
	go get -u -v github.com/jteeuwen/go-bindata/go-bindata

$(VCDIR)/vcert_darwin_amd64.go: bindata vcert/bins/$(VCERTVERSION)/darwin/*ert
	cp vcert/bins/$(VCERTVERSION)/darwin/*ert vcert/bins/$(VCERTVERSION)/vcert/$(BINNAME)
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o $(VCDIR)/vcert_darwin_amd64.go -pkg embedded vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_darwin_386.go: bindata vcert/bins/$(VCERTVERSION)/darwin/vcert86
	cp vcert/bins/$(VCERTVERSION)/darwin/*86 vcert/bins/$(VCERTVERSION)/vcert/$(BINNAME)
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o $(VCDIR)/vcert_darwin_386.go -pkg embedded vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_linux_amd64.go: bindata vcert/bins/$(VCERTVERSION)/linux/vcert
	cp vcert/bins/$(VCERTVERSION)/linux/*ert vcert/bins/$(VCERTVERSION)/vcert/$(BINNAME)
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o $(VCDIR)/vcert_linux_amd64.go -pkg embedded vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_linux_386.go: bindata vcert/bins/$(VCERTVERSION)/linux/vcert86
	cp vcert/bins/$(VCERTVERSION)/linux/*86 vcert/bins/$(VCERTVERSION)/vcert/$(BINNAME)
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o $(VCDIR)/vcert_linux_386.go -pkg embedded vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_windows_amd64.go: bindata vcert/bins/$(VCERTVERSION)/windows/vcert.exe
	cp vcert/bins/$(VCERTVERSION)/windows/*rt.exe vcert/bins/$(VCERTVERSION)/vcert/$(BINNAME).exe
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o $(VCDIR)/vcert_windows_amd64.go -pkg embedded vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_windows_386.go: bindata vcert/bins/$(VCERTVERSION)/windows/vcert86.exe
	cp vcert/bins/$(VCERTVERSION)/windows/*86.exe vcert/bins/$(VCERTVERSION)/vcert/$(BINNAME).exe
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o $(VCDIR)/vcert_windows_386.go -pkg embedded vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

clean:
	rm $(VCDIR)/vcert_*.go
