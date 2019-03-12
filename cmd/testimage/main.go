package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	f, err := os.OpenFile("../output/test.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	png.Encode(f, img)
	os.Exit(0)
}
