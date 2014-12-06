all: fmt build

fmt:
	go fmt ./...

build: deps
	cp -f perl/perl.go perl/perl._go
	sed -i.bak 's|$$GOPATH|${GOPATH}/src|g' perl/perl.go
	rm perl/perl.go.bak
	#FIXME errors ignored so perl.go is restored
	-go build .
	mv -f perl/perl._go perl/perl.go

deps: libperl

libperl: vendor/perl-5.20.1/perl

vendor/perl-5.20.1/perl: vendor/perl-5.20.1/Makefile
	cd vendor/perl-5.20.1; make

vendor/perl-5.20.1/Makefile: |vendor/perl-5.20.1
	cd vendor/perl-5.20.1; sh ./Configure -de

vendor/perl-5.20.1: vendor/perl-5.20.1.tar.gz
	cd vendor; tar -xf perl-5.20.1.tar.gz

vendor/perl-5.20.1.tar.gz: |vendor
	wget -O vendor/perl-5.20.1.tar.gz http://www.cpan.org/src/5.0/perl-5.20.1.tar.gz

vendor:
	mkdir -p vendor

.PHONY: all fmt build deps libperl
