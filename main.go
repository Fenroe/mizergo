package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"golang.org/x/image/draw"
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

func resizeImage(img image.Image) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/2, img.Bounds().Max.Y/2))
	draw.CatmullRom.Scale(rgba, rgba.Rect, img, img.Bounds(), draw.Over, nil)
	return rgba
}

func saveImage(path string, img *image.RGBA) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = jpeg.Encode(file, img, nil)
	return err
}

func getPercentage(x, y int) float64 {
	return float64((y / x) * 100)
}

func main() {
	fmt.Println("Hi banana!")
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(wd)
	args := os.Args[1:]
	path := args[0]
	image, err := getImageFromFilePath(path)
	if err != nil {
		fmt.Println("Oh no!")
		return
	}
	fmt.Printf("width:%v, height:%v\n", image.Bounds().Dx(), image.Bounds().Dy())
	resizedImage := resizeImage(image)
	err = saveImage(path, resizedImage)
	if err != nil {
		fmt.Println(err)
		return
	}
	image, err = getImageFromFilePath(path)
	if err != nil {
		fmt.Println("Oh no!")
		return
	}
	fmt.Printf("width:%v, height:%v\n", image.Bounds().Dx(), image.Bounds().Dy())
}
