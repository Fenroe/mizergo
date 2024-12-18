package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func getImageFromFilePath(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	return image, err
}

func main() {
	fmt.Println("Hi banana!")
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(wd)
}
