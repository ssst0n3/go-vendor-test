package main

import (
	"st0n3"
	"fmt"
)

func main() {
    source()
    st0n3.Crash()
}

func source() {
	a()
}

func a() {
    sink()
}

func sink() {
    fmt.Println("sink")
}
