.PHONY: clean help default build

VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
VCERTVERSION := $(shell grep "const VCertVersion " version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)

BUILD_DIR=$(shell pwd)

VCDIR=vcert

help:
	@echo 'Management commands for the govcert library:'
	@echo
	@echo 'Usage:'
	@echo '    make build     Use go-bindata to create embedable versions of the VCert binary.'
	@echo '    make clean     Clean build assets.'
	@echo

default: clean build

build: $(VCDIR)/vcert_darwin_amd64.go $(VCDIR)/vcert_darwin_386.go $(VCDIR)/vcert_linux_amd64.go $(VCDIR)/vcert_linux_386.go $(VCDIR)/vcert_windows_amd64.go $(VCDIR)/vcert_windows_386.go

$(VCDIR)/vcert_darwin_amd64.go: vcert/bins/$(VCERTVERSION)/darwin/VCert
	cp vcert/bins/$(VCERTVERSION)/darwin/VCert vcert/bins/$(VCERTVERSION)/vcert/VCert
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o ../../vcert_darwin_amd64.go -pkg vcert vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_darwin_386.go: vcert/bins/$(VCERTVERSION)/darwin/VCert86
	cp vcert/bins/$(VCERTVERSION)/darwin/VCert86 vcert/bins/$(VCERTVERSION)/vcert/VCert
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o ../../vcert_darwin_386.go -pkg vcert vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_linux_amd64.go: vcert/bins/$(VCERTVERSION)/linux/VCert
	cp vcert/bins/$(VCERTVERSION)/linux/VCert vcert/bins/$(VCERTVERSION)/vcert/VCert
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o ../../vcert_linux_amd64.go -pkg vcert vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_linux_386.go: vcert/bins/$(VCERTVERSION)/linux/VCert86
	cp vcert/bins/$(VCERTVERSION)/linux/VCert86 vcert/bins/$(VCERTVERSION)/vcert/VCert
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o ../../vcert_linux_386.go -pkg vcert vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_windows_amd64.go: vcert/bins/$(VCERTVERSION)/windows/VCert.exe
	cp vcert/bins/$(VCERTVERSION)/windows/VCert.exe vcert/bins/$(VCERTVERSION)/vcert/VCert.exe
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o ../../vcert_windows_amd64.go -pkg vcert vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

$(VCDIR)/vcert_windows_386.go: vcert/bins/$(VCERTVERSION)/windows/VCert86.exe
	cp vcert/bins/$(VCERTVERSION)/windows/VCert86.exe vcert/bins/$(VCERTVERSION)/vcert/VCert.exe
	cd vcert/bins/$(VCERTVERSION) && go-bindata -o ../../vcert_windows_386.go -pkg vcert vcert
	rm vcert/bins/$(VCERTVERSION)/vcert/*

clean:
	rm $(VCDIR)/vcert_*
