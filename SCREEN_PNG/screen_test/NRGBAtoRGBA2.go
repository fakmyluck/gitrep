package main

import (
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"os"
)

type picture struct {
	img image.Image
	ptr *image.RGBA
	dim dimensions
	sym string
}

type dimensions struct {
	Dx, Dy int //dlinna shirina
}

func main() {

	filename := "888"
	imag := openimg(filename)

	// save file
	outputFile, err := os.Create(filename + ".png")
	if err != nil {
		fmt.Println(err)
		return
	}
	png.Encode(outputFile, imag)
	fmt.Println("File created.")

}

func openimg(filename string) *image.RGBA {

	RAWscr, err := os.Open("../pics/" + filename + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer RAWscr.Close()

	img, _, err := image.Decode(RAWscr)
	if err != nil {
		fmt.Println(err)
	}
	nrgba := img.(*image.NRGBA)

	rgba := image.NewRGBA(nrgba.Rect)

	rgba.Stride = nrgba.Stride

	rgba.Pix = nrgba.Pix
	return rgba

}
