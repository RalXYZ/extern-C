package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L${SRCDIR}/endpoint/ -lc_procedure
#include <stdlib.h>
#include "endpoint/c_procedure.h"
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	caller_stack := rand.Intn(10_000)
	fmt.Printf("[%04d] start of golang main procedure\n", caller_stack)

	caller_language := C.CString("golang")
	defer C.free(unsafe.Pointer(caller_language))
	C.c_procedure(caller_language, C.int(caller_stack))

	fmt.Printf("[%04d] end of golang main procedure\n", caller_stack)
}
