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
	// WB := returnimage("2_WB").(*image.NRGBA)

	var dvumern [][]uint8
	//	y := 0
	var tmpstr byte
	for y := 0; y < B.Bounds().Dy(); y++ {
		for x := 0; x < B.Bounds().Dx(); x++ {
			tmp := (255 - B.Pix[x*4+y*B.Stride]) / 73
			tmpstr = ' '
			if tmp == 0 {
				tmpstr = 'o'
			} else if tmp == 1 {
				tmpstr = '-'
			} else if tmp == 2 {
				tmpstr = '.'
			}
			fmt.Print("", string(tmpstr))
		}
		fmt.Print("\t")
		for x := 0; x < B.Bounds().Dx(); x++ {
			tmp := (W.Pix[x*4+y*B.Stride] / 73)
			tmpstr = ' '
			if tmp == 0 {
				tmpstr = 'o'
			} else if tmp == 1 {
				tmpstr = '-'
			} else if tmp == 2 {
				tmpstr = '.'
			}
			fmt.Print("", string(tmpstr))
		}
		fmt.Print("\n")
	}
	// for y := 0; y < B.Bounds().Dy(); y++ {
	// 	for x := 0; x < B.Bounds().Dx(); x++ {

	// 		r := B.Pix[x*4+y*B.Stride]
	// 		g := B.Pix[x*4+1+y*B.Stride]
	// 		b := B.Pix[x*4+2+y*B.Stride]
	// 		//avg := r/3 + g/3 + b/3
	// 		//avg1 := (uint16(r) + uint16(g) + uint16(b)) / 3
	// 		//fmt.Printf("avg: %v %v\t", avg, 255-avg)
	// 		fmt.Printf("%v %v %v   \t", r, g, b)

	// 		nr := W.Pix[x*4+y*B.Stride]
	// 		ng := W.Pix[x*4+1+y*B.Stride]
	// 		nb := W.Pix[x*4+2+y*B.Stride]

	// 	proverka_nr:
	// 		for i := 0; i <= len(dvumern); i++ {
	// 			if i == len(dvumern) {
	// 				dvumern = append(dvumern, []uint8{nr})
	// 			}
	// 			if dvumern[i][0] == nr {
	// 				for n := 1; n <= len(dvumern[i]); n++ {
	// 					if n == len(dvumern[i]) {
	// 						dvumern[i] = append(dvumern[i], 255-r)
	// 						break proverka_nr
	// 					}
	// 					if 255-r == dvumern[i][n] {
	// 						break proverka_nr
	// 					}
	// 				}
	// 			}
	// 		}

	// 		//avg = r/3 + g/3 + b/3
	// 		//avg2 := (uint16(r) + uint16(g) + uint16(b)) / 3
	// 		//fmt.Printf("%v\t", avg)
	// 		//fmt.Printf("W:  %v %v %v    \tavg: %v\n", r, g, b, avg)
	// 		fmt.Printf("%v %v %v   \t", nr, ng, nb)

	// 		r = WB.Pix[x*4+y*B.Stride]
	// 		g = WB.Pix[x*4+1+y*B.Stride]
	// 		b = WB.Pix[x*4+2+y*B.Stride]
	// 		//avg = r/3 + g/3 + b/3
	// 		//avg3 := (uint16(r) + uint16(g) + uint16(b)) / 3
	// 		//fmt.Println(avg)
	// 		//fmt.Printf("WB: %v %v %v    \tavg: %v\n\n", r, g, b, avg)
	// 		//fmt.Printf("%v\t%v\t%v\n", avg3-(255-avg1), avg3-avg2, avg3)
	// 		fmt.Printf("%v %v %v\n", r, g, b)
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(dvumern)
	// fmt.Println("Pitaetsa sozdat' file")
	// outputFile, err := os.Create("B_r.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// e := png.Encode(outputFile, NRtoR(B))
	// if e != nil {
	// 	fmt.Println("FAILED to create, err=", e)
	// 	return
	// }

	// fmt.Println("Pitaetsa sozdat' file")
	// outputFile, err = os.Create("W_r.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// e = png.Encode(outputFile, NRtoR(W))
	// if e != nil {
	// 	fmt.Println("FAILED to create, err=", e)
	// 	return
	// }

	// fmt.Println("Pitaetsa sozdat' file")
	// outputFile, err = os.Create("WB_r.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// e = png.Encode(outputFile, NRtoR(WB))
	// if e != nil {
	// 	fmt.Println("FAILED to create, err=", e)
	// 	return
	// }

}

func NRtoR(img *image.NRGBA) image.Image {
	myImage := image.NewGray(image.Rect(img.Rect.Min.X, img.Rect.Min.Y, img.Rect.Max.X, img.Rect.Max.Y))

	for n := 0; n*4 < len(img.Pix); n++ {
		//fmt.Println("n: ", n)
		myImage.Pix[n] = img.Pix[n*4]
	}
	//fmt.Println("myImage.Stride: ", myImage.Stride)
	//myImage.Stride = img.Stride

	return myImage
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
