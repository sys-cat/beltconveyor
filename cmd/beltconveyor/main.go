package main

import (
	"fmt"
	"os"

	"github.com/sys-cat/beltconveyor"
)

func main() {
	item := beltconveyor.New("../output/test.png")
	if err := item.Open(); err != nil {
		fmt.Printf("Open faild, %v\n", err)
		os.Exit(1)
	}
	if err := item.Decode(); err != nil {
		fmt.Printf("Decode faild, %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v", item)

	os.Exit(0)
}
