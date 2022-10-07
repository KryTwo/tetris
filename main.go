/*
TODO:
1. Взаимодействие с фигурой I
2. Доработать спавн
3. Реализовать быстрое падение фигуры
4. Добавить фантом?
5...
*/
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

type field [10][20]Cell

var Field field

func CreateField(f *field) {
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {
			f[c][r].Column = c
			f[c][r].Row = r
			f[c][r].Fill = 0
			f[c][r].Fixed = 0
		}
	}
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
func showFieldOnce() {

	for r := 0; r < cap(Field[0]); r++ {
		for c := 0; c < cap(Field); c++ {
			if Field[c][r].Fill == 1 {
				fmt.Print("◽")
			} else {
				fmt.Print("◾")
			}
			//fmt.Printf("%d-%d|%d", Field[c][r].Column, Field[c][r].Row, Field[c][r].CenterOfFigure)
			fmt.Print(Field[c][r].Fall)
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(50 * time.Millisecond)

}

func showField() {
	//for {
	for r := 0; r < cap(Field[0]); r++ {
		for c := 0; c < cap(Field); c++ {
			if Field[c][r].Fill == 1 {
				fmt.Print("◽")
			} else {
				fmt.Print("◾")
			}
			//fmt.Printf("%d-%d", Field[c][r].Column, Field[c][r].Row)
			//fmt.Print(Field[c][r].CenterOfFigure)
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(1000 * time.Millisecond)
	//}
}

func showCenterOfFigure() {
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {

			fmt.Print(Field[c][r].CenterOfFigure, " ")
		}
		fmt.Println()
	}
}

func showFigure(f [4][4]Cell) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			fmt.Print(Field[c][r].Fill)
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

func SpawnAdvancedFigure(a Figure, s int) {
	ActFigure = a //getRandFigure()
	spawnCol := s //:= getRand(8)

	switch ActFigure {
	case O:
		spawnCol = s //getRand(9)
	case I:
		spawnCol = s //getRand(7)
	}

	fc := 0
	fr := 0
	for r := 0; r < len(ActFigure); r++ {
		for c := spawnCol; c < len(ActFigure)+spawnCol; c++ {
			if ActFigure[fr][fc].Fill == 1 {
				Field[c][r].Fill = 1
				Field[c][r].Fall = 1
				if ActFigure[fr][fc].RotationCenter == 1 {
					Field[c][r].CenterOfFigure = 1
				}
			}
			fc++
		}
		fc = 0
		fr++
	}
}

var ActFigure Figure

func SpawnFigure() {
	//спавн только на пустое место в разумных пределах от центра

	ActFigure = getRandFigure()
	spawnCol := getRand(8)

	switch ActFigure {
	case O:
		spawnCol = getRand(9)
	case I:
		spawnCol = getRand(7)
	default:

	}

	fc := 0
	fr := 0
	for r := 0; r < len(ActFigure); r++ {
		for c := spawnCol; c < len(ActFigure)+spawnCol; c++ {
			if ActFigure[fr][fc].Fill == 1 {
				Field[c][r].Fill = 1
				Field[c][r].Fall = 1
				if ActFigure[fr][fc].RotationCenter == 1 {
					Field[c][r].CenterOfFigure = 1
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

func getLowerCells() int {
	var lr int
	for r := 19; r > 0; r-- {
		for c := 0; c < 10; c++ {
			if Field[c][r].Fall == 1 {
				lr = r
				return lr
			}
		}
	}
	return lr
}

func clearLine(row int) {
	for c := 0; c < 10; c++ {
		Field[c][row].Fill = 0
		Field[c][row].Fixed = 0
		Field[c][row].CenterOfFigure = 0
	}

	//time.Sleep(200 * time.Millisecond)
	// something here place counter (payer score)
}

func moveAllUpperCellsDown(row int) {
	for r := row; r > 0; r-- {
		for c := 0; c < 10; c++ {
			Field[c][r].Fill = Field[c][r-1].Fill
			Field[c][r].Fixed = Field[c][r-1].Fixed
			Field[c][r].CenterOfFigure = Field[c][r-1].CenterOfFigure

		}
	}

}

func checkFullLine() {

	for r := 19; r >= 0; r-- {
		temp := 0
		for c := 0; c < 10; c++ {
			temp = temp + Field[c][r].Fixed
		}
		if temp == 10 {
			clearLine(r)
			moveAllUpperCellsDown(r)
		}
	}
}

func fixFigure() {
	for r := 19; r > -1; r-- {
		for c := 0; c < 10; c++ {
			if Field[c][r].Fall == 1 {
				Field[c][r].Fall = 0
				Field[c][r].Fixed = 1
				Field[c][r].CenterOfFigure = 0
			}
		}
	}

	checkFullLine()
}

func findFigureCells() [4][2]int {
	var res [4][2]int
	var rc, rr int
	for r := 19; r > 0; r-- {
		for c := 0; c < 10; c++ {

			if Field[c][r].Fall == 1 {
				res[rc][rr] = c
				res[rc][rr+1] = r
				rc++
				rr = 0
			}
		}
	}
	return res
}

// FindCenterOfFigure return col and row where are placed center of the falling figure
func FindCenterOfFigure() (int, int) {
	var col, row int = -1, -1
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {
			if Field[c][r].CenterOfFigure == 1 {
				col = Field[c][r].Column
				row = Field[c][r].Row
				return col, row

			}
		}
	}
	return col, row
}

//func (c *Cell) clear() {
//	c.Fill = 0
//	c.Fixed = 0
//	c.CenterOfFigure = 0
//	c.Fall = 0
//}

func canRotate(temp field, col, row int) bool {
	var tc, tr int
	for c := col - 1; c < col+2; c++ {
		for r := row + 1; r > row-2; r-- {
			if Field[c][r].Fixed == 1 && temp[tc][tr].Fill == 1 {
				return false
			}
			tc++
		}
		tc = 0
		tr++
	}
	tr = 0
	return true
}

func rollback(temp field) {
	Field = temp
}

func tryMoveAndRotate(temp field) bool {
	col, row := FindCenterOfFigure()

	if col == 0 && canMove("left") {
		MoveFigure("left")
		if canRotate(temp, col, row) {
			RotateFigureOld(1)
			return false
		} else {
			rollback(temp)
			return false
		}
	}

	if col == 9 && canMove("right") {
		MoveFigure("right")
		if canRotate(temp, col, row) {
			RotateFigureOld(1)
			return false
		} else {
			rollback(temp)
			return false
		}
	}

	if canRotate(temp, col, row) {
		RotateFigureOld(1)
	} else {
		rollback(temp)
	}

	//
	//if canMove("left") {
	//	MoveFigure("left")
	//	col, row = FindCenterOfFigure()
	//	if canRotate(temp, col, row) {
	//		RotateFigureOld(1)
	//		return false
	//	} else {
	//		Field = temp
	//	}
	//
	//} else if canMove("right") {
	//	MoveFigure("right")
	//	col, row = FindCenterOfFigure()
	//	if canRotate(temp, col, row) {
	//		RotateFigureOld(1)
	//		return false
	//	} else {
	//		Field = temp
	//	}
	//}
	return false
}

func TryMove(again int) bool {
	if again == 1 {
		return false
	}

	col, _ := FindCenterOfFigure()

	if col == 0 && canMove("right") {
		MoveFigure("right")
	} else if col == 9 && canMove("left") {
		MoveFigure("right")
	} else if canMove("left") {
		MoveFigure("left")
	} else if canMove("right") {
		MoveFigure("right")
	} else {
		return false
	}
	return true
}

func TryRotate() bool {
	col, row := FindCenterOfFigure()
	switch ActFigure {
	case I:
		var tempFigure [4]Cell

		if col > 7 {
			col = 7
		}

		if col == 0 {
			col = 1
		}

		var t int
		for r := row - 1; r < row+3; r++ {
			for c := col - 1; c < col+3; c++ {
				if Field[c][r].Fall == 1 {
					tempFigure[t] = Field[c][r]
					Field[c][r].Fill = 0
					Field[c][r].Fall = 0
					Field[c][r].CenterOfFigure = 0
					t++
				}
			}
		}
		t = 0

		var position string
		if tempFigure[0].Column == tempFigure[1].Column {
			position = "v" // vertical    |
		} else {
			position = "h" // horizontal  _
		}

		// пробуем повернуть
		if position == "v" {
			for c := col - 1; c < col+3; c++ {
				if Field[c][row].Fixed == tempFigure[t].Fill {
					return false
				}
				Field[c][row].Fill = tempFigure[t].Fill
				Field[c][row].Fall = tempFigure[t].Fall
				Field[c][row].CenterOfFigure = tempFigure[t].CenterOfFigure
				t++
			}
		} else {
			for r := row - 1; r < row+3; r++ {
				if Field[col][r].Fixed == tempFigure[t].Fill {
					return false
				}
				Field[col][r].Fill = tempFigure[t].Fill
				Field[col][r].Fall = tempFigure[t].Fall
				Field[col][r].CenterOfFigure = tempFigure[t].CenterOfFigure
				t++
			}
		}
		return true

	default:
		var tempFigure [3][3]Cell
		var tc, tr int
		for r := row - 1; r < row+2; r++ {
			for c := col - 1; c < col+2; c++ {
				tempFigure[tc][tr] = Field[c][r]
				Field[c][r].Fill = 0
				Field[c][r].Fall = 0
				Field[c][r].CenterOfFigure = 0
				tc++
			}
			tc = 0
			tr++
		}
		tr = 0

		// пробуем повернуть
		for c := col - 1; c < col+2; c++ {
			for r := row + 1; r > row-2; r-- {
				if Field[c][r].Fixed == 1 && tempFigure[tc][tr].Fill == 1 {
					return false
				}
				Field[c][r].Fill = tempFigure[tc][tr].Fill
				Field[c][r].Fall = tempFigure[tc][tr].Fall
				Field[c][r].CenterOfFigure = tempFigure[tc][tr].CenterOfFigure
				tc++
			}
			tc = 0
			tr++
		}
		tr = 0
		return true
	}
}

func Rotation() {

	col, _ := FindCenterOfFigure()
	var again int // сколько раз было попыток
	temp := Field

	switch ActFigure {
	case O:
		return
	case I:
		if !TryRotate() {
			rollback(temp)
		}
		return
	default:
		if col == 0 || col == 9 {
			if !TryMove(again) {
				rollback(temp)
				return
			}
			again++
			col, _ = FindCenterOfFigure()
		}
		if !TryRotate() {
			rollback(temp)
			if !TryMove(again) {
				rollback(temp)
				return
			} else {
				if !TryRotate() {
					rollback(temp)
					return
				} else {
					return
				}
			}
		}
	}
}

func RotateFigureOld(again int) {
	// need to fix unlimited rotate with tryMoveAndRotate
	if ActFigure == O {
		return // fmt.Println("ты дурак?")
	}

	col, row := FindCenterOfFigure() // return col + row

	//if ActFigure == I {
	//	if col == 0 {
	//		MoveFigure("right")
	//	} else if col == 9 {
	//		MoveFigure("left")
	//		MoveFigure("left")
	//	}
	//}

	// записываем состояние поля во времянку temp
	temp := Field

	if again == 0 {
		tryMoveAndRotate(temp)
	}

	var tc, tr int

	// пробуем повернуть
	for c := col - 1; c < col+2; c++ {
		for r := row + 1; r > row-2; r-- {
			if Field[c][r].Fixed == 1 && temp[tc][tr].Fill == 1 {
				fmt.Println("can't rotate")
				rollback(temp)               // откат
				do := tryMoveAndRotate(temp) //пробуем сместить и заново повернуть

				if !do { // если невозможно повернуть, то выходим
					fmt.Println("can't rotate ever")
					return
				}
				return //проверяем всего один раз, далее выходим
			}
			tr++
		}
		tr = 0
		tc++
	}
	tc = 0

	for c := col - 1; c < col+2; c++ {
		for r := row + 1; r > row-2; r-- {
			Field[c][r].Fill = temp[tc][tr].Fill
			Field[c][r].Fall = temp[tc][tr].Fall
			Field[c][r].CenterOfFigure = temp[tc][tr].CenterOfFigure
			tc++
		}
		tc = 0
		tr++

	}
	tr = 0

}

func canMove(dir string) bool {
	var pos [4][2]int
	var rc, rr int
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {
			if Field[c][r].Fall == 1 {
				pos[rc][rr] = Field[c][r].Column
				pos[rc][rr+1] = Field[c][r].Row
				rc++
			}
		}
	}

	max := pos[0]
	min := pos[0]
	for _, el := range pos {
		if el[0] > max[0] {
			max = el
		}

		if el[0] < min[0] {
			min = el
		}
	}

	switch dir {
	case "left":
		if min[0] == 0 {
			return false
		}
		for _, el := range pos {
			if Field[el[0]-1][el[1]].Fixed == 1 {
				return false
			}
		}
	case "right":
		if max[0] == 9 {
			return false
		}
		for _, el := range pos {
			if Field[el[0]+1][el[1]].Fixed == 1 {
				return false
			}
		}
	}
	return true
}

func MoveFigure(dir string) {
	if !canMove(dir) {
		return
	}

	switch dir {
	case "left":
		for r := 0; r < 20; r++ {
			for c := 0; c < 10; c++ {
				if Field[c][r].Fall == 1 && c != 0 {
					Field[c][r].Fill = 0
					Field[c-1][r].Fill = 1
					Field[c][r].Fall = 0
					Field[c-1][r].Fall = 1

					if Field[c][r].CenterOfFigure == 1 {
						Field[c][r].CenterOfFigure = 0
						Field[c-1][r].CenterOfFigure = 1
					}
				}
			}
		}
	case "right":
		for r := 0; r < 20; r++ {
			for c := 9; c != -1; c-- {
				if Field[c][r].Fall == 1 && c != 9 {
					Field[c][r].Fall = 0
					Field[c+1][r].Fall = 1
					Field[c][r].Fill = 0
					Field[c+1][r].Fill = 1

					if Field[c][r].CenterOfFigure == 1 {
						Field[c][r].CenterOfFigure = 0
						Field[c+1][r].CenterOfFigure = 1
					}
				}
			}
		}
	}
}

// FallFigureOnce for technical use
func FallFigureOnce(ch chan int) {

	key := <-ch
	if key != 0 {
		showFieldOnce()
		switch key {
		case 65517:
			RotateFigureOld(0)
		case 65515:
			MoveFigure("left")
			showFieldOnce()
		case 65514:
			MoveFigure("right")
			showFieldOnce()
		case 65516:
			//func fastFall()
		}
	}

	// проверка достижения нижней линии
	//if i > 15 {
	//	lowerRow := getLowerCells()
	//	if lowerRow == 19 {
	//
	//		fixFigure()
	//		fmt.Println("rich to the end of field")
	//		return
	//	}
	//}

	// проверка падения на другую фигуру
	t := findFigureCells()
	for _, row := range t {
		if Field[row[0]][row[1]+1].Fixed == 1 {
			fixFigure()
			return
		}
	}

	for r := 18; r != -1; r-- { //выяснить почему с 18 строки
		for c := 0; c < 10; c++ {
			if Field[c][r].Fall == 1 {
				if Field[c][r].CenterOfFigure == 1 {
					Field[c][r].CenterOfFigure = 0
					Field[c][r+1].CenterOfFigure = 1
				}
				Field[c][r].Fill = 0
				Field[c][r+1].Fill = 1

				Field[c][r].Fall = 0
				Field[c][r+1].Fall = 1

			}
		}

	}
	showField()

	time.Sleep(5 * time.Millisecond)
}

func fastFall() {
	/*
		1. найти все нижние ячейки
		2. найти минимальное расстояние от одной из нижних ячеек до ближайшей fix
		3. провести разово нужное количество итераций (2) /// или сразу прописать фигуру снизу без лишних итераций

		---------

		1.
	*/
}

func FallFigure(ch chan int) {
	for i := 0; i < 19; i++ {

		key := <-ch
		if key != 0 {
			showFieldOnce()
			switch key {
			case 65517:
				RotateFigureOld(0)
			case 65515:
				MoveFigure("left")
				showFieldOnce()
			case 65514:
				MoveFigure("right")
				showFieldOnce()
			case 65516:
				fastFall()
			}
		}

		// проверка достижения нижней линии
		if i > 15 {
			lowerRow := getLowerCells()
			if lowerRow == 19 {

				fixFigure()
				fmt.Println("rich to the end of field")
				return
			}
		}

		// проверка падения на другую фигуру
		t := findFigureCells()
		for _, row := range t {
			if Field[row[0]][row[1]+1].Fixed == 1 {
				fixFigure()
				return
			}
		}

		for r := 18; r != -1; r-- { //выяснить почему с 18 строки
			for c := 0; c < 10; c++ {
				if Field[c][r].Fall == 1 {
					if Field[c][r].CenterOfFigure == 1 {
						Field[c][r].CenterOfFigure = 0
						Field[c][r+1].CenterOfFigure = 1
					}
					Field[c][r].Fill = 0
					Field[c][r+1].Fill = 1

					Field[c][r].Fall = 0
					Field[c][r+1].Fall = 1

				}
			}

		}
		showField()

		time.Sleep(5 * time.Millisecond)
	}

}

func startGame(ch chan int) {
	CreateField(&Field)
	CreateFigure()

	showField()

	SpawnAdvancedFigure(I, 0)
	showField()
	Rotation()
	showField()
	MoveFigure("left")
	showField()

	Rotation()
	showField()

	fmt.Println("simulation is finished")
}

func main() {

	ch := make(chan int)
	go startGame(ch)
	for {
		s, _ := getKeyTimeout(50 * time.Millisecond)

		if s != 0 {
			ch <- int(s)
		} else {
			ch <- 0
		}
	}
}

func getKeyTimeout(tm time.Duration) (ch rune, err error) {
	if err = keyboard.Open(); err != nil {
		return
	}
	defer keyboard.Close()

	var (
		chChan  = make(chan rune, 1)
		errChan = make(chan error, 1)

		timer = time.NewTimer(tm)
	)
	defer timer.Stop()

	go func(chChan chan<- rune, errChan chan<- error) {
		_, s, err := keyboard.GetSingleKey()
		if err != nil {
			errChan <- err
			return
		}
		chChan <- rune(s)
	}(chChan, errChan)

	select {
	case <-timer.C:
		return
	case ch = <-chChan:
	case err = <-errChan:
	}

	return
}
