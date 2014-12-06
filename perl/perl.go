package perl

/*
#cgo CFLAGS: -D_THREAD_SAFE -pthread -I../vendor/perl-5.20.1
#cgo LDFLAGS: -fstack-protector -L/usr/local/lib -L$GOPATH/github.com/ian-kent/purl/vendor/perl-5.20.1 -lperl -ldl -lm -lutil -lc -fno-common -fno-strict-aliasing -pipe -fstack-protector -I/usr/local/include -I/usr/local/lib/perl5/5.20.1/darwin-2level/CORE
#include "c/purl.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

var perlMutex sync.Mutex
var xsMap = make(map[string]*func(...string) string)

// Purl is a Perl interpreter instance.
//
// You must call Init before calling any other functions.
// You must call Destroy when finished with the instance.
type Purl struct {
	init    bool
	destroy bool
}

// Init initialises the Perl interpreter instance
func (p *Purl) Init() {
	perlMutex.Lock()
	if !p.init {
		C.PurlInit()
		p.init = true
		p.destroy = false
	}
	perlMutex.Unlock()
}

// Destroy destroys the Perl interpreter instance
func (p *Purl) Destroy() {
	perlMutex.Lock()
	if !p.destroy {
		C.PurlDestroy()
		p.destroy = true
		p.init = false
	}
	perlMutex.Unlock()
}

// RegisterXS makes a Go function available in Perl
// via an XS callback.
//
// It currently only supports a variadic string input
// and a single scalar string output.
func (p *Purl) RegisterXS(name string, f func(...string) string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	xsMap[name] = &f
	p.Eval(`
package main {
	*{"` + name + `"} = sub { Purl::XS->Invoke("` + name + `") };
}
`)
}

// Eval evaluates Perl code.
//
// Interpreter state is persisted until Destroy is called.
func (p *Purl) Eval(src string) string {
	csrc := C.CString(src)
	defer C.free(unsafe.Pointer(csrc))

	cres := C.EvalPerl(csrc)
	if cres == nil {
		return ""
	}

	// Not sure if this is required? causes memory issues
	//defer C.free(unsafe.Pointer(cres))
	return C.GoString(cres)
}

func newString(s string) *C.SV {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	str := C.Perl_newSVpvn(cs, C.STRLEN(len(s)))
	return str
}
