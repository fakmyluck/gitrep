package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
	"slucker/SCREEN_PNG/screenshot"
	"slucker/SCREEN_PNG/tim"
)

type picture struct {
	img image.Image
	ptr *image.RGBA
	dim dimensions
	sym string
}

type grayP struct {
	img image.Image
	ptr *image.Gray
	dim dimensions
	sym string
}

type maximum struct {
	X, Y int
}

type dimensions struct {
	Dx, Dy int //dlinna shirina
}

type controll struct {
	img       grayP
	ots, nach coords
}

type coords struct {
	X, Y int16
}
type rawstr struct {
	hz, name, dezh, start, end string
}

type todisplay struct {
	HZ    uint32
	name  byte
	dezh  byte
	start tim.Time
	end   tim.Time
	sub   int32
}

// func savetofile(filename string, a, e todisplay) {

// }

func main() {

	var num [12]grayP

	//sDy, sDx := Screenshot.dim.Dy, Screenshot.dim.Dx
	hz := controll{img: createpic("HZ", "HZ"), ots: coords{Y: 14}}
	Andres := controll{img: createpic("Andres", "andres")}
	Eduard := controll{img: createpic("Eduard", "eduard")}
	dezh := controll{img: createpic("D", "dezh")}
	start := controll{img: createpic("START", "START"), ots: coords{Y: -3}}
	start.ots.X = int16(Andres.img.dim.Dx) + 50
	end := controll{img: createpic("END", "END"), ots: coords{Y: -4}}
	end.ots.X = int16(Andres.img.dim.Dx) + 68

	var max maximum
	var scan rawstr

	for n, s := 0, ""; n <= 9; n++ {
		s = string('0' + n)
		num[n] = createpic(s, s)
		num[n].dim.maxpix(&max)
	}

	num[10] = createpic(".", "dot")
	num[10].dim.maxpix(&max)

	num[11] = createpic(":", "dot_dot")
	num[11].dim.maxpix(&max)

	var disE, disA []todisplay
	var key_input string
	var enable_screenshot bool = false

	for p := 0; p < 15; p++ {
		fmt.Println("[help] -помощ")
	input_switch:
		for {

			fmt.Scan(&key_input)
			switch key_input {
			case "e", "E", "Exit", "EXIT", "exit", "у", "учше", "[exit]":
				return
			case "reset", "RESET", "R", "r", "Reset", "к", "куыуе", "[reset]":
				disE = nil
				fmt.Println("\nСписок Э. отчищен.")
			case "reseta", "RESETA", "RA", "ra", "Reseta", "куыуеф", "кф":
				disA = nil
				fmt.Println("\nСписок А. отчищен.")
			case "", "s", "S", "screen", "[s]", "ы":
				break input_switch
			case "h", "help", "Help", "HELP", "рудз", "р", "[help]":
				fmt.Println("\n\n[screen] или [s] -добавить наряд (наряд должен быть на экране от [НЗ:] до [Факт. заверш.:]")
				fmt.Println("[reset] или [r] -отчистить список		(ra)")
				fmt.Println("[exit] или [e] -выход")
				fmt.Println("[es] [ds] -enable/disable screenshot")
			case "enablescreenshot", "enables", "es":
				enable_screenshot = true
				fmt.Println("Screenshot enabled.")
			case "disablescreenshot", "disables", "ds":
				enable_screenshot = false
				fmt.Println("Screenshot disabled.")
			default:
				fmt.Println("! Неверный ввод.")
			}
			key_input = ""
		}

		Screenshot := takescreenshot("Screenshot")
		scan.hz = Screenshot.findvert(Screenshot.findCoords(hz), num)
		if scan.hz == "" {
			fmt.Println("ERROR scan.hz (убрать выделение с номера наряда)") ///gfdgdfg pererabotka errora DODEAT' <<<<<<<<<<<<<<
		}

		if Screenshot.findCoords(Andres) != (coords{-1, -1}) {
			scan.name = Andres.img.sym
		} else if Screenshot.findCoords(Eduard) != (coords{-1, -1}) {
			scan.name = Eduard.img.sym
		} else {
			fmt.Println("не нашёл имени")
		}

		if Screenshot.findCoords(dezh) != (coords{-1, -1}) {
			scan.dezh = dezh.img.sym
		}

		scan.start = Screenshot.findvert(Screenshot.findCoords(start), num)
		scan.end = Screenshot.findvert(Screenshot.findCoords(end), num)

		//FUNC podshet
		tmpdisp := rawtodisplay(scan)

		if tmpdisp.name == 'A' {
			removeold(tmpdisp, &disA)
			checkfordupes(tmpdisp, &disA)
		} else if tmpdisp.name == 'E' {
			removeold(tmpdisp, &disE)
			checkfordupes(tmpdisp, &disE)
		}

		displayloop(disA, disE, tmpdisp.name)
		if enable_screenshot {
			Screenshot.printscreen("obvedenniy_3.png")
		}
	}
}

func displayforloop(disp []todisplay) {
	for n := 0; n < len(disp); n++ {
		fmt.Printf("%v, %v.%02d %02d:%02d - %v.%02d %02d:%02d [%.2f]", disp[n].HZ /**/, disp[n].start.Day, disp[n].start.Mon, disp[n].start.Hour, disp[n].start.Min /**/, disp[n].end.Day, disp[n].start.Mon, disp[n].end.Hour, disp[n].end.Min /**/, float32(disp[n].sub)/3600)
		if disp[n].dezh == 'D' {
			fmt.Printf(" (%.2f)", countdezh(disp)/3600)
		}
		fmt.Println()
	}
}

func countdezh(disp []todisplay) float32 {
	var dezh int32
	var sum int32
	for i := 0; i < len(disp); i++ {
		if (disp)[i].dezh != 0 {
			dezh = (disp)[i].sub
		} else {
			sum += (disp)[i].sub
		}
	}
	return float32(dezh - sum)
}

func display(disp todisplay) {
	fmt.Printf("%v, %v.%02d %02d:%02d - %v.%02d %02d:%02d [%.2f]\n", disp.HZ /**/, disp.start.Day, disp.start.Mon, disp.start.Hour, disp.start.Min /**/, disp.end.Day, disp.start.Mon, disp.end.Hour, disp.end.Min /**/, float32(disp.sub)/3600)
}

func displayloop(disAndres []todisplay, disEduard []todisplay, name byte) {
	if name == 'A' {
		if disEduard != nil {
			fmt.Println("Eduard:")
			displayforloop(disEduard)
		}
		fmt.Println("\nAndres:")
		displayforloop(disAndres)
	} else if name == 'E' {
		if disAndres != nil {
			fmt.Println("Andres:")
			displayforloop(disAndres)
		}
		fmt.Println("\nEduard:")
		displayforloop(disEduard)

	} else {
		fmt.Printf("tmp.name [%v] failed!\n", name)
	}
}

func checkfordupes(tmpdisp todisplay, disp *[]todisplay) {
	if *disp == nil {
		*disp = append((*disp), tmpdisp)
		return
	}
	for i := 0; i < len(*disp); i++ {
		if tmpdisp.HZ == (*disp)[i].HZ {
			if tmpdisp.sub != (*disp)[i].sub {
				(*disp)[i] = tmpdisp
			}
			return
		}
	}
	*disp = append((*disp), tmpdisp)
}

func removeold(tmpdisp todisplay, disp *[]todisplay) {
	if *disp == nil {
		return
	}
	n := len(*disp) - 1
	for i := n; i >= 0; i-- {
		if tmpdisp.start.Units-(*disp)[i].start.Units > 50000 {
			if n == i {
				if n == 0 {
					(*disp)[i] = tmpdisp
					return
				}
				n--
			} else {
				(*disp)[i] = (*disp)[n]
				n--
			}
		}
	}
	if n != len(*disp)-1 {
		(*disp) = (*disp)[0:n]
	}
}

// func disptobyte(a todisplay) []byte {
// 	var bytearray []byte //4 1 1 (2*6)*2 2 -> 32
// 	bytearray = append(bytearray, convert.Uint32tobyte(a.HZ)...)
// 	bytearray = append(bytearray, a.name)
// 	bytearray = append(bytearray, a.dezh)
// 	bytearray = append(bytearray, rettimeinbytes(a.start)...)
// 	bytearray = append(bytearray, rettimeinbytes(a.end)...)
// 	bytearray = append(bytearray, convert.Uint16tobyte(a.sub)...)
// 	return bytearray
// }

// func rettimeinbytes(a tim.Time) []byte {
// 	var bytearray []byte
// 	bytearray = append(bytearray, convert.Uint16tobyte(a.Day)...)
// 	bytearray = append(bytearray, convert.Uint16tobyte(a.Mon)...)
// 	bytearray = append(bytearray, convert.Uint16tobyte(a.Year)...)
// 	bytearray = append(bytearray, convert.Uint16tobyte(a.Hour)...)
// 	bytearray = append(bytearray, convert.Uint16tobyte(a.Min)...)
// 	return append(bytearray, convert.Uint16tobyte(a.Sec)...)
// }

// // func frombytetodisplay(bytearray []byte) todisplay {
// // 	var a todisplay
// // 	a.HZ = convert.Bytetouint32(bytearray[0:4])
// // 	a.name = bytearray[4]
// // 	a.dezh = bytearray[5]
// // 	a.start = frombytetotime(bytearray[6:18])
// // 	a.end = frombytetotime(bytearray[18:30])
// // 	a.sub = convert.Bytetouint16(bytearray[30:32])
// // 	return a
// // }

// func frombytetotime(bytearray []byte) tim.Time {
// 	var a tim.Time
// 	a.Day = convert.Bytetouint16(bytearray[0:2])
// 	a.Mon = convert.Bytetouint16(bytearray[2:4])
// 	a.Year = convert.Bytetouint16(bytearray[4:6])
// 	a.Hour = convert.Bytetouint16(bytearray[6:8])
// 	a.Min = convert.Bytetouint16(bytearray[8:10])
// 	a.Sec = convert.Bytetouint16(bytearray[10:12])
// 	return a
// }

func rawtodisplay(scan rawstr) (disp todisplay) {
	disp.HZ = loopstring(scan.hz)
	switch scan.name {
	case "Andres":
		disp.name = 'A'
	case "Eduard":
		disp.name = 'E'
	}
	if scan.dezh != "" {
		disp.dezh = 'D'
	}
	disp.start = tim.ToTime(scan.start)
	disp.end = tim.ToTime(scan.end)

	disp.start.Units = tim.TimeToUnits(disp.start)
	disp.end.Units = tim.TimeToUnits(disp.end)

	disp.sub = disp.end.Units - disp.start.Units //tim.Subtime(disp.start, disp.end)
	return
}

func loopstring(s string) (sum uint32) {
	var mult uint32 = 1
	for i := len(s) - 1; i >= 0; i-- {
		sum += (uint32((s)[i]) - '0') * mult
		mult *= 10
	}
	return
}

func (screenshot picture) findCoords(poiskS controll) coords {
	for y := 0; y < screenshot.dim.Dy-poiskS.img.dim.Dy; y++ { //max.Y eto maximum visoti kartinki (sDy-max.Y)
		for x := 0; x < screenshot.dim.Dx-poiskS.img.dim.Dx; x++ {

			if screenshot.searchPic(x, y, poiskS.img) { //zapisat' "word" esli naidena cifra
				return coords{int16(x) + poiskS.ots.X, int16(y) + poiskS.ots.Y}
			}
		}
	}
	return coords{-1, -1}
}

func takescreenshot(symbol string) picture {

	var char picture
	char.sym = symbol

	var err error
	char.ptr, err = screenshot.CaptureScreen()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	char.img = image.Image(char.ptr)

	char.dim.Dy, char.dim.Dx = char.ptr.Bounds().Dy(), char.ptr.Bounds().Dx()
	return char
}

func createpic(symbol, filename string) grayP {
	var char grayP
	char.sym = symbol

	RAWscr, err := os.Open("pics/Gray/" + filename + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer RAWscr.Close()

	char.img, _, err = image.Decode(RAWscr)
	if err != nil {
		fmt.Println(err)
	}

	char.ptr = char.img.(*image.Gray)
	char.dim.Dy, char.dim.Dx = char.ptr.Bounds().Dy(), char.ptr.Bounds().Dx()
	return char
}

func (bigPic *picture) findvert(cords coords, num [12]grayP) string {
	if cords.X < 0 {
		return ""
	}
	x, y := int(cords.X), int(cords.Y)
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

func strtoi16(s string) int16 {
	var sum int16
	i := len(s) - 1
	var mult int16 = 1
	for ; i >= 0 && (s)[i] != ' ' && (s)[i] != ':' && (s)[i] != '.'; i-- {
		sum += (int16((s)[i]) - '0') * mult
		mult *= 10
	}
	return sum
}

func (bigPic *picture) searchPic(x, y int, obj grayP) bool {
	var e byte

	if bigPic.dim.Dx-x < obj.dim.Dx || bigPic.dim.Dy-y < obj.dim.Dy {
		return false
	}
	for Ny := 0; Ny < obj.dim.Dy; Ny++ {
		for Nx := 0; Nx < obj.dim.Dx; Nx++ {

			_, G, _, _ := bigPic.ptr.At(x+Nx, y+Ny).RGBA()
			_, g, _, _ := obj.ptr.At(Nx, Ny).RGBA()
			if g != G {
				e++
				if e > 1 {
					return false //goto skipSETRGBA
				}
			}
		}
	}

	if e == 1 {
		fmt.Println("searchPic: одна ошибка в символе (е=", e, ")")
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

func (screenshot *picture) greenbox(x, y int, num grayP) {
	for zy := 0; zy < num.dim.Dy; zy++ {
		screenshot.ptr.SetRGBA(x+0, y+zy, color.RGBA{0, 255, 0, 255})
		screenshot.ptr.SetRGBA(x+num.dim.Dx-1, y+zy, color.RGBA{0, 255, 0, 255})
	}

	for zx := 0; zx < num.dim.Dx; zx++ {
		screenshot.ptr.SetRGBA(x+zx, y+0, color.RGBA{0, 255, 0, 255})
		screenshot.ptr.SetRGBA(x+zx, y+num.dim.Dy-1, color.RGBA{0, 255, 0, 255})

	}
}

func (leng dimensions) maxpix(Max *maximum) {
	if leng.Dx > int(Max.X) {
		Max.X = leng.Dx
	}
	if leng.Dy > int(Max.Y) {
		Max.Y = leng.Dy
	}
}

func (screenshot picture) printscreen(filename string) {
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	png.Encode(outputFile, screenshot.img)
	fmt.Printf("File [%v] created.\n", filename)
}
