package main

import (
	"fmt"
	"image"
	"os"
)
/*	b:=B.(*image.NRGBA)
	w:=W.(*image.NRGBA)
	wb:=wb.(*image.NRGBA)*/
func main() {
	B := returnimage("2_B")
	W := returnimage("2_W")
	WB := returnimage("2_WB")

	y:=0

	for x:=B.Bounds.Min.X ; x <B.Bounds.Max.X ; x++{
		r,g,b,_:=B.At(x,y)
		avg:=(r+g+b)/3

		fmt.Printf("B: %v %v %v \tavg: %v\n",r,g,b,avg)
		
		r,g,b,_=W.At(x,y)
		avg=(r+g+b)/3
		fmt.Printf("W: %v %v %v \tavg: %v\n",r,g,b,avg)
		
		r,g,b,_=WB.At(x,y)
		avg=(r+g+b)/3
		fmt.Printf("WB: %v %v %v \tavg: %v\n\n",r,g,b,avg)
		
	}
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
