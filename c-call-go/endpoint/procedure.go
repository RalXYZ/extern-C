package main

import "C"

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {}


//export go_procedure
func go_procedure(caller_language string, caller_stack int) {
	rand.Seed(time.Now().Unix())
	callee_stack := rand.Intn(10_000)
	fmt.Printf("[%04d] start of a golang procedure, called by %s [%04d]\n", callee_stack, caller_language, caller_stack)

	fmt.Printf("[%04d] end of a golang procedure\n", callee_stack)
}
