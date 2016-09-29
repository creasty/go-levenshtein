package levenshtein

/*
#cgo CFLAGS: -I .

#include <stdint.h>
#include <stddef.h>

#include "levenshtein.h"
*/
import "C"

import (
	"unsafe"
)

func Distance(a, b string) int {
	if a == b {
		return 0
	}

	aRune := []rune(a)
	bRune := []rune(b)
	aLen := len(aRune)
	bLen := len(bRune)

	if aLen == 0 {
		return bLen
	}
	if bLen == 0 {
		return aLen
	}

	ca := (*C.int32_t)(unsafe.Pointer(&aRune[0]))
	cb := (*C.int32_t)(unsafe.Pointer(&bRune[0]))

	cdist := C.levenshtein(ca, C.size_t(len(aRune)), cb, C.size_t(len(bRune)))
	return int(cdist)
}
