package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Can't exec Run, %v", err)
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Can't exec Run, %v", err)
		}
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //MacOS with intel proc
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Can't exec Run, %v", err)
		}
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

type Cell struct {
	Column         int
	Row            int
	Fill           int // пусто не пусто
	Fixed          int // зафиксировано
	Fall           int // в состоянии падения
	CenterOfFigure int
}

type Field [10][20]Cell

func CreateField() *Field {
	var field Field
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {

			field[c][r].Column = c
			field[c][r].Row = r
			field[c][r].Fill = 0
			field[c][r].Fixed = 0
		}
	}
	return &field
}

// getRand return random int using lim (if lim = 7 then return 0-6)
func getRand(lim int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(lim)
}

type FigureCell struct {
	RotationCenter int
	Fill           int
}

type Figure [4][4]FigureCell

var (
	J Figure
	L Figure
	T Figure
	O Figure
	Z Figure
	S Figure
	I Figure
)

func CreateFigure() {
	I[1][0].Fill = 1
	I[1][1].Fill = 1
	I[1][2].Fill = 1
	I[1][3].Fill = 1
	I[1][1].RotationCenter = 1

	J[1][0].Fill = 1
	J[1][1].Fill = 1
	J[1][2].Fill = 1
	J[2][2].Fill = 1
	J[1][1].RotationCenter = 1

	L[1][0].Fill = 1
	L[1][1].Fill = 1
	L[1][2].Fill = 1
	L[2][0].Fill = 1
	L[1][1].RotationCenter = 1

	O[1][0].Fill = 1
	O[1][1].Fill = 1
	O[2][0].Fill = 1
	O[2][1].Fill = 1

	S[1][1].Fill = 1
	S[1][2].Fill = 1
	S[2][0].Fill = 1
	S[2][1].Fill = 1
	S[1][1].RotationCenter = 1

	T[1][0].Fill = 1
	T[1][1].Fill = 1
	T[1][2].Fill = 1
	T[2][1].Fill = 1
	T[1][1].RotationCenter = 1

	Z[1][0].Fill = 1
	Z[1][1].Fill = 1
	Z[2][1].Fill = 1
	Z[2][2].Fill = 1
	Z[1][1].RotationCenter = 1

}

func showField(f Field) {
	for r := 0; r < cap(f[0]); r++ {
		for c := 0; c < cap(f); c++ {
			//if f[c][r].Fill == 1 {
			//	fmt.Print("◽")
			//} else {
			//	fmt.Print("◾")
			//}
			fmt.Print(f[c][r].Fill)
		}
		fmt.Println()
	}
}

func showCenterOfFigure(f Field) {
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {

			fmt.Print(f[c][r].CenterOfFigure, " ")
		}
		fmt.Println()
	}
}

func showFigure(f [4][4]Cell) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			fmt.Print(f[c][r].Fill)
		}
		fmt.Println()
	}
}

func getRandFigure() Figure {
	rnd := getRand(7)
	switch rnd {
	case 0:
		return O
	case 1:
		return L
	case 2:
		return J
	case 3:
		return T
	case 4:
		return Z
	case 5:
		return S
	case 6:
		return I
	default:
		return O
	}
}

func SpawnAdvancedFigureNew(a Figure, s int, f *Field) {
	actFigure := a //getRandFigure()
	spawnCol := s  //:= getRand(8)

	switch actFigure {
	case O:
		spawnCol = s //getRand(9)
	case I:
		spawnCol = s //getRand(7)
	}

	fc := 0
	fr := 0
	for r := 0; r < len(actFigure); r++ {
		for c := spawnCol; c < len(actFigure)+spawnCol; c++ {
			if actFigure[fr][fc].Fill == 1 {
				f[c][r].Fill = 1
				f[c][r].Fall = 1
				if actFigure[fr][fc].RotationCenter == 1 {
					f[c][r].CenterOfFigure = 1
				}
			}
			fc++
		}
		fc = 0
		fr++
	}
}

func SpawnFigureNew(f *Field) {
	//спавн только на пустое место в разумных пределах от центра

	actFigure := getRandFigure()
	spawnCol := getRand(8)

	switch actFigure {
	case O:
		spawnCol = getRand(9)
	case I:
		spawnCol = getRand(7)
	default:

	}

	fc := 0
	fr := 0
	for r := 0; r < len(actFigure); r++ {
		for c := spawnCol; c < len(actFigure)+spawnCol; c++ {
			if actFigure[fr][fc].Fill == 1 {
				f[c][r].Fill = 1
				f[c][r].Fall = 1
				if actFigure[fr][fc].RotationCenter == 1 {
					f[c][r].CenterOfFigure = 1
				}
			}
			fc++
		}
		fc = 0
		fr++
	}
}

func canFall() bool {
	return false
}

func getLowerCells(f *Field) int {
	var lr int
	for r := 19; r > 0; r-- {
		for c := 0; c < 10; c++ {
			if f[c][r].Fall == 1 {
				lr = r
				return lr
			}
		}
	}
	return lr
}

func clearLine(f *Field, row int) {
	for c := 0; c < 10; c++ {
		f[c][row].Fill = 0
		f[c][row].Fixed = 0
		f[c][row].CenterOfFigure = 0
	}
	showField(*f)
	time.Sleep(200 * time.Millisecond)
	// something here place counter (payer score)
}

func moveAllUpperCellsDown(f *Field, row int) {
	for r := row; r > 0; r-- {
		for c := 0; c < 10; c++ {
			f[c][r].Fill = f[c][r-1].Fill
			f[c][r].Fixed = f[c][r-1].Fixed
			f[c][r].CenterOfFigure = f[c][r-1].CenterOfFigure
			showField(*f)
		}
	}

}

func checkFullLine(f *Field) {

	for r := 19; r >= 0; r-- {
		temp := 0
		for c := 0; c < 10; c++ {
			temp = temp + f[c][r].Fixed
		}
		if temp == 10 {
			clearLine(f, r)
			moveAllUpperCellsDown(f, r)
		}
	}
}

func fixFigure(f *Field) {
	for r := 19; r > 0; r-- {
		for c := 0; c < 10; c++ {
			if f[c][r].Fall == 1 {
				f[c][r].Fall = 0
				f[c][r].Fixed = 1
				f[c][r].CenterOfFigure = 0
			}
		}
	}

	checkFullLine(f)
}

func findFigureCells(f *Field) [4][2]int {
	var res [4][2]int
	var rc, rr int
	for r := 19; r > 0; r-- {
		for c := 0; c < 10; c++ {

			if f[c][r].Fall == 1 {
				res[rc][rr] = c
				res[rc][rr+1] = r
				rc++
				rr = 0
			}
		}
	}
	return res
}

func FallFigure(f *Field, ch chan string) {
	for i := 0; i < 19; i++ {

		if <-ch != "" {
			fmt.Println("ch correct transfer to FallFigure")
			time.Sleep(1 * time.Second)
		}

		// проверка достижения нижней линии
		if i > 15 {
			lowerRow := getLowerCells(f)
			if lowerRow == 19 {
				showField(*f)
				fixFigure(f)
				fmt.Println("rich to the end of field")
				return
			}
		}

		// проверка падения на другую фигуру
		t := findFigureCells(f)
		for _, row := range t {
			if f[row[0]][row[1]+1].Fixed == 1 {
				fixFigure(f)
				return
			}
		}

		for r := 18; r > 0; r-- { //выяснить почему с 18 строки
			for c := 0; c < 10; c++ {
				if f[c][r].Fall == 1 {
					if f[c][r].CenterOfFigure == 1 {
						f[c][r].CenterOfFigure = 0
						f[c][r+1].CenterOfFigure = 1
					}
					f[c][r].Fill = 0
					f[c][r+1].Fill = 1

					f[c][r].Fall = 0
					f[c][r+1].Fall = 1
				}
			}

		}
		showField(*f)
		time.Sleep(100 * time.Millisecond)
	}
}

// FindCenterOfFigure return col and row where are placed center of the falling figure
func FindCenterOfFigure(f *Field) (int, int) {
	var col, row int = -1, -1
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {
			if f[c][r].CenterOfFigure == 1 {
				col = f[c][r].Column
				row = f[c][r].Row
				break

			}
		}

		if col != -1 {
			break
		}
	}
	return col, row
}

// RotateFigure get f - main *field and dir - direction rotation "right" or "left"
//func RotateFigure(f *Field, dir string) {
//	col, row := FindCenterOfFigure(f) // return col + row
//	var temp [4][4]Cell
//	var tc, tr int
//
//	for r := row - 1; r < row+2; r++ {
//		for c := col - 1; c < col+2; c++ {
//			temp[tr][tc] = f[r][c]
//			tc++
//		}
//		tc = 0
//		tr++
//	}
//
//	switch dir {
//	case "left":
//		t2c := col - 1
//		t2r := row - 1
//		for c := 0; c < 3; c++ {
//			for r := 2; r > -1; r-- {
//				f[t2r][t2c] = temp[r][c]
//				t2c++
//			}
//			t2c = 0
//			t2r++
//		}
//	case "right":
//		t2c := col - 1
//		t2r := row - 1
//		for c := 2; c > -1; c-- {
//			for r := 0; r < 3; r++ {
//				f[t2r][t2c] = temp[r][c]
//				t2c++
//			}
//			t2c = 0
//			t2r++
//		}
//	}
//
//	fmt.Println()
//	showFigure(temp)
//	fmt.Println()
//
//}

// MoveFigure -- f - main *field and dir - direction
func MoveFigure(f *Field, dir string) {
	//col, row := FindCenterOfFigure(f)
	switch dir {
	case "left":
		for r := 0; r < 20; r++ {
			for c := 0; c < 10; c++ {
				if f[c][r].Fill == 1 {
					if f[c][r].Fill == 1 {
						f[c][r].Fill = 0
						f[c-1][r].Fill = 1

						if f[c][r].CenterOfFigure == 1 {
							f[c][r].CenterOfFigure = 0
							f[c-1][r].CenterOfFigure = 1
						}
					}
				}
			}
		}
	case "right":
		for r := 0; r < 20; r++ {
			for c := 9; c != -1; c-- {
				if f[c][r].Fill == 1 {
					f[c][r].Fill = 0
					f[c+1][r].Fill = 1

					if f[c][r].CenterOfFigure == 1 {
						f[c][r].CenterOfFigure = 0
						f[c+1][r].CenterOfFigure = 1
					}
				}
			}
		}
	}
}

func getKey(chKey chan uint16) {
	for {
		_, s, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if s != 0 {
			chKey <- uint16(s)
		} else {
			chKey <- 0
		}

		//switch s {
		//case 65517:
		//	fmt.Println("вверх - команда повернуть")
		//case 65515:
		//	fmt.Println("влево - команда сдвиг влево")
		//case 65514:
		//	fmt.Println("вправо - команда сдвиг вправо")
		//case 65516:
		//	fmt.Println("вниз - команда быстро опустить")
		//}
	}

}

func startGame(ch chan string) {
	field := CreateField()
	CreateFigure()

	for {
		SpawnFigureNew(field)
		FallFigure(field, ch)
	}

}

func main() {

	ch := make(chan string)
	chKey := make(chan uint16)
	go startGame(ch)
	go getKey(chKey)
	for r := range chKey {
		switch r {
		case 65517:
			ch <- "turn"
			fmt.Println("повернуть")
		case 65515:
			fmt.Println("влево")
		case 65514:
			fmt.Println("вправо")
		case 65516:
			fmt.Println("вниз")
		default:

		}
		ch <- ""
	}

	//showField(*field)

	//time.Sleep(60 * time.Second)

	//for i := 0; i < 10; i++ {
	//SpawnAdvancedFigureNew(Z, 0, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(Z, 2, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(Z, 4, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(Z, 6, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(Z, 7, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(L, 0, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(J, 7, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(I, 3, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(O, 0, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(O, 2, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(O, 4, field)
	//FallFigure(field)
	//SpawnAdvancedFigureNew(L, 6, field)
	//FallFigure(field)
	//}

	//for i := 0; i < 40; i++ {
	//	SpawnFigureNew(field)
	//	FallFigure(field)
	//}
}
