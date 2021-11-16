package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

/*	b:=B.(*image.NRGBA)
	w:=W.(*image.NRGBA)
	wb:=wb.(*image.NRGBA)*/
func main() {
	B := returnimage("2_B").(*image.NRGBA)
	W := returnimage("2_W").(*image.NRGBA)
	WB := returnimage("2_WB").(*image.NRGBA)

	//	y := 0
	for y := 0; y < B.Bounds().Dy(); y++ {
		for x := 0; x < B.Bounds().Dx(); x++ {

			r := B.Pix[x*4+y*B.Stride]
			g := B.Pix[x*4+1+y*B.Stride]
			b := B.Pix[x*4+2+y*B.Stride]
			//avg := r/3 + g/3 + b/3
			//avg1 := (uint16(r) + uint16(g) + uint16(b)) / 3
			//fmt.Printf("avg: %v %v\t", avg, 255-avg)
			fmt.Printf("%v %v %v   \t", 255-r, 255-g, 255-b)

			r = W.Pix[x*4+y*B.Stride]
			g = W.Pix[x*4+1+y*B.Stride]
			b = W.Pix[x*4+2+y*B.Stride]
			//avg = r/3 + g/3 + b/3
			//avg2 := (uint16(r) + uint16(g) + uint16(b)) / 3
			//fmt.Printf("%v\t", avg)
			//fmt.Printf("W:  %v %v %v    \tavg: %v\n", r, g, b, avg)
			fmt.Printf("%v %v %v   \t", r, g, b)

			r = WB.Pix[x*4+y*B.Stride]
			g = WB.Pix[x*4+1+y*B.Stride]
			b = WB.Pix[x*4+2+y*B.Stride]
			//avg = r/3 + g/3 + b/3
			//avg3 := (uint16(r) + uint16(g) + uint16(b)) / 3
			//fmt.Println(avg)
			//fmt.Printf("WB: %v %v %v    \tavg: %v\n\n", r, g, b, avg)
			//fmt.Printf("%v\t%v\t%v\n", avg3-(255-avg1), avg3-avg2, avg3)
			fmt.Printf("%v %v %v\n", r, g, b)
		}
		fmt.Println()
	}
}

func returnimage(name string) image.Image {
	RAWscr, err := os.Open(name + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer RAWscr.Close()

	img, _, err := image.Decode(RAWscr)
	if err != nil {
		fmt.Println(err)
	}
	return img
}
