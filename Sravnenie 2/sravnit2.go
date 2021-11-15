package main

import (
	"fmt"
	"image"
	"os"
)

func main() {
	B := returnimage("2_B")
	W := returnimage("2_W")
	WB := returnimage("2_WB")

	x,y:=0,0
	Bcol:=B.At(x,y)
	B.(*image.NRGBA).Pix.
	tst,_,_,_:=B.At(x,y).RGBA()
	
	fmt.Println("B:\t",B.At(x,y),"\t")
	fmt.Println("W:\t",W.At(x,y),"\t")
	fmt.Println("WB:\t",WB.At(x,y),"\t")
}

func returnimage(name string) image.Image {
	RAWscr, err := os.Open(name + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer RAWscr.Close()

	img, _, err = image.Decode(RAWscr)
	if err != nil {
		fmt.Println(err)
	}
	return img
}
