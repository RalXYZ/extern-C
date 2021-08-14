package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L${SRCDIR}/../c/ -lc_procedure
#include <stdlib.h>
#include "../c/c_procedure.h"
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"
	"math/rand"
)

func main() {}


//export go_procedure
func go_procedure(s string) {
	rand.Seed(time.Now().Unix())
	stack := rand.Intn(10_000)
	fmt.Printf("[%04d] start of a golang procedure, called by %s\n", stack, s)

	cStr := C.CString("golang")
	C.c_procedure(cStr)
	C.free(unsafe.Pointer(cStr))

	fmt.Printf("[%04d] end of a golang procedure\n", stack)
}
