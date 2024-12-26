package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"

	"golang.org/x/image/draw"
)

type settings struct {
	width               int64
	height              int64
	maintainAspectRatio bool
	path                string
	output              string
	
}

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

func getSettings(args []string) (settings, error) {
	settings := settings{
		width:               0,
		height:              0,
		maintainAspectRatio: true,
		path:                "",
	}

	for i, arg := range args {
		switch arg {
		case "-w", "--w", "-width", "--width":
			width, err := strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				return settings, err
			}
			settings.width = width
		case "-h", "--h", "-height", "--height":
			height, err := strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				return settings, err
			}
			settings.height = height
		case "-ar", "--ar", "-aspect", "--aspect":
			maintainAspectRatio, err := strconv.ParseBool(args[i+1])
			if err != nil {
				return settings, err
			}
			settings.maintainAspectRatio = maintainAspectRatio
		case "-p", "--p", "-path", "--path":
			if len(args) <= i+1 {
				return settings, errors.New("path flag provided without appropriate value")
			}
			path := args[i+1]
			settings.path = path
		case "-o", "--o", "-output", "--output":
			if len(args) <= i+1 {
				return settings, errors.New("output flag provided without appropriate value")
			}
			output := args[i+1]
			settings.output = output
		default:
			continue
		}
	}
	return settings, nil
}

func main() {
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
