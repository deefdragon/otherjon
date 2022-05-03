package main

import (
	"fmt"

	"github.com/deefdragon/otherjon/player"
)

func main() {
	p1 := player.New("Xanathar")
	fmt.Println(p1.GetStatus())

	//simulate a mock battle using the player, printing the stats at the beginning and the end.

	//Write tests for EVERY method. Aim for 100% of statements covered.
}
