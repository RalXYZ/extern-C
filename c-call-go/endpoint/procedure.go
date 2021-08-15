package main

import "C"

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {}


//export go_procedure
func go_procedure(s string, caller_stack int) {
	rand.Seed(time.Now().Unix())
	stack := rand.Intn(10_000)
	fmt.Printf("[%04d] start of a golang procedure, called by %s [%04d]\n", stack, s, caller_stack)

	fmt.Printf("[%04d] end of a golang procedure\n", stack)
}
