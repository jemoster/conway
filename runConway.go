package main

import (
	"./conway"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [*.cw]\n", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	game := new(conway.Game)
	err := game.FromFile(os.Args[1])

	if err != nil {
		fmt.Printf("Map failed to load, Error: %s\n", err)
	}

	game.Print()
	for i:=0; i<10; i++ {
		fmt.Println("----------")
		game.Step();
		game.Print()
	}
}
