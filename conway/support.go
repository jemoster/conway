package conway

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Game struct {
	Map []byte
	Rows int
	Cols int
}

func  (game *Game) index(row int, col int) (i int){
	if(row<0){ row = game.Rows-1 }
	if(col<0){ col = game.Cols-1 }
	if(row>game.Rows-1){ row = 0 }
	if(col>game.Cols-1){ col = 0 }

	i = game.Rows*row + col
	return i
}

func (game *Game) Print() {
	for i := 0; i < game.Cols; i++ {
		for j := 0; j < game.Rows; j++ {
			fmt.Printf("%c", game.Map[i*game.Cols+j])
		}
		fmt.Printf("\n")
	}
}

func (game *Game) Step() {
	for row:=0; row<game.Rows; row++ {
		for col:=0; col<game.Cols; col++{
			//Get live neighbors
			live := 0;
			live += isLive(game.Map[game.index(row-1,col-1)])
			live += isLive(game.Map[game.index(row-1,col)])
			live += isLive(game.Map[game.index(row-1,col+1)])
			live += isLive(game.Map[game.index(row,col-1)])
			live += isLive(game.Map[game.index(row,col-1)])
			live += isLive(game.Map[game.index(row+1,col-1)])
			live += isLive(game.Map[game.index(row+1,col)])
			live += isLive(game.Map[game.index(row+1,col+1)])
			
			//Apply rules
			if(isLive(game.Map[game.index(row,col)])==1){
				if(live<2 || live>3){
					game.Map[game.index(row,col)] = 'd'
				}
				
			} else if(live==3) {
				game.Map[game.index(row,col)] = 'n'	
			}
		}
	}
	for row:=0; row<game.Rows; row++ {
		for col:=0; col<game.Cols; col++{
			if(game.Map[game.index(row,col)]=='d') {
				game.Map[game.index(row,col)] = '_'
			} else if (game.Map[game.index(row,col)]=='n'){
				game.Map[game.index(row,col)] = 'r'
			}
		}
	}
}
func isLive(cell byte) (live int) {
	live = 0
	if(cell=='r' || cell=='d'){
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
	game.Rows = 0;
	game.Cols = 0;
	for ;; {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		game.Map = append(game.Map,line[0:len(line)-1]...)
		game.Rows++
		game.Cols = len(line)-1
	}
	return
}
