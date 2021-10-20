package main

import (
	"fmt"
	"image"
	"image/color"
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

	var num [12]picture
	var control [4]picture
	Screenshot, maxY, maxX := createpic("Screenshot", "tstsmple"), 0, 0
	sDy, sDx := Screenshot.dim.Dy, Screenshot.dim.Dx
	//var tX,tY int
	control[0] = createpic("HZ", "HZ")
	control[1] = createpic("Andres", "andres")
	control[2] = createpic("START", "START")
	control[3] = createpic("END", "END")

	var hz, start, end string
	var andres bool

	//var word string
	var c uint8 = 0

	for n, s := 0, ""; n <= 9; n++ {

		s = string('0' + n)
		num[n] = createpic(s, s)
		maxX, maxY = num[n].dim.maxpix(maxX, maxY)
	}

	num[10] = createpic(".", "dot")
	maxX, maxY = num[10].dim.maxpix(maxX, maxY)

	num[11] = createpic(":", "dot_dot")
	maxX, maxY = num[11].dim.maxpix(maxX, maxY)

	for y := 0; y < sDy-maxY; y++ {
		for x := 0; x < sDx-maxX; x++ {

			if Screenshot.searchPic(x, y, control[c]) { //zapisat' "word" esli naidena cifra
				switch c {
				case 0:
					{
						hz = Screenshot.findvert(x, y+12, num)
						y = y + 12
						x = x + 40
					}
				case 1:
					andres = true
					y = y + 256

				case 2:
					start = Screenshot.findvert(x+control[2].dim.Dx+30, y-3, num)

				case 3:
					end = Screenshot.findvert(x+control[3].dim.Dx+40, y-4, num)

				case 4:
					fmt.Println("ERROR case4!")
				}
				c++
				if c > 3 {
					goto printresults
				}
			}

		}

	}

printresults:

	fmt.Printf("Hz: %v\n", hz)
	fmt.Printf("andres: %v\n", andres)
	fmt.Printf("Start: %v\n", start)
	fmt.Printf("End: %v\n", end)

	//fmt.Println(word)

	// outputFile, err := os.Create("obvedenniy_3.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// png.Encode(outputFile, Screenshot.img)
	// fmt.Println("File created.")

}

func createpic(symbol, filename string) picture {
	var char picture
	char.sym = symbol

	RAWscr, err := os.Open(filename + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer RAWscr.Close()

	char.img, _, err = image.Decode(RAWscr)
	if err != nil {
		fmt.Println(err)
	}

	char.ptr = char.img.(*image.RGBA)
	char.dim.Dy, char.dim.Dx = char.ptr.Bounds().Dy(), char.ptr.Bounds().Dx()
	return char
}

func (bigPic picture) findvert(x, y int, num [12]picture) string {
	x0 := x
	var word string

	for xt := x; xt+6 < bigPic.dim.Dx && xt-x0 < 10; xt++ {
		for yt := y; yt < y+10; yt++ {

			for n := 0; n < 12; n++ {
				if bigPic.searchPic(xt, yt, num[n]) {
					word += num[n].sym
					x0 = xt
				}
			}

		}
	}
	return word
}

func (bigPic picture) findhor(x, y int, num [12]picture) string {
	x0 := x
	var word string

	for yt := y; yt < y+11; yt++ {
		for xt := x; xt+6 < bigPic.dim.Dx && xt-x0 < 8; xt++ {

			for n := 0; n < 12; n++ {
				if bigPic.searchPic(xt, yt, num[n]) {
					word += num[n].sym
					x0 = xt
				}
			}

		}
	}
	return word
}

func (bigPic *picture) searchPic(x, y int, obj picture) bool {

	if bigPic.dim.Dx-x < obj.dim.Dx || bigPic.dim.Dy-y < obj.dim.Dy {
		return false
	}
	for Ny := 0; Ny < obj.dim.Dy; Ny++ {
		for Nx := 0; Nx < obj.dim.Dx; Nx++ {
			if obj.ptr.At(Nx, Ny) != bigPic.ptr.At(x+Nx, y+Ny) {

				return false //goto skipSETRGBA
			}
		}
	}
	bigPic.greenbox(x, y, obj)
	return true
}

func (screenshot *picture) greenbox(x, y int, num picture) {
	for zy := 0; zy < num.dim.Dy; zy++ {
		screenshot.ptr.SetRGBA(x+0, y+zy, color.RGBA{0, 255, 0, 255})
		screenshot.ptr.SetRGBA(x+num.dim.Dx-1, y+zy, color.RGBA{0, 255, 0, 255})
	}

	for zx := 0; zx < num.dim.Dx; zx++ {
		screenshot.ptr.SetRGBA(x+zx, y+0, color.RGBA{0, 255, 0, 255})
		screenshot.ptr.SetRGBA(x+zx, y+num.dim.Dy-1, color.RGBA{0, 255, 0, 255})

	}
}

func (leng dimensions) maxpix(MaxX, MaxY int) (int, int) {
	if leng.Dx > int(MaxX) {
		MaxX = leng.Dx
	}
	if leng.Dy > int(MaxY) {
		MaxY = leng.Dy
	}
	return MaxX, MaxY
}
