// https://leetcode.com/contest/warm-up-contest/problems/lexicographical-numbers/

package main

import (
	"fmt"
	"os"
)

type state struct {
	facing      rune
	next_coords [2]float64
	now_coords  [2]float64
	turn        rune
}

func next_state(position state, rune_ip rune) state {
	if position.facing == 'N' {
		if rune_ip == 'G'{
			position.turn = 'G'
			position.facing = 'N'
			position.next_coords = position.now_coords
		}else if rune_ip = 'L'{
			position.turn = 'L'
			position.facing = 'W'
			position.next_coords = position.now_coords
		} else if rune_ip = 'R'{
			position.turn = 'R'
			position.facing = 'E'
			position.next_coords = position.now_coords
		}
	}


}

func main() {
	errands := "GGLLRRG"
	isRobotBounded(errands)
}

func isRobotBounded(instructions string) bool {

	return true

}

/*
func move(s []string) []int {
	if char == 'N' {
		if movement == 'G' {

		}

	}
}
*/
