package perl

/*
#cgo CFLAGS: -D_THREAD_SAFE -pthread -I../vendor/perl-5.20.1
#cgo LDFLAGS: -fstack-protector -L/usr/local/lib -L/Users/ikent/dev/src/github.com/ian-kent/purl/vendor/perl-5.20.1 -lperl -ldl -lm -lutil -lc -fno-common -DPERL_DARWIN -fno-strict-aliasing -pipe -fstack-protector -I/usr/local/include -I/usr/local/lib/perl5/5.20.1/darwin-2level/CORE
#include "c/purl.h"
*/
import "C"

func Test() {
	C.RunPurl()
}
