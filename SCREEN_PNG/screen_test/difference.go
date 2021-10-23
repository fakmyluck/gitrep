package main

import (
	"fmt"
	"image"
	"image/color"
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

	filename1, filename2 := "8_fail", "8"
	image1 := openimg(filename1)
	image2 := openimg(filename2)

	fmt.Printf("%T,%T", image1, image2)
	compareIMG(image1, image2)
	// save file
	outputFile, err := os.Create("difference.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	png.Encode(outputFile, image1)
	fmt.Println("File created.")

}

func compareIMG(imgOrigin, imgComp *image.RGBA) {

	if imgOrigin.Rect != imgComp.Rect {
		fmt.Printf("RECT %v!=%v\n", imgOrigin.Rect, imgComp.Rect)
		return
	}

	for x := 0; x < imgOrigin.Rect.Max.X; x++ {
		for y := 0; y < imgOrigin.Rect.Max.Y; y++ {
			if imgOrigin.At(x, y) != imgComp.At(x, y) {
				imgOrigin.Set(x, y, color.RGBA{255, 0, 0, 255})
				fmt.Printf("Diff found! (%v,%v)\n", x, y)
			}
		}
	}
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
	nrgba := img.(*image.RGBA)
	return nrgba

}
