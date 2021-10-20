package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
)

type picture struct{
img image.Image   
ptr *image.RGBA
sDx,sDy,Dx,Dy int  //dlinna shirina
}

func main() {

num:=[12]picture

	RAWscr, err := os.Open("tstsmple.png")
	if err != nil {
		fmt.Println(err)
	}
	defer RAWscr.Close()

	MainScr, _, err := image.Decode(RAWscr)
	if err != nil {
		fmt.Println(err)
	}

	//var numPrime [12]image.Image
	//var num [12]*image.RGBA
	for n := '0'; n <= '9'; n++ {

		RAWscr, err = os.Open(string(n) + ".png")
		if err != nil {
			fmt.Println(err)
		}

		num.img[n], _, err = image.Decode(RAWscr)
		if err != nil {
			fmt.Println(err)
		}
		num.ptr[n] = num.img[n].(*image.RGBA)

	num[n].Dx, num[n].Dx := num[n].ptr.Bounds().Dy(), num[n].ptr.Bounds().Dx()
	}
	// switch num0.(type) {
	// case *image.RGBA:
	// 	fmt.Println("RGBA")
	// 	// i in an *image.RGBA
	// case *image.NRGBA:
	// 	fmt.Println("nRGBA")
	// 	// i in an *image.NRBGA
	// }
	//Scr.At(x+nx, y+ny)

	Scr := MainScr.(*image.RGBA)
	sDy, sDx := Scr.Rect.Dy(), Scr.Rect.Dx()



	for y := 0; y < sDy; y++ {
		if y >= sDy-num0.Bounds().Dy() {
			break
		}

		for x := 0; x < sDx; x++ {
/*
			if x >= sDx-num0.Bounds().Dx() {
				break
			}*/

for n:=0;n<12;n++{
			for ny := 0; ny < Dy; ny++ {
				for nx := 0; nx < Dx; nx++ {
					if num[n].ptr.At(nx, ny) != Scr.At(x+nx, y+ny) {
						//	fmt.Printf("failed at: %v, %v\n", x, y)
						goto skipSETRGBA
					}
				}
			}



			//fmt.Println("Success!")
			for zy := 0; zy < Dy; zy++ {
				Scr.SetRGBA(x+0, y+zy, color.RGBA{0, 255, 0, 255})
				Scr.SetRGBA(x+Dx-1, y+zy, color.RGBA{0, 255, 0, 255})
			}

			for zx := 0; zx < Dx; zx++ {
				Scr.SetRGBA(x+zx, y+0, color.RGBA{0, 255, 0, 255})
				Scr.SetRGBA(x+zx, y+Dy-1, color.RGBA{0, 255, 0, 255})
			}
		skipSETRGBA:
}//kovichka ot for
		}

	}

	// //	fmt.Println(MainScr.At(0, 0))
	// Scr.SetRGBA(0, 0, color.RGBA{0, 255, 0, 255})
	// Scr.SetRGBA(0, 1, color.RGBA{0, 255, 0, 255})
	// Scr.SetRGBA(0, 2, color.RGBA{0, 255, 0, 255})
	// Scr.SetRGBA(0, 3, color.RGBA{0, 255, 0, 255})
	// Scr.SetRGBA(1, 1, color.RGBA{0, 255, 0, 255})
	// Scr.SetRGBA(2, 1, color.RGBA{0, 255, 0, 255})
	// Scr.SetRGBA(3, 1, color.RGBA{0, 255, 0, 255})
	// Scr.SetRGBA(75, 75, color.RGBA{3, 3, 3, 255})
	// //	fmt.Println(MainScr.At(0, 0))

	// for y := 0; y < 10; y++ {

	// 	for zy := 0; zy < Dy; zy++ {
	// 		for zx := 0; zx < Dx; zx++ {
	// 			if num0.At(zx, zy) != Scr.At(zx, zy+y) {
	// 				fmt.Printf("F %v\t%v\n", num0.At(zx, zy), Scr.At(zx, zy+y))
	// 				goto skipSETRGBA2
	// 			}

	// 		}
	// 	}
	// 	fmt.Printf("SUCCESS!?")
	// 	//fmt.Printf("FOUND at: %v, %v\n", y, Dx)
	// 	for zy := 0; zy < Dy; zy++ {
	// 		Scr.SetRGBA(0, zy, color.RGBA{0, 255, 0, 255})
	// 		Scr.SetRGBA(Dy-1, zy, color.RGBA{0, 255, 0, 255})
	// 	}

	// 	for zx := 0; zx < Dx; zx++ {
	// 		Scr.SetRGBA(zx, 0, color.RGBA{0, 255, 0, 255})
	// 		Scr.SetRGBA(zx, Dy-1, color.RGBA{0, 255, 0, 255})
	// 	}

	// skipSETRGBA2:
	// }

	// //fmt.Println(MainScr)

	outputFile, err := os.Create("obvedenniy_3.png")
	if err != nil {
		fmt.Println(err)
	}

	png.Encode(outputFile, MainScr)
}
