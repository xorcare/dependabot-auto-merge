package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xorcare/pointer"
)

func main() {
	fmt.Println("Hello", pointer.Bool(true), uuid.New())
}
