package levenshtein

/*
#cgo CFLAGS: -I .

#include <stdlib.h>
#include "levenshtein.h"
*/
import "C"

import (
	"unsafe"
)

func Distance(a, b string) int {
	ca := C.CString(a)
	defer C.free(unsafe.Pointer(ca))

	cb := C.CString(b)
	defer C.free(unsafe.Pointer(cb))

	cdist := C.levenshtein(a, C.int(len(a)), b, C.int(len(b)))
	return int(cdist)
}
