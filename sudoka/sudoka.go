package main

import "fmt"

var dbg [2]int

func PrintGrid(grid [9][9][11]int) {

	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	for y, grid_row := range grid {
		if y%3 == 0 && y != 0 {
			fmt.Print(" -----------------------\n")
		}
		for x, grid_column := range grid_row {
			if x%3 == 0 && x != 0 {
				fmt.Print(" |")
			}
			fmt.Print(colorGreen) //green
			if grid_column[9] == 9 {
				grid_column[9] = '*' - '0'
				fmt.Print(colorCyan) //cyan
			}
			if grid_column[10] == 1 {
				fmt.Print(colorWhite)
			}
			if grid_column[10] == 2 {
				fmt.Print(colorPurple) //purp
			}

			if grid_column[10] == 3 { // ---
				fmt.Print(colorYellow)
				//	grid[y][x][10] = 2
			}

			if grid_column[10] == 4 { //	|
				fmt.Print(colorBlue)
				//	grid[y][x][10] = 2
			}

			if grid_column[10] == 5 { //	|
				fmt.Print(colorRed)
				//	grid[y][x][10] = 2
			}

			fmt.Printf(" %v", string(grid_column[9]+1+'0'))

			fmt.Print(colorReset)
		}

		fmt.Println()
	}

	fmt.Println()
}

func MinusOne(grid [9][9]int) [9][9]int {
	for y, row := range grid {
		for x, cell := range row {
			if 0 < cell && cell < 10 {
				grid[y][x] = cell - 1
			} else {
				grid[y][x] = 9
			}
		}
	}
	return grid
}

func Returnifone(posibl [11]int) int { // nahodit' edinstvenniy "1"
	sum := 0
	n := 9
	for i, num := range posibl[:9] {
		sum += num
		if num == 1 {
			n = i
		}
	}

	if sum != 1 {
		n = 9
	}

	return n
}

func GridInit(unsolved [9][9]int) [9][9][11]int {
	var grid [9][9][11]int

	dbg[0] = dbg[1]
	dbg[1] = 3

	for Y := 0; Y < 9; Y++ {
		for X := 0; X < 9; X++ {
			if unsolved[Y][X] != 9 {
				grid[Y][X][9] = unsolved[Y][X]
			} else {
				grid[Y][X][9] = 9
				for i := 0; i < 9; i++ {
					grid[Y][X][i] = 1
				}
			}
		}
	}
	return grid
}

func ErrorCheck(grid [9][9][11]int, e *bool) {
	var sumX, sumY [9]int

	for Y := 0; Y < 9; Y++ {
		for X := 0; X < 9; X++ {

			if grid[Y][X][9] != 9 {
				sumX[grid[Y][X][9]] += 1

			}
			if grid[X][Y][9] != 9 {
				sumY[grid[X][Y][9]] += 1

			}
		}

		for i := 0; i < 9; i++ {
			if sumX[i] > 1 || sumY[i] > 1 { // proverit' CHE S ETIM??
				fmt.Printf("\nERROR dupe+! i=%v  sumX=%v  sumY=%v\n", i, sumX[i], sumY[i])
				fmt.Println("Prev func: ", dbg)
				*e = true
			}
		}

		for i := 0; i < 9; i++ {
			sumX[i] = 0
			sumY[i] = 0
		}
	}

	n := 0
	for Z := 0; Z < 9; Z += 3 {
		for C := 0; C < 9; C += 3 {
			n = 0
			for zi := Z; zi < 3; zi++ {
				for ci := C; ci < 3; ci++ {
					if grid[Z][C][9] != 9 {
						sumX[n] += grid[Z][C][9]
					}
					n++
				}
			}

			for i := 0; i < 9; i++ {

				if sumX[i] > 1 {
					fmt.Printf("\nERROR dupe []! i=%v  sumX=%v \n", i, sumX[i])
					fmt.Println("Prev func: ", dbg)
					*e = true
				}
				sumX[i] = 0
			}
		}
	}
}

func LineCheck(grid [9][9][11]int) [9][9][11]int { //1
	var tmp int

	dbg[0] = dbg[1] // <--- eto nomer prigrammi dlya ErrorCheck, dodelat' & *
	dbg[1] = 1

	for Y := 0; Y < 9; Y++ {
		for X := 0; X < 9; X++ {

			if grid[Y][X][9] != 9 {
				tmp = grid[Y][X][9]
				for i := 0; i < 9; i++ {
					grid[i][X][tmp] = 0
					grid[Y][i][tmp] = 0

				}
				//grid [Y][X][tmp]=1
			}
		}
	}
	return SquareCheck(grid)
}

func SquareCheck(grid [9][9][11]int) [9][9][11]int { //2
	var snum [9]int

	dbg[0] = dbg[1]
	dbg[1] = 2

	for baseY := 0; baseY < 9; baseY += 3 {
		for baseX := 0; baseX < 9; baseX += 3 {

			for i := 0; i < 9; i++ { //onulenie
				snum[i] = 1
			}

			for c := 0; c < 2; c++ {
				for Y := baseY; Y < baseY/3*3+3; Y++ {
					for X := baseX; X < baseX/3*3+3; X++ {
						if c == 0 && grid[Y][X][9] != 9 {
							snum[grid[Y][X][9]] = 0
						} else if c == 1 {
							for i := 0; i < 9; i++ {
								grid[Y][X][i] *= snum[i]
							}

						}
					}
				}
			}
		}
	}
	return grid
}

/*
func PrintOnes(grid [9][9][11]int, a int) {

	for iY := 0; iY < 9; iY++ {
		for iX := 0; iX < 9; iX++ {
			if grid[iY][iX][9] == a {
				fmt.Print(a+1, " ")
			} else if grid[iY][iX][9] != 9 {
				fmt.Print("o ")
			} else {
				fmt.Print(grid[iY][iX][a], " ")
			}
			if (iX+1)%3 == 0 {
				fmt.Print(" ")
			}
		}
		if (iY+1)%3 == 0 {
			fmt.Print("\n")
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println(a + 1)
}*/

func SetCell(grid *[9][9][11]int, Y int, X int) {

	dbg[0] = dbg[1]
	dbg[1] = 4

	SetThisNumZero := grid[Y][X][9]
	//PrintOnes(*grid, SetThisNumZero)

	for iY := Y / 3 * 3; iY < Y/3*3+3; iY++ {
		for iX := X / 3 * 3; iX < X/3*3+3; iX++ {
			grid[iY][iX][SetThisNumZero] = 0 // <
		}
	}

	for i := 0; i < 9; i++ {
		grid[i][X][SetThisNumZero] = 0
		grid[Y][i][SetThisNumZero] = 0
	}
	//PrintOnes(grid, SetThisNumZero)
}

func SolveOne(grid *[9][9][11]int, CC *bool) {	// if 001000000 => [9]=2

	*CC = false
	dbg[0] = dbg[1]
	dbg[1] = 5

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x][9] == 9 {

				grid[y][x][9] = Returnifone(grid[y][x])
				if grid[y][x][9] != 9 {

					grid[y][x][10] = 1

					SetCell(grid, y, x)
					*CC = true

					return
				}
			}
		}
	}
}

func SolveTwo(grid *[9][9][11]int, CC *bool) {	//esli iz 9ti jachejek, 8 zanjati
	*CC = false
	dbg[0] = dbg[1]
	dbg[1] = 6

	var sum_y, sum_x [9]int

	for y := 0; y < 9; y++ {
		for i := 0; i < 9; i++ {
			sum_x[i] = 0
			sum_y[i] = 0
		}
		for x := 0; x < 9; x++ {
			for i := 0; i < 9; i++ {
				sum_y[i] += grid[y][x][i]
				sum_x[i] += grid[x][y][i]
			}
		}

		//fmt.Printf("sumy%v sumx%v\n", sum_y, sum_x)

		for i := 0; i < 9; i++ { // NACHINAJA OTSUDA HERNJA KAKAJATA <<<<<
			if sum_y[i] == 1 {
				for n := 0; n < 9; n++ {
					if grid[y][n][i] == 1 {
						grid[y][n][9] = i
						grid[y][n][10] = 2
						SetCell(&*grid, y, n)
						*CC = true
						fmt.Print("sum_y: ")
						fmt.Println(sum_y)
						return
					}
				}
			}

			if sum_x[i] == 1 {
				for n := 0; n < 9; n++ {
					//fmt.Printf("grid[%v][%v][%v]=%v\n", n, y, i, grid[n][y][i])
					if grid[n][y][i] == 1 {
						grid[n][y][9] = i
						grid[n][y][10] = 2
						SetCell(&*grid, n, y)
						*CC = true
						fmt.Print("sum_x: ")
						fmt.Println(sum_x)
						return
					}
				}
			}
		}
		//// >>> zakanchivaetsa hernya

	}
}

func ClearHidden(grid *[9][9][11]int, CC *bool){	//esli grid[3][1][4] && grid[3][2][4] == 1 --> grid[3][0:9][4]=0 
	var sum_num [9]int
	var Y_check,X_check int
	for base_y:=0;base_y<9;base_y+=3{
		for base_x:=0;base_x<9;base_x+=3{

			for y:=0; y<base_y+3 ; y++{
				for x:=0; x<base_x+3 ; x++{
					for i:=0 ; i<9 ; i++{
						sum_num[i]+=grid[x][y][i]
					}
				}
			}

			Y_check,X_check=9,9
		for n:=0 ; n<9 ; n++{
			if sum_num[n] ==2 || sum_num[n]==3{

				for y:=0; y<base_y+3 ; y++{
					for x:=0; x<base_x+3 ; x++{
						if grid[y][x][n]==1{
							if Y_check==9{
								Y_check = y
								X_check = x
							}else if Y_

							
							
						}
					}
				}
				//proverit' stojat li v liniju?
			} 
		}

		for i:=0 ; i<9 ; i++{	//onulenie
			sum_num[i]=0
		}
		}

	}
}

func Solvesudoku(unsolved [9][9]int) {
	e := false
	CanContinue := true
	grid := GridInit(MinusOne(unsolved))

	ErrorCheck(grid, &e)

	grid = LineCheck(grid)

	for CanContinue {
		SolveOne(&grid, &CanContinue) //<------------ CHOTO NETO

		//ErrorCheck(grid, &e)
		MetaCheck(grid, &e)
		if e == true {
			fmt.Println("ERR")
			break
		}

		SolveTwo(&grid, &CanContinue)

		MetaCheck(grid, &e)
		ErrorCheck(grid, &e)
		if e == true {
			fmt.Println("ERR")
			break
		}
	}

	ErrorCheck(grid, &e)

	PrintGrid(grid)
}

func MetaCheck(grid [9][9][11]int, e *bool) {

	MetaArray := [9][9]int{
		{3, 8, 2, 1, 7, 5, 4, 9, 6},
		{4, 7, 5, 3, 6, 9, 2, 1, 8},
		{9, 1, 6, 2, 4, 8, 7, 5, 3},

		{8, 9, 4, 6, 2, 7, 5, 3, 1},
		{5, 2, 7, 8, 1, 3, 6, 4, 9},
		{1, 6, 3, 5, 9, 4, 8, 2, 7},

		{7, 3, 8, 4, 5, 1, 9, 6, 2},
		{6, 5, 1, 9, 8, 2, 3, 7, 4},
		{2, 4, 9, 7, 3, 6, 1, 8, 5}}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if grid[x][y][9] != 9 {
				if grid[x][y][9] != MetaArray[x][y]-1 {
					*e = true
					fmt.Printf("%v != MetaArray[%v][%v](%v)\tRIP\n", grid[x][y][9]+1, x, y, MetaArray[x][y])
				}
			}
		}
	}
}

func main() {

	sudo := [9][9]int{
		{0, 8, 0 /*|*/, 0, 0, 0 /*|*/, 0, 9, 0}, // 3 8 2	1 7 5	4 9 6
		{0, 7, 0 /*|*/, 0, 6, 0 /*|*/, 2, 1, 0}, // 4 7 5	3 6 9	2 1 8
		{0, 0, 6 /*|*/, 0, 4, 8 /*|*/, 7, 0, 0}, // 9 1 6	2 4 8	7 5 3
		//---------|-------------|----------
		{8, 0, 0 /*|*/, 0, 0, 0 /*|*/, 5, 3, 0}, // 8 9 4	6 2 7	5 3 1
		{0, 2, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0}, // 5 2 7	8 1 3	6 4 9
		{1, 6, 3 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0}, // 1 6 3	5 9 4	8 2 7
		//---------|-------------|----------
		{0, 0, 0 /*|*/, 4, 0, 1 /*|*/, 9, 0, 0}, // 7 3 8	4 5 1	9 6 2
		{0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 7, 0}, // 6 5 1	9 8 2	3 7 4
		{2, 0, 9 /*|*/, 7, 0, 0 /*|*/, 0, 0, 5}} // 2 4 9	7 3 6	1 8 5

	//	{5, 3, 0 /*|*/, 0, 7, 0 /*|*/, 0, 0, 0},
	//	{6, 0, 0 /*|*/, 1, 9, 5 /*|*/, 0, 0, 0},
	//	{0, 9, 8 /*|*/, 0, 0, 0 /*|*/, 0, 6, 0},
	//	//---------|-------------|----------
	//	{8, 0, 0 /*|*/, 0, 6, 0 /*|*/, 0, 0, 3},
	//	{4, 0, 0 /*|*/, 8, 0, 3 /*|*/, 0, 0, 1},
	//	{7, 0, 0 /*|*/, 0, 2, 0 /*|*/, 0, 0, 6},
	//	//---------|-------------|----------
	//	{0, 6, 0 /*|*/, 0, 0, 0 /*|*/, 2, 8, 0},
	//	{0, 0, 0 /*|*/, 4, 1, 9 /*|*/, 0, 0, 5},
	//	{0, 0, 0 /*|*/, 0, 8, 0 /*|*/, 0, 7, 9}}
	Solvesudoku(sudo)

}
