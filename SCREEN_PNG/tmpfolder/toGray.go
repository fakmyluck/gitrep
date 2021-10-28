package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
)

func main() {
	openfile("0")
	fmt.Print("\n")
	openfile("1")
	fmt.Println()
	openfile("2")
	fmt.Print("\n")
	openfile("3")
	fmt.Print("\n")
	openfile("4")
	fmt.Print("\n")
	openfile("5")
	fmt.Print("\n")
	openfile("6")
	fmt.Print("\n")
	openfile("7")
	fmt.Print("\n")
	openfile("8")
	fmt.Print("\n")
	openfile("9")
	fmt.Print("\n")
	openfile("andres")
	fmt.Print("\n")
	openfile("dot")
	fmt.Print("\n")
	openfile("dot_dot")
	fmt.Print("\n")
	openfile("END")
	fmt.Print("\n")
	openfile("START")
	fmt.Print("\n")
	openfile("HZ")
	fmt.Print("\n")
}

func NRtoG(img *image.RGBA) image.Image {
	myImage := image.NewGray(image.Rect(img.Rect.Min.X, img.Rect.Min.Y, img.Rect.Max.X, img.Rect.Max.Y))

	for n := 0; n*4 < len(img.Pix); n++ {
		//fmt.Println("n: ", n)
		myImage.Pix[n] = img.Pix[n*4+1]
	}
	//fmt.Println("myImage.Stride: ", myImage.Stride)
	//myImage.Stride = img.Stride

	return myImage
}

func openfile(s string) {
	RAWscr, err := os.Open("../pics/" + s + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer RAWscr.Close()

	numPrime, _, err := image.Decode(RAWscr)
	if err != nil {
		fmt.Println(err)
	}

	// nrg := numPrime.(*image.NRGBA)

	//print_black(numPrime)
	fmt.Println("Pitaetsa sozdat' file")
	outputFile, err := os.Create("../pics/Gray/" + s + "_g.png")
	if err != nil {
		fmt.Println(err)
	}

	e := png.Encode(outputFile, NRtoG(numPrime.(*image.RGBA)))
	if e != nil {
		fmt.Println("FAILED to create "+s+"  ", err)
	}
	fmt.Println("Created: " + s + "_g.png")
	//png.Encode(outputFile, numzero)
}

func print_black(pic image.Image) {

	fmt.Println()
	for y := 0; y < pic.Bounds().Dy(); y++ {
		for x := 0; x < pic.Bounds().Dx(); x++ {

			tmpCol := pic.At(x, y).(color.RGBA)
			if tmpCol.G < 120 {

				fmt.Print("o")
			} else {
				fmt.Print(" ")

			}
			fmt.Print(" ")
		}
		fmt.Print("\n")

	}

}
