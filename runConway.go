package main

import (
	"./conway"
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [*.cw] [Step limit (-1 for none)]\n", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	game := new(conway.Game)
	err := game.FromFile(os.Args[1])

	if err != nil {
		fmt.Printf("Map failed to load, Error: %s\n", err)
	}

	stepLimit, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Invalid step limit, Error: %s", err)
	}

	game.Print()
	fmt.Println("Rows:%s", game.Rows)
	fmt.Println("Cols:%s", game.Cols)
	for i := 0; i < int(stepLimit); i++ {
		time.Sleep(200 * time.Millisecond)
		game.Step()
		game.Print()
	}
}
