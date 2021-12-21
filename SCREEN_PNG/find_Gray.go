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
	"time"
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

type controll struct {
	img       grayP
	ots, nach coords
}
type controllArray struct {
	hz     controll
	Andres controll
	Eduard controll
	dezh   controll
	start  controll
	end    controll
	num    [12]grayP
	max    maximum
}

// func savetofile(filename string, a, e todisplay) {//
	
	// }
	
func main() {
		
	cnt := createcontroll()
	
	var disE, disA []todisplay
	var key_input string
	var enable_screenshot bool = false
	var timer_enabled bool = false
	var tmpdisp todisplay
	var tStart, tEnd time.Time
	
	fmt.Println("timer_enabled", timer_enabled)
	
	fmt.Println("[help] -помощ")
	for p := 0; p < 15; p++ {
		
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
				if timer_enabled {
					tStart = time.Now()
				}
				break input_switch
			case "h", "help", "Help", "HELP", "рудз", "р", "[help]":
				fmt.Println("\n\n\t1. Открыть новый наряд")
				fmt.Println("\t2. Заполнить")
				fmt.Println("\t3. СОХРОНИТЬ наряд [F12]")
				fmt.Println("\t4. В данной програме написать [screen] или [s]")
				fmt.Println("\t5. перейти к пункту 1.\n")
				fmt.Println("[screen] или [s] -добавить наряд (наряд должен быть на экране от [НЗ:] до [Факт. заверш.:]")
				fmt.Println("[reset] или [r] -отчистить список          (ra)")
				fmt.Println("[exit] или [e] -выход")
				fmt.Println("[d] debug time")
				fmt.Println("[es] [ds] -enable/disable screenshot")
				fmt.Println("[input] [i] input raw data")
				fmt.Println("[dela] [dele] delete Andres/Eduard")
			case "enablescreenshot", "enables", "es":
				enable_screenshot = true
				fmt.Println("Screenshot enabled.")
			case "disablescreenshot", "disables", "ds":
				enable_screenshot = false
				fmt.Println("Screenshot disabled.")
			case "debug", "DEBUG", "Debug", "d", "D", "в", "В", "time", "T", "t", "е", "Е":
				if timer_enabled {
					timer_enabled = false
					fmt.Println("Timer disabled.")
					} else {
						timer_enabled = true
						fmt.Println("Timer enabled.")
					}
			case "input","i","Input":{//ruchnoy vvod
				timer_enabled = false
				tmpdisp=inputdate()
				goto skip_sceenshot
			}
			case "delA","dela","вудф":{
				if disA!=nil{
					deletedisp(disA)
				}else{
					fmt.Println("disA == nil")
				}
				goto displayloop:
					
			}
			case "delE","dele","вуду":{
				if disE!=nil{
					deletedisp(disE)
				}else{
					fmt.Println("disE == nil")
				}
				goto displayloop:	
			}
			default:
				fmt.Println("! Неверный ввод.")
			}
				key_input = ""
			}
			
			Screenshot := takescreenshot("Screenshot")
			
			//FUNC podshet
			tmpdisp = rawtodisplay(Screenshot.scanit(cnt))
			skip_sceenshot:

			if tmpdisp.name == 'A' {
				removeold(tmpdisp, &disA)
				checkfordupes(tmpdisp, &disA)
			} else if tmpdisp.name == 'E' {
				removeold(tmpdisp, &disE)
				checkfordupes(tmpdisp, &disE)
			}
			goto displayloop:
				displayloop(disA, disE, tmpdisp.name)
				if enable_screenshot {
					Screenshot.printscreen("obvedenniy_3.png")
				}
				
				if timer_enabled {
					tEnd = time.Now()
					fmt.Println("Elapse time:", tEnd.Sub(tStart))
				}
	}
				
}
			
func deletedisp(disp []todisplay) []todisplay{
	fmt.Print("Index to remove?(net proverki na 0> & <10 & nil): ")
	fmt.Scan(&key_input)	
	vvedenniyind:=byte(key_input[0]-48)	//sdelat' proverku na "nil?>>?>???"

	last=len(disp)-1
	if last > vvedenniyind-1{
		disp[vvedenniyind-1]=disp[last]
	}
	disp=disp[:last-1]
	
	return disp
}

func createcontroll() (cnt controllArray) {
cnt.hz = controll{img: createpic("HZ", "HZ"), ots: coords{Y: 14}}
cnt.Andres = controll{img: createpic("Andres", "andres")}
cnt.Eduard = controll{img: createpic("Eduard", "eduard")}
cnt.dezh = controll{img: createpic("D", "dezh")}
cnt.start = controll{img: createpic("START", "START"), ots: coords{Y: -3}}
cnt.start.ots.X = int16(cnt.Andres.img.dim.Dx) + 50
cnt.end = controll{img: createpic("END", "END"), ots: coords{Y: -4}}
cnt.end.ots.X = int16(cnt.Andres.img.dim.Dx) + 68

for n, s := 0, ""; n <= 9; n++ {
s = string('0' + n)
cnt.num[n] = createpic(s, s)
cnt.num[n].dim.maxpix(&cnt.max)
}
cnt.num[10] = createpic(".", "dot")
cnt.num[10].dim.maxpix(&cnt.max)
cnt.num[11] = createpic(":", "dot_dot")
cnt.num[11].dim.maxpix(&cnt.max)

return
}

func (Screenshot *picture) scanit(cnt controllArray) (scan rawstr) {
scan.hz = Screenshot.findvert(Screenshot.findCoords(cnt.hz), cnt.num)
if scan.hz == "" {
fmt.Println("ERROR scan.hz (убрать выделение с номера наряда)")
}
if Screenshot.findCoords(cnt.Andres) != (coords{-1, -1}) {
scan.name = cnt.Andres.img.sym
} else if Screenshot.findCoords(cnt.Eduard) != (coords{-1, -1}) {
scan.name = cnt.Eduard.img.sym
} else {
fmt.Println("не нашёл имени")
}
if Screenshot.findCoords(cnt.dezh) != (coords{-1, -1}) {
scan.dezh = cnt.dezh.img.sym
}
scan.start = Screenshot.findvert(Screenshot.findCoords(cnt.start), cnt.num)
scan.end = Screenshot.findvert(Screenshot.findCoords(cnt.end), cnt.num)
return
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

func inputdate() todisplay{
var disp todisplay
var scankey string
fmt.Print("Введите НЗ: ")
fmt.Scan(&scankey)
disp.HZ=scankey

fmt.Print("Име (А/Е): ")
fmt.Scan(&scankey)
if scankey=="A"{
disp.name='A'
}else if scankey=="E"{
disp.name='E'
}else {
fmt.Println("ERR input name(disp.name= default ['A'])")
}

fmt.Print("Наряд дежурство? (Y/N?): ")
fmt.Scan(&scankey)
if scankey==Y{
disp.dezh="D"
}
disp.dezh=scankey

fmt.Print("Начало: ")
fmt.Scan(&scankey)
disp.start=tim.ToTime(scankey)

fmt.Print("Введите НЗ: ")
fmt.Scan(&scankey)
disp.end=tim.ToTime(scankey)


disp.start.Units = tim.TimeToUnits(disp.start)
disp.end.Units = tim.TimeToUnits(disp.end)

disp.sub = disp.end.Units - disp.start.Units 

return disp
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
	last := len(*disp) - 1
	for i := last; i >= 0; i-- {
		dT := tmpdisp.start.Units - (*disp)[i].start.Units
		if dT > 604_800 {
			fmt.Printf("Разница во времени более недели(%v дня),\nстарые наряды не удалены, ([R] чтобы стереть всё)", dT/86_400)
			} else if dT > 50000 {
				
				if i == last {
					if last == 0 {
						*disp = nil
						return
					}
					last--
					} else {
						(*disp)[i] = (*disp)[last]
						last--
					}
				}
				
			}
			
			if last != len(*disp)-1 {
				(*disp) = (*disp)[0 : last+1]
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
						// func strtoi16(s string) int16 {
							// 	var sum int16
							// 	i := len(s) - 1
							// 	var mult int16 = 1
							// 	for ; i >= 0 && (s)[i] != ' ' && (s)[i] != ':' && (s)[i] != '.'; i-- {
								// 		sum += (int16((s)[i]) - '0') * mult
								// 		mult *= 10
								// 	}
								// 	return sum
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
								
								RAWscr, err := os.Open("pics/Red/" + filename + ".png")
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
													//break
												}
												bigPic.greenbox(xt, yt, num[n])
												word += num[n].sym
												
												x0 = xt
												tmp = n
											}
										}
										
									}
								}
								
								if word != "" {
									return word
									
								}
								
								// Probivat' Siniy
								for xt := x; xt+6 < bigPic.dim.Dx && xt-x0 < 10; xt++ {
									for yt := y; yt < y+10; yt++ {
										
										for n := 0; n < 12; n++ {
											
											if bigPic.searchBluePic(xt, yt, num[n]) {
												
												if xt-x0-num[tmp].dim.Dx > 2 {
													word += " "
													y = yt
													break
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
							
							func (bigPic *picture) searchBluePic(x, y int, obj grayP) bool {
								var e byte
								
								if bigPic.dim.Dx-x < obj.dim.Dx || bigPic.dim.Dy-y < obj.dim.Dy {
									return false
								}
								e = 0
								for Ny := 0; Ny < obj.dim.Dy; Ny++ {
									for Nx := 0; Nx < obj.dim.Dx; Nx++ {
										R := (255 - bigPic.ptr.RGBAAt(x+Nx, y+Ny).R) / 110
										r := obj.ptr.GrayAt(Nx, Ny).Y / 110
										
										if r != R {
											e++
											if e > 2 {
												return false //goto skipSETRGBA
											}
										}
									}
								}
								
								if e != 0 {
									fmt.Printf("searchPic: ошибок в (%v) символе: %v\n", obj.sym, e)
								}
								bigPic.greenbox(x, y, obj)
								return true
							}
							
							func (bigPic *picture) searchPic(x, y int, obj grayP) bool {
								var e byte
								
								if bigPic.dim.Dx-x < obj.dim.Dx || bigPic.dim.Dy-y < obj.dim.Dy {
									return false
								}
								e = 0
								for Ny := 0; Ny < obj.dim.Dy; Ny++ {
									for Nx := 0; Nx < obj.dim.Dx; Nx++ {
										R := bigPic.ptr.RGBAAt(x+Nx, y+Ny).R
										r := obj.ptr.GrayAt(Nx, Ny).Y
										
										if r != R {
											e++
											if e > 2 {
												return false //goto skipSETRGBA
											}
										}
									}
								}
								
								if e != 0 {
									fmt.Printf("searchPic: ошибок в (%v) символе: %v\n", obj.sym, e)
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
							