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

type time struct {
	day   int
	mon   int
	year  int
	hour  int
	minut int
	sec   int
}

type dimensions struct {
	Dx, Dy int //dlinna shirina
}

func main() {

	var control picture
	//	Screenshot := takescreenshot("Screenshot")
	Screenshot := createpic("Screen", "888")
	sDy, sDx := Screenshot.dim.Dy, Screenshot.dim.Dx
	//var tX,tY int
	control = createpic("8", "8")

	for y := 0; y < sDy-control.dim.Dy; y++ {
		for x := 0; x < sDx-control.dim.Dx; x++ {

			if Screenshot.searchPic(x, y, control) { //zapisat' "word" esli naidena cifra

			}

		}

	}

	// fmt.Printf("myImg type: %T\n", Screenshot.img)
	// fmt.Println(word)

	outputFile, err := os.Create("obvedenniy_3.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	png.Encode(outputFile, Screenshot.img)
	fmt.Println("File created.")

}

func toTime(s string) time {
	T := time{
		sec:   loopstring(&s),
		minut: loopstring(&s),
		hour:  loopstring(&s),
		year:  loopstring(&s),
		mon:   loopstring(&s),
		day:   loopstring(&s),
	}
	return T
}

func loopstring(s *string) int {

	var sum int
	//fmt.Println(len((*s)))
	i := len(*s) - 1

	mult := 1
	for ; i >= 0 && (*s)[i] != ' ' && (*s)[i] != ':' && (*s)[i] != '.'; i-- {
		sum += (int((*s)[i]) - '0') * mult
		mult *= 10
	}
	if i >= 0 {
		*s = (*s)[:i]
	}

	return sum
}

func createpic(symbol, filename string) picture {
	var char picture
	char.sym = symbol

	RAWscr, err := os.Open("../pics/" + filename + ".png")
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

func (bigPic *picture) findvert(x, y int, num [12]picture) string {
	x0 := x
	var word string
	var tmp int

	bigPic.redcross(x, y)

	for xt := x; xt+6 < bigPic.dim.Dx && xt-x0 < 10; xt++ {
		for yt := y; yt < y+10; yt++ {

			for n := 0; n < 12; n++ {
				if bigPic.searchPic(xt, yt, num[n]) {

					if xt-x0-num[tmp].dim.Dx > 2 {
						word += " "
					}
					bigPic.greenbox(xt, yt, num[n])
					word += num[n].sym

					x0 = xt
					tmp = n
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

func (screenshot *picture) redcross(x, y int) {
	screenshot.ptr.SetRGBA(x, y, color.RGBA{255, 0, 0, 255}) //RED +
	screenshot.ptr.SetRGBA(x-1, y, color.RGBA{255, 0, 0, 255})
	screenshot.ptr.SetRGBA(x+1, y, color.RGBA{255, 0, 0, 255})
	screenshot.ptr.SetRGBA(x, y-1, color.RGBA{255, 0, 0, 255})
	screenshot.ptr.SetRGBA(x, y+1, color.RGBA{255, 0, 0, 255})
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
