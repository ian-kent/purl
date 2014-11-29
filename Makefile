all: build

build: deps
	go build .

work: deps
	go build -work .

deps:
	mkdir -p vendor
	if [ ! -e vendor/perl-5.20.1.tar.gz ]; then wget -O vendor/perl-5.20.1.tar.gz http://www.cpan.org/src/5.0/perl-5.20.1.tar.gz; fi
	if [ ! -e vendor/perl-5.20.1 ]; then cd vendor; tar -xf perl-5.20.1.tar.gz; cd perl-5.20.1; sh ./Configure -de; make; fi

.PHONY: all build work deps
