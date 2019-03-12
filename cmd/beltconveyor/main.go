package main

import (
	"fmt"
	"os"

	"github.com/sys-cat/beltconveyor"
)

func main() {
	item := beltconveyor.New("")
	fmt.Printf("%#v", item)

	os.Exit(0)
}
