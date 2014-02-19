package conway

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Game struct {
	Map  []byte
	Rows int
	Cols int
}

func (game *Game) index(row int, col int) (i int) {
	if row < 0 {
		row = game.Rows - 1
	}
	if col < 0 {
		col = game.Cols - 1
	}
	if row > game.Rows-1 {
		row = 0
	}
	if col > game.Cols-1 {
		col = 0
	}

	i = game.Cols*row + col
	return i
}

func (game *Game) Print() {
	fmt.Printf("\033[1;1H")
	for row := 0; row < game.Rows; row++ {
		for col := 0; col < game.Cols; col++ {
			fmt.Printf("%c", game.Map[game.index(row, col)])
		}
		fmt.Printf("\n")
	}
}

func (game *Game) Step() {
	for row := 0; row < game.Rows; row++ {
		go func(){for col := 0; col < game.Cols; col++ {
			game.StepCell(row,col)
		}
		}()
	}
	for row := 0; row < game.Rows; row++ {
		go func(){for col := 0; col < game.Cols; col++ {
			game.UpdateCell(row,col)
		}
		}()
	}
}



func (game *Game) UpdateCell(row int, col int) {
	if game.Map[game.index(row, col)] == 'd' {
		game.Map[game.index(row, col)] = '_'
	} else if game.Map[game.index(row, col)] == 'n' {
		game.Map[game.index(row, col)] = 'r'
	}
}

func (game *Game) StepCell(row int, col int) {

	//Get live neighbors
	live := 0
	live += isLive(game.Map[game.index(row-1, col-1)])
	live += isLive(game.Map[game.index(row-1, col)])
	live += isLive(game.Map[game.index(row-1, col+1)])
	live += isLive(game.Map[game.index(row, col-1)])
	live += isLive(game.Map[game.index(row, col+1)])
	live += isLive(game.Map[game.index(row+1, col-1)])
	live += isLive(game.Map[game.index(row+1, col)])
	live += isLive(game.Map[game.index(row+1, col+1)])

	//Apply rules
	if isLive(game.Map[game.index(row, col)]) == 1 {
		if live < 2 || live > 3 {
			game.Map[game.index(row, col)] = 'd'
		}

	} else if live == 3 && game.Map[game.index(row, col)] != 'r' {
		game.Map[game.index(row, col)] = 'n'
	}
	//game.Map[game.index(row,col)] = 48+byte(live)
}

func isLive(cell byte) (live int) {
	live = 0
	if cell == 'r' || cell == 'd' {
		live = 1
	}
	return
}

func (game *Game) FromFile(name string) (err error) {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	fileinfo, err := file.Stat()

	r := bufio.NewReaderSize(file, int(fileinfo.Size()))

	game.Load(r)

	return err
}

func (game *Game) Load(r *bufio.Reader) (err error) {
	game.Rows = 0
	game.Cols = 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		game.Map = append(game.Map, line[0:len(line)-1]...)
		game.Rows++
		game.Cols = len(line) - 1
	}
	//fmt.Println(game.Map)
	return
}
