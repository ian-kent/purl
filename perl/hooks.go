package perl

import "C"
import (
	"fmt"
	"unsafe"
)

//export PurlTest
func PurlTest() {
	fmt.Println("Hi from Perl::Test!")
}

//export PurlXSHook
func PurlXSHook(fnAddr unsafe.Pointer, delegate *C.char, narg C.int, svArgsPtr unsafe.Pointer) (svOutResult unsafe.Pointer) {
	dn := C.GoString(delegate)
	if d, ok := xsMap[dn]; ok {
		df := *d
		sv := newString(df())
		svOutResult = unsafe.Pointer(sv)
	} else {
		panic("Unknown XS hook")
	}
	return
}
