package main

import (
	"fmt"
	"strings"
	"strconv"
)

func contains(s []int, e interface{}) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func check(i Player) bool {
	switch {
	case contains(i.Place, 1) && contains(i.Place, 2) && contains(i.Place, 3): // ---
		return true
	case contains(i.Place, 4) && contains(i.Place, 5) && contains(i.Place, 6): // ---
		return true
	case contains(i.Place, 7) && contains(i.Place, 8) && contains(i.Place, 9): // ---
		return true
	case contains(i.Place, 1) && contains(i.Place, 4) && contains(i.Place, 7): // |||
		return true
	case contains(i.Place, 2) && contains(i.Place, 5) && contains(i.Place, 8): // |||
		return true
	case contains(i.Place, 3) && contains(i.Place, 6) && contains(i.Place, 9): // |||
		return true
	case contains(i.Place, 1) && contains(i.Place, 5) && contains(i.Place, 9): // \\\
		return true
	case contains(i.Place, 3) && contains(i.Place, 5) && contains(i.Place, 7): // ///
		return true
	}
	return false
}


var Board = 
`
|1|2|3|
|-----|
|4|5|6|
|-----|
|7|8|9|
`

type Player struct {
	Symbol string
	Place []int
	Win bool
}

func main() {
	x := Player{Symbol: "x"}
	o := Player{Symbol: "o"}

	var response int

	players := []Player{x,o}

	var win bool
	fmt.Printf(Board)
	
	for i := 0; !win; i++ {
		y := 0
		if(i == 2) {
			y = 1
			i = 0
		}
		fmt.Printf("Player %v\n", players[i].Symbol)
		fmt.Scanln(&response)

		if response == 0 || response > 9 {
			fmt.Println("Too big or too small number or string")
			i = y
			fmt.Printf(Board)
			continue
		}

		if contains(players[0].Place, response) || contains(players[1].Place, response) {
			fmt.Println("This place is taken")
			i = y
			fmt.Printf(Board)
			continue
		}

		Board = strings.ReplaceAll(Board, strconv.Itoa(response), players[i].Symbol)
		fmt.Printf(Board)

		players[i].Place = append(players[i].Place, response)

		win = check(players[i])
		players[i].Win = win

		if (len(players[0].Place) >= 4 && len(players[1].Place) >= 4 && (len(players[0].Place) < len(players[1].Place) || len(players[0].Place) > len(players[1].Place))) {
			fmt.Printf("Tie!")
			return
		}
	}
	
	for _, j := range players {
		if(j.Win) {
			fmt.Printf("GG %v, You win!", j.Symbol)
		}
	}
}

