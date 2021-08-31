package main

import "fmt"

var dbg [3]int

func PrintGrid(grid *[9][9][11]int) {

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

			if grid_column[10] > 20 { // KRASNOE UBRAT'
				grid_column[10] -= 20
			}

			if grid_column[10] == 1 {
				fmt.Print(colorWhite)
			}
			if grid_column[10] == 2 {
				fmt.Print(colorPurple) //purp
			}

			if grid_column[10] == 3 { // ---
				fmt.Print(colorYellow)
				//grid[y][x][10] = 2
			}

			if grid_column[10] == 4 { //	|
				fmt.Print(colorBlue)
				//grid[y][x][10] = 2
			}

			if grid_column[10] == 5 { //	|
				fmt.Print(colorRed)
				//grid[y][x][10] = 2
			}

			if grid_column[10] > 20 { //	|
				fmt.Print(colorRed)
				grid[y][x][10] -= 20
			}

			fmt.Printf(" %v", string(grid_column[9]+1+'0'))

			fmt.Print(colorReset)
		}

		fmt.Println()
	}

	fmt.Println()
}

//1
func MinusOne(grid [9][9]int) [9][9]int {
	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 1
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

//2
func Returnifone(posibl [11]int) int { // nahodit' edinstvenniy "1"
	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 2

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

//3
func GridInit(unsolved [9][9]int) [9][9][11]int {
	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 3

	var grid [9][9][11]int

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
	var dbg_string [3]string

	for n := 0; n < 3; n++ {
		switch dbg[n] {
		case 1:
			dbg_string[n] = "MinusOne"

		case 2:
			dbg_string[n] = "Returnifone"

		case 3:
			dbg_string[n] = "GridInit"

		case 4:
			dbg_string[n] = "LineCheck"

		case 5:
			dbg_string[n] = "SquareCheck"

		case 6:
			dbg_string[n] = "PrintOnes"
		case 7:
			dbg_string[n] = "SetCell"

		case 8:
			dbg_string[n] = "SolveOne"

		case 9:
			dbg_string[n] = "SolveTwo"

		case 10:
			dbg_string[n] = "SolveCube"

		case 11:
			dbg_string[n] = "ClearHidden"

		default:
			dbg_string[n] = "!!ErrorErrora!!"
		}

	}

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

				fmt.Println("2prev func: ", dbg_string[0], "\n1prev func: ", dbg_string[1], "\nError func: ", dbg_string[2])
				*e = true
				PrintGrid(&grid)
			}
		}

		for i := 0; i < 9; i++ {
			sumX[i] = 0
			sumY[i] = 0
		}
	}

	//[kvadrat]
	for base_Z := 0; base_Z < 9; base_Z += 3 {
		for base_C := 0; base_C < 9; base_C += 3 {

			for zi := base_Z; zi < 3; zi++ {
				for ci := base_C; ci < 3; ci++ {
					if grid[zi][ci][9] != 9 {
						sumX[grid[zi][ci][9]] += 1
						if sumX[grid[zi][ci][9]] > 1 {
							fmt.Printf("\nERROR dupe [kvadrat] num:%v %v/%v \n", grid[zi][ci][9]+1, base_Z/3, base_C/3)
							fmt.Println("2prev func: ", dbg_string[0], "\n1prev func: ", dbg_string[1], "\nError func: ", dbg_string[2])
							*e = true
							PrintGrid(&grid)
						}
					}
				}
			}
			for i, _ := range sumX {
				sumX[i] = 0
			}
		}
	}
}

//4
func LineCheck(grid [9][9][11]int) [9][9][11]int {
	var tmp int

	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 4

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

//5
func SquareCheck(grid [9][9][11]int) [9][9][11]int {
	var snum [9]int

	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 5

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

//6
func PrintOnes(grid *[9][9][11]int, a int) {
	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 6

	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	// colorPurple := "\033[35m"
	// colorCyan := "\033[36m"

	for iY := 0; iY < 9; iY++ {
		for iX := 0; iX < 9; iX++ {

			if grid[iY][iX][10] == 7 {
				fmt.Print(colorGreen)
			} else if grid[iY][iX][10] == 7 {
				fmt.Print(colorRed)
			}

			if grid[iY][iX][9] == a {
				fmt.Print(a+1, " ")

			} else if grid[iY][iX][9] != 9 {
				fmt.Print("o ")

			} else {
				fmt.Print(grid[iY][iX][a], " ")

			}
			grid[iY][iX][10] = 0
			fmt.Print(colorReset)
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
}

//7
func SetCell(grid *[9][9][11]int, Y int, X int) {
	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 7

	SetThisNumZero := grid[Y][X][9]

	for iY := Y / 3 * 3; iY < Y/3*3+3; iY++ {
		for iX := X / 3 * 3; iX < X/3*3+3; iX++ {
			grid[iY][iX][SetThisNumZero] = 0 // <
		}
	}

	for i := 0; i < 9; i++ {
		grid[Y][X][i] = 0
		grid[i][X][SetThisNumZero] = 0
		grid[Y][i][SetThisNumZero] = 0
	}
}

//8
func SolveOne(grid *[9][9][11]int, CC *bool) { // if 001000000 => [9]=2

	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 8

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x][9] == 9 {

				grid[y][x][9] = Returnifone(grid[y][x])
				if grid[y][x][9] != 9 {

					grid[y][x][10] = 21

					SetCell(grid, y, x)
					*CC = true

					return
				}
			}
		}
	}
}

//9
func SolveTwo(grid *[9][9][11]int, CC *bool) { //esli summa[i] na linii == 1 => SetCell

	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 9

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
						grid[y][n][10] = 22

						SetCell(&*grid, y, n)
						*CC = true
						return
					}
				}
			}

			if sum_x[i] == 1 {
				for n := 0; n < 9; n++ {
					//fmt.Printf("grid[%v][%v][%v]=%v\n", n, y, i, grid[n][y][i])
					if grid[n][y][i] == 1 {

						grid[n][y][9] = i
						grid[n][y][10] = 22

						SetCell(&*grid, n, y)
						*CC = true
						return
					}
				}
			}
		}
		//// >>> zakanchivaetsa hernya

	}
}

//10
func SolveTwoCube(grid *[9][9][11]int, CC *bool) { // esli summa[i] v 3x3 kvadrate == 1 => SetCell

	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 10

	//fmt.Println("vhod v solve2cube\n 1 2 3 4 5 6 7 8 9")
	var sum_cube [9]int

	for base_y := 0; base_y < 9; base_y += 3 {
		for base_x := 0; base_x < 9; base_x += 3 {
			for i, _ := range sum_cube { //onulenie
				sum_cube[i] = 0
			}

			for y := base_y; y < base_y+3; y++ { //	summa vseh shansov v 3x3 kvadrate (potom esli kakoita == 1 => setcell)
				for x := base_x; x < base_x+3; x++ {
					for i, OneNine := range grid[y][x][0:9] {
						sum_cube[i] += OneNine
					}
				}
			}

			for i, sumi := range sum_cube {
				if sumi == 1 {

					for y := base_y; y < base_y+3; y++ {
						for x := base_x; x < base_x+3; x++ {
							if grid[y][x][i] == 1 { //
								//	fmt.Printf("\033[32m%v[%v][%v][%v]=%v\n\033[0m", grid[y][x][i], y, x, i, i)
								if grid[y][x][9] != 9 {
									fmt.Printf("ERROR!\nTryting to swap %v with %v\n", grid[y][x][9], i) // <---== Ubrat' ESLI VSE OK
								}
								*CC = true
								grid[y][x][9] = i
								SetCell(*&grid, y, x)
								grid[y][x][10] = 23 //YELLOW
							}

						}
					}
				}

			}

		}
	}

}

// NAJTI V CHEM PROBLEMA
//11
func ClearHidden(grid *[9][9][11]int, CC *bool) { //esli grid[3][1][4] && grid[3][2][4] == 1 --> grid[3][0:9][4]=0
	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 11

	var sum_num [9]int
	var Y_check, X_check int

	var y_change, x_change, Z_skip int
	var counter_ptr *int

	for base_y := 0; base_y < 9; base_y += 3 {
		for base_x := 0; base_x < 9; base_x += 3 {

			for y := base_y; y < base_y+3; y++ {
				for x := base_x; x < base_x+3; x++ {
					for i := 0; i < 9; i++ {
						if grid[y][x][9] == i && grid[y][x][i] == 1 {
							fmt.Printf("\nCHTOTO POSHLO NE TAK: grid[%v][%v][9] == %v && grid[y][x][i] == 1\n", y, x, i)
						}
						sum_num[i] += grid[y][x][i]

					}
				}
			}

			fmt.Print(sum_num, " ")
			/// PRODOLZHIT' NIZHE
			Y_check, X_check = 9, 9
			for n := 0; n < 9; n++ {
				if sum_num[n] == 2 || sum_num[n] == 3 {

					for y := base_y; y < base_y+3; y++ {
						for x := base_x; x < base_x+3; x++ {
							if grid[y][x][n] == 1 {
								if Y_check == 9 {
									Y_check = y
									X_check = x
								} else {

									if Y_check != y {
										Y_check = 11
									}
									if X_check != x {
										X_check = 11
									}

								}
							}
						}
					}

					if X_check == 9 && Y_check == 9 {
						fmt.Printf("????err??CHECK == 9\t n: %v  check%v/%v   sum:%v\n", n, base_y, base_x, sum_num[n])
						//break
					}

					if X_check == 11 && Y_check == 11 {
						fmt.Printf("errCHECK == 11\t")
						break
					}

					//proverit' stojat li v liniju?

					/// !! NEZABIT' OBNULAT' VSE!!!
					y_change = 0
					x_change = 0

					if Y_check != 11 {
						y_change = Y_check
						counter_ptr = &x_change
						Z_skip = base_x
					} else if X_check != 11 {
						x_change = X_check
						counter_ptr = &y_change
						Z_skip = base_y
					} else {
						break
					}

					if y_change == 11 || x_change == 11 {
						fmt.Println("ERROR")
					}

					//			fmt.Println("n:", n+1, "  y_change:", y_change, "   x_change:", y_change)
					//		PrintOnes(&*grid, n)
					for *counter_ptr = 0; *counter_ptr < 9; *counter_ptr++ {
						grid[y_change][x_change][10] = 8
						if *counter_ptr < Z_skip || *counter_ptr > Z_skip+2 {
							grid[y_change][x_change][n] = 0
							grid[y_change][x_change][10] = 7
						}
					}
					//		PrintOnes(grid, n)

					// if Z!=11 {

					// 	for Z// sdelat' wtobi for propuskal base_y/x +3
					// }
				}
			}

			for i := 0; i < 9; i++ { //onulenie
				sum_num[i] = 0
			}
		}
		fmt.Println()
	}
}

func Solvesudoku(unsolved [9][9]int) {
	e := false
	CanContinue := true
	grid := GridInit(MinusOne(unsolved))

	ErrorCheck(grid, &e)
	grid = LineCheck(grid)
	PrintGrid(&grid)

	for CanContinue {
		CanContinue = false

		SolveOne(&grid, &CanContinue)

		ErrorCheck(grid, &e)
		//MetaCheck(grid, &e)
		if e == true {
			fmt.Println("ERRor")
			break
		}

		SolveTwo(&grid, &CanContinue)

		//MetaCheck(grid, &e)
		ErrorCheck(grid, &e)
		if e == true {
			fmt.Println("ERRor")
			break
		}

		SolveTwoCube(&grid, &CanContinue)
		//PrintGrid(&grid)
		//MetaCheck(grid, &e)
		ErrorCheck(grid, &e)
		if e == true {
			fmt.Println("ERRor")
			break
		}

	}

	//PrintOnes(&grid, 2)
	ClearHidden(&grid, &CanContinue)

	ErrorCheck(grid, &e)

	fmt.Println()
	fmt.Println()
	PrintGrid(&grid)
	fmt.Println()

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

		{0, 0, 0 /*|*/, 0, 9, 1 /*|*/, 0, 8, 0},
		{0, 0, 0 /*|*/, 0, 0, 3 /*|*/, 0, 1, 0},
		{0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 2, 7, 4},
		//---------|-------------|----------
		{0, 8, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 2},
		{0, 0, 1 /*|*/, 0, 0, 0 /*|*/, 6, 0, 0},
		{0, 9, 0 /*|*/, 6, 0, 5 /*|*/, 0, 0, 0},
		//---------|-------------|----------
		{2, 0, 0 /*|*/, 0, 7, 0 /*|*/, 8, 0, 0},
		{7, 0, 4 /*|*/, 1, 0, 0 /*|*/, 0, 0, 0},
		{0, 0, 0 /*|*/, 0, 4, 0 /*|*/, 0, 0, 0}}

	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// //---------|-------------|----------
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// //---------|-------------|----------
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0},
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0}}

	// {0, 8, 0 /*|*/, 0, 0, 0 /*|*/, 0, 9, 0}, // 3 8 2	1 7 5	4 9 6
	// {0, 7, 0 /*|*/, 0, 6, 0 /*|*/, 2, 1, 0}, // 4 7 5	3 6 9	2 1 8
	// {0, 0, 6 /*|*/, 0, 4, 8 /*|*/, 7, 0, 0}, // 9 1 6	2 4 8	7 5 3
	// //---------|-------------|----------
	// {8, 0, 0 /*|*/, 0, 0, 0 /*|*/, 5, 3, 0}, // 8 9 4	6 2 7	5 3 1
	// {0, 2, 0 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0}, // 5 2 7	8 1 3	6 4 9
	// {1, 6, 3 /*|*/, 0, 0, 0 /*|*/, 0, 0, 0}, // 1 6 3	5 9 4	8 2 7
	// //---------|-------------|----------
	// {0, 0, 0 /*|*/, 4, 0, 1 /*|*/, 9, 0, 0}, // 7 3 8	4 5 1	9 6 2
	// {0, 0, 0 /*|*/, 0, 0, 0 /*|*/, 0, 7, 0}, // 6 5 1	9 8 2	3 7 4
	// {2, 0, 9 /*|*/, 7, 0, 0 /*|*/, 0, 0, 5}} // 2 4 9	7 3 6	1 8 5

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
