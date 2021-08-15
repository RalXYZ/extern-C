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
	stack := rand.Intn(10_000)
	fmt.Printf("[%04d] start of a golang procedure\n", stack)

	cStr := C.CString("golang")
	defer C.free(unsafe.Pointer(cStr))
	C.c_procedure(cStr, C.int(stack))

	fmt.Printf("[%04d] end of a golang procedure\n", stack)
}
