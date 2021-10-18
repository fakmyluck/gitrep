package main

import (
	"fmt"
	"image"
	_ "image/png" // Register JPEG format

	// Register PNG  format
	"log"
	"os"
)

func main() {

	RAWscr, err := os.Open("tstsmple.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer RAWscr.Close()

	num0, err := os.Open("0.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer num0.Close()

	_, frm, err := image.Decode(RAWscr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(frm)
	/*
		for y := 0; y < MainScr.Bounds().Dy(); y++ {
			for x := 0; x < MainScr.Bounds().Dx(); x++ {

				if x >= MainScr.Bounds().Dx()-num0.Bounds().Dx(){
					break
				}

				SearchNum(y,x,M)


			}

			if x >= MainScr.Bounds().Dy()-num0.Bounds().Dy(){
				break
			}

		}*/
	//fmt.Println(MainScr)
}

//func (p *)
