/*
TODO:
1. (done) Взаимодействие с фигурой I
2. (done) Реализовать быстрое падение фигуры
3. (done) Доработать спавн
4. Добавить фантом?
5. Тестирование
6. Оптимизация
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
			//fmt.Printf("%d-%d|%d", Field[c][r].Column, Field[c][r].Row, Field[c][r].Fall)
			//fmt.Print(Field[c][r].Fall)
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
	time.Sleep(50 * time.Millisecond)
	//}
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

var ActFigure Figure

func SpawnFigureAdvanced(a Figure, s int) {
	ActFigure = a //getRandFigure()
	spawnCol := s //:= getRand(8)
	temp := Field

	switch ActFigure {
	case O:
		spawnCol = s //getRand(9)
	case I:
		spawnCol = s //getRand(7)
	}

	fc := 0
	fr := 1
	for r := 0; r < len(ActFigure)-1; r++ {
		for c := spawnCol; c < len(ActFigure)+spawnCol; c++ {
			if ActFigure[fr][fc].Fill == 1 && Field[c][r].Fixed == 1 {
				rollback(temp)
				return
			}

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

func SpawnFigure() bool {
	temp := Field
	ActFigure = getRandFigure()
	var spawnCol int

	switch ActFigure {
	case O:
		spawnCol = getRand(9)
	case I:
		spawnCol = getRand(7)
	default:
		spawnCol = getRand(8)
	}

	fc := 0
	fr := 1
	for r := 0; r < len(ActFigure)-1; r++ {
		for c := spawnCol; c < len(ActFigure)+spawnCol; c++ {
			if ActFigure[fr][fc].Fill == 1 && Field[c][r].Fixed == 1 {
				rollback(temp)
				return true
			}

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
	return false
}

func getBottomCells() int {
	for r := 19; r > 0; r-- {
		for c := 0; c < 10; c++ {
			if Field[c][r].Fall == 1 {
				return r
			}
		}
	}
	return 0
}

func getTopCells() int {
	for r := 0; r < 20; r++ {
		for c := 0; c < 10; c++ {
			if Field[c][r].Fall == 1 {
				return r
			}
		}
	}
	return 0
}

func clearLine(row int) {
	for c := 0; c < 10; c++ {
		Field[c][row].Fill = 0
		Field[c][row].Fixed = 0
		Field[c][row].CenterOfFigure = 0
	}
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
			temp = 0
			r++
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

type pos struct {
	col int
	row int
}

type figureCells struct {
	left   pos
	right  pos
	bottom pos
	top    pos
	center pos
	all    [4]pos
}

// findFigureCells returns 'position' struct, where have boundary values of figure cells
func findFigureCells() figureCells {
	var position figureCells
	var x int
	for r := 19; r > -1; r-- {
		for c := 0; c < 10; c++ {
			if Field[c][r].Fall == 1 {
				position.all[x].col = Field[c][r].Column
				position.all[x].row = Field[c][r].Row
				x++
			}
			if Field[c][r].CenterOfFigure == 1 {
				position.center.col = c
				position.center.row = r
			}
		}
	}

	position.left = position.all[0]
	position.right = position.all[0]
	for _, el := range position.all {
		if el.col < position.left.col {
			position.left.col = el.col
			position.left.row = el.row
		}
		if el.col > position.right.col {
			position.right.col = el.col
			position.right.row = el.row
		}
	}

	position.bottom = position.all[0]
	position.top = position.all[0]
	for _, el := range position.all {
		if el.row > position.bottom.row {
			position.bottom.col = el.col
			position.bottom.row = el.row
		}
		if el.row < position.top.row {
			position.top.col = el.col
			position.top.row = el.row
		}
	}

	return position
}

func rollback(temp field) {
	Field = temp
}

func TryMove(again int) bool {
	if again == 1 {
		return false
	}

	p := findFigureCells()

	if p.right.col == 1 && canMove(p, "right") {
		MoveFigure("right")
	} else if p.left.col == 8 && canMove(p, "left") {
		MoveFigure("left")
	} else if canMove(p, "left") {
		MoveFigure("left")
	} else if canMove(p, "right") {
		MoveFigure("right")
	} else {
		return false
	}
	return true
}

func TryRotate() bool {
	p := findFigureCells()
	switch ActFigure {
	case I:
		var tempFigure [4]Cell

		if p.center.col > 7 {
			p.center.col = 7
		}

		if p.center.col == 0 {
			p.center.col = 1
		}

		var t int
		for r := p.center.row - 1; r < p.center.row+3; r++ {
			for c := p.center.col - 1; c < p.center.col+3; c++ {
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

		var dir string
		if tempFigure[0].Column == tempFigure[1].Column {
			dir = "v" // vertical    |
		} else {
			dir = "h" // horizontal  _
		}

		// пробуем повернуть
		if dir == "v" {
			for c := p.center.col - 1; c < p.center.col+3; c++ {
				if Field[c][p.center.row].Fixed == tempFigure[t].Fill {
					return false
				}
				Field[c][p.center.row].Fill = tempFigure[t].Fill
				Field[c][p.center.row].Fall = tempFigure[t].Fall
				Field[c][p.center.row].CenterOfFigure = tempFigure[t].CenterOfFigure
				t++
			}
		} else {
			for r := p.center.row - 1; r < p.center.row+3; r++ {
				if Field[p.center.col][r].Fixed == tempFigure[t].Fill {
					return false
				}
				Field[p.center.col][r].Fill = tempFigure[t].Fill
				Field[p.center.col][r].Fall = tempFigure[t].Fall
				Field[p.center.col][r].CenterOfFigure = tempFigure[t].CenterOfFigure
				t++
			}
		}
		return true

	default:
		var tempFigure [3][3]Cell
		var tc, tr int
		for r := p.center.row - 1; r < p.center.row+2; r++ {
			for c := p.center.col - 1; c < p.center.col+2; c++ {
				if Field[c][r].Fall == 1 {
					tempFigure[tc][tr].Fall = 1
					tempFigure[tc][tr].Fill = 1
					Field[c][r].Fill = 0
					Field[c][r].Fall = 0

					if Field[c][r].CenterOfFigure == 1 {
						tempFigure[tc][tr].CenterOfFigure = 1
						Field[c][r].CenterOfFigure = 0
					}
				}
				tc++
			}
			tc = 0
			tr++
		}
		tr = 0
		// пробуем повернуть
		for c := p.center.col - 1; c < p.center.col+2; c++ {
			for r := p.center.row + 1; r > p.center.row-2; r-- {
				if Field[c][r].Fixed == 1 && tempFigure[tc][tr].Fall == 1 {
					return false
				}
				if Field[c][r].Fixed == 0 {
					Field[c][r].Fill = tempFigure[tc][tr].Fill
					Field[c][r].Fall = tempFigure[tc][tr].Fall
					Field[c][r].CenterOfFigure = tempFigure[tc][tr].CenterOfFigure
				}
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
	p := findFigureCells()

	var again int // сколько раз было попыток
	temp := Field

	switch ActFigure {
	case O:
		return
	case I:
		if p.center.row == 0 || p.center.row > 17 {
			return
		}
		if !TryRotate() {
			rollback(temp)
		}
		return
	default:
		if p.center.row == 0 || p.center.row == 19 {
			return
		}

		if p.center.col == 0 || p.center.col == 9 {
			if !TryMove(again) {
				rollback(temp)
				return
			}
			again++
			p = findFigureCells()
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

func canMove(cells figureCells, dir string) bool {
	switch dir {
	case "left":
		// упор влево
		if cells.left.col == 0 {
			return false
		}
		// слева фикс для каждой ячейки
		for _, c := range cells.all {
			if Field[c.col-1][c.row].Fixed == 1 {
				return false
			}
		}
	case "right":
		// упор вплаво
		if cells.right.col == 9 {
			return false
		}
		// слева фикс для каждой ячейки
		for _, c := range cells.all {
			if Field[c.col+1][c.row].Fixed == 1 {
				return false
			}
		}
	}
	return true
}

func MoveFigure(dir string) {
	if !canMove(findFigureCells(), dir) {
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

func fastFall() {
	cells := findFigureCells()

	defer showFieldOnce()

	for i := cells.bottom.row; i < 19; i++ {
		cells = findFigureCells()
		// проверка достижения нижней линии
		if cells.bottom.row == 19 {
			return
		}

		// проверка падения на другую фигуру
		for _, c := range cells.all {
			if Field[c.col][c.row+1].Fixed == 1 {
				fixFigure()
				return
			}
		}

		for r := 18; r != -1; r-- {
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
	}
}

// FallFigureOnce technical use only
func FallFigureOnce() {
	// проверка достижения нижней линии
	cells := findFigureCells()
	if cells.bottom.row == 19 {
		fixFigure()
		fmt.Println("rich to the end of field")
		return
	}

	// проверка падения на другую фигуру
	for _, c := range cells.all {
		if Field[c.col][c.row+1].Fixed == 1 {
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

func FallFigure() {
	for i := 0; i < 19; i++ {
		switch a {
		case "next": // only after FastFall
			time.Sleep(50 * time.Millisecond)
			a = "run"
			return
		case "pause": // after move and rotate
			for a != "run" {
				time.Sleep(10 * time.Millisecond)
			}
		}

		cells := findFigureCells()
		showFieldOnce()
		// проверка достижения нижней линии
		if i > 15 {
			if cells.bottom.row == 19 {
				fixFigure()
				showFieldOnce()
				return
			}
		}

		// проверка падения на другую фигуру
		for _, c := range cells.all {
			if Field[c.col][c.row+1].Fixed == 1 {
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
		showFieldOnce()
		time.Sleep(500 * time.Millisecond)
	}
	fixFigure()
}

func gameProcess() {
	// нужен лок действия, когда нет падающей фигуры
	var score int
	exit := false
	for exit == false {
		exit = SpawnFigure()
		FallFigure()
		score++
	}
	fmt.Println("Game over")
	fmt.Printf("Your score: %d\n", score)

	fmt.Println("simulation is finished")
}

func doActions(ch chan int, chA chan string) {
	for {
		key := <-ch
		if key != 0 {
			switch key {
			case 65517:
				a = "pause"
				Rotation()
				showFieldOnce()
				a = "run"
			case 65515:
				a = "pause"
				MoveFigure("left")
				showFieldOnce()
				a = "run"
			case 65514:
				a = "pause"
				MoveFigure("right")
				showFieldOnce()
				a = "run"
			case 65516:
				a = "next"
				fastFall()
				fixFigure()
			case 27:
				ForceExit()
			}

		}
	}
}

var a string

func main() {
	CreateField(&Field)
	CreateFigure()
	ch := make(chan int)
	chA := make(chan string)
	go gameProcess()
	go getKey(ch)
	go doActions(ch, chA)
	go func() {
		time.Sleep(300 * time.Second)
		os.Exit(1)
	}()
	for range chA {
		a = <-chA
	}

}

func ForceExit() {
	os.Exit(1)
}

func getKey(ch chan int) {
	for {
		s, _ := getKeyTimeout(50 * time.Millisecond)
		time.Sleep(20 * time.Millisecond)
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
