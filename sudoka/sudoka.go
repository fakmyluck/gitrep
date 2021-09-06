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

func SumRow(row [9]int) int {
	var sum int

	for _, num := range row {
		sum += num
	}
	return sum
}

func SumGridRow(gridrow [11]int) int {
	var sum int

	for i := 0; i < 9; i++ {
		sum += gridrow[i]
	}
	return sum
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
	colorCyan := "\033[36m"
	fmt.Println(a + 1)
	for iY := 0; iY < 9; iY++ {
		for iX := 0; iX < 9; iX++ {

			if grid[iY][iX][10] == 7 {
				fmt.Print(colorGreen)
			} else if grid[iY][iX][10] == 6 {
				fmt.Print(colorRed)
			}

			if grid[iY][iX][9] == a {
				fmt.Print(colorCyan, a+1, " ")

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

	fmt.Printf("\n8/6%v\t\t!!!\n9/6%v\n", grid[7][5], grid[8][5])
	fmt.Printf("6/3%v\n", grid[6][3])
	PrintGrid(&*grid)
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
						grid[y][n][10] = 2

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
						grid[n][y][10] = 2

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

								if grid[y][x][9] != 9 {
									fmt.Printf("ERROR!\nTryting to swap %v with %v\n", grid[y][x][9], i) // <---== Ubrat' ESLI VSE OK
								}
								*CC = true
								grid[y][x][9] = i

								grid[y][x][10] = 3 //YELLOW
								SetCell(grid, y, x)
							}

						}
					}
				}

			}

		}
	}

}

//11
func ClearHidden(grid *[9][9][11]int) { //esli grid[3][1][4] && grid[3][2][4] == 1 --> grid[3][0:9][4]=0
	dbg[0] = dbg[1]
	dbg[1] = dbg[2]
	dbg[2] = 11

	//tmparray := grid
	var xck, yck int

	for cnum := 0; cnum < 9; cnum++ {

		for base_y := 0; base_y < 9; base_y += 3 {
			for base_x := 0; base_x < 9; base_x += 3 {

				xck, yck = 9, 9

				for y := base_y; y < base_y+3; y++ {
					for x := base_x; x < base_x+3; x++ {

						if grid[y][x][9] == cnum {
							goto skipsquare
						}
						if grid[y][x][cnum] == 1 {
							if yck == 9 {
								yck, xck = y, x
							} else {
								if yck != y {
									yck = 11
								}
								if xck != x {
									xck = 11
								}
							}
						}
					}
				}

				if yck == 9 {
					fmt.Printf("cheto poshlo ne tak: yck=%v  xck=%v  cnum=%v\n", yck, xck, cnum)
				}

				if yck < 9 {
					for i := 0; i < 9; i++ {
						if i < base_x || i > base_x+2 {
							grid[yck][i][cnum] = 0
							grid[yck][i][10] = 6 //debug
						}
					}
					//PrintOnes(grid, cnum)
				}

				if xck < 9 {
					for i := 0; i < 9; i++ {
						if i < base_y || i > base_y+2 {
							grid[i][xck][cnum] = 0
							grid[i][xck][10] = 6 //debug
						}
					}
					//PrintOnes(grid, cnum)
				}
			skipsquare:
			}
		} //base_y

	} //cnum
}

func FullCellLine(grid *[9][9][11]int) {
	fmt.Println("SURPRISE")
	var sumyx, sumxy [9][9]int
	// var sumxy[9] int		REMOVE
	var cell1, cell2 int

	for y := 0; y < 9; y++ {

		for n := 0; n < 9; n++ {

			for x := 0; x < 9; x++ {

				sumyx[n][x] = grid[y][x][n] //grid[y][x][n] => esli 1 to v kletke mozhet bit' cifra 'n'
				sumxy[n][x] = grid[x][y][n] //sum_[n][x] => esli 1 to cifra 'n' nahoditsa kakih kletkah
			}
		}
		cell1, cell2 = 99, 99
		//proveeka na sovpodenie
		for n1 := 0; n1 < 9-1; n1++ {
			/*	if SumRow(sumyx[n1]) == 2 { //POstrochno ----

				for n2 := n1 + 1; n2 < 9; n2++ {
					if sumyx[n1] == sumyx[n2] {
						//√ naideno

						fmt.Printf("√yx y:%v\n%v %v\n%v %v\n\n", y, sumyx[n1], grid[y][n1], sumyx[n2], grid[y][n2])

						for i := 0; i < 9; i++ {
							if sumyx[n1][i] == 1 {
								if cell1 == 99 {
									cell1 = i
								} else {
									cell2 = i
									break
								}
							}
						}
						fmt.Print("\n", sumyx[n1], ' ', grid[y][n1], '\n', sumyx[n1], ' ', grid[y][n2], '\n')
						for i := 0; i < 9; i++ {

							if i != n1 && i != n2 {
								grid[y][cell1][i] = 0
								grid[y][cell2][i] = 0
								//grid[y][X][10]=6
							}

						}

						fmt.Print("\n", grid[y][n1], '\n', grid[y][n2], '\n')
						fmt.Printf("cell1:%v  cell2:%v  n1:%v n2:%v", cell1, cell2, n1, n2)
					}
				}

			}*/
			//ispravit' verhnuju po (takzhe kak nizhe)
			//neuveren wto pravel'no
			if SumRow(sumxy[n1]) == 2 { // ||| po ryadam

				for n2 := n1 + 1; n2 < 9; n2++ {
					if sumxy[n1] == sumxy[n2] {

						for i := 0; i < 9; i++ { //basically i == x
							if sumxy[n1][i] == 1 {
								if cell1 == 99 {
									cell1 = i
								} else {
									cell2 = i
									break
								}
							}
						}

						//√ naideno					//PS! gridn1 NEPRAVEL'no!?
						// fmt.Printf("\n\n|  1 2 3 4 5 6 7 8 9\t\t√xy x:%v\ns1%v%v gridn1%v\ns2%v%v gridn2%v\n\n", y+1, sumxy[n1], n1+1, grid[cell1][y], sumxy[n2], n2+1, grid[cell2][y])
						if !(grid[cell1][y] == grid[cell2][y] && SumGridRow(grid[cell1][y]) == 2) {
							for i := 0; i < 9; i++ {

								if i != n1 && i != n2 {
									grid[cell1][y][i] = 0
									grid[cell2][y][i] = 0
									//grid[x][y][10]=6
								}

							}
						}

						// fmt.Printf("\ns1%v gridn1%v\ns2%v gridn2%v\n\n", sumxy[n1], grid[cell1][y], sumxy[n2], grid[cell2][y])
						// fmt.Printf("cell1:%v  cell2:%v  n1:%v n2:%v", cell1+1, cell2+1, n1+1, n2+1)
					}
				}
				//fmt.Println()
				//PrintGrid(&*grid)
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
		CanContinue = false
		ClearHidden(&grid)

		SolveOne(&grid, &CanContinue)

		ErrorCheck(grid, &e)
		MetaCheck(grid, &e)
		if e == true {
			fmt.Println("ERRor")
			break
		}

		SolveTwo(&grid, &CanContinue)

		MetaCheck(grid, &e)
		ErrorCheck(grid, &e)
		if e == true {
			fmt.Println("ERRor")
			break
		}

		SolveTwoCube(&grid, &CanContinue)

		MetaCheck(grid, &e)
		ErrorCheck(grid, &e)
		if e == true {
			fmt.Println("ERRor")
			break
		}
		FullCellLine(&grid)

	}

	//PrintOnes(&grid, 2)

	ErrorCheck(grid, &e)

	fmt.Println()
	fmt.Println()
	PrintGrid(&grid)
	fmt.Println()

}

func MetaCheck(grid [9][9][11]int, e *bool) {

	MetaArray := [9][9]int{

		{6, 2, 7, 4, 9, 1, 5, 8, 3},
		{8, 4, 5, 7, 2, 3, 9, 1, 6},
		{1, 3, 9, 8, 5, 6, 2, 7, 4},

		{3, 8, 6, 9, 1, 7, 4, 5, 2},
		{5, 7, 1, 2, 3, 4, 6, 9, 8},
		{4, 9, 2, 6, 8, 5, 1, 3, 7},

		{2, 6, 3, 5, 7, 9, 8, 4, 1},
		{7, 5, 4, 1, 6, 8, 3, 2, 9},
		{9, 1, 8, 3, 4, 2, 7, 6, 5}}

	// {3, 8, 2, 1, 7, 5, 4, 9, 6},
	// {4, 7, 5, 3, 6, 9, 2, 1, 8},
	// {9, 1, 6, 2, 4, 8, 7, 5, 3},

	// {8, 9, 4, 6, 2, 7, 5, 3, 1},
	// {5, 2, 7, 8, 1, 3, 6, 4, 9},
	// {1, 6, 3, 5, 9, 4, 8, 2, 7},

	// {7, 3, 8, 4, 5, 1, 9, 6, 2},
	// {6, 5, 1, 9, 8, 2, 3, 7, 4},
	// {2, 4, 9, 7, 3, 6, 1, 8, 5}}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if grid[x][y][9] != 9 {
				if grid[x][y][9] != MetaArray[x][y]-1 {
					*e = true
					fmt.Printf("%v != MetaArray[%v][%v](%v)\tRIP\n", grid[x][y][9]+1, x+1, 1+y, MetaArray[x][y])
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
