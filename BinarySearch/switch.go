package main

import (
	"fmt"
)

var guess, a int //between 1 and 300
var v, mid int

/*Voici une description en pseudo-code de la recherche dichotomique :
1 - Let min = 1 and max = n.
2 - Guess the average of max and min , rounded down so that it is an integer.
3 - If you guessed the number, stop. You found it!
4 - If the guess was too low, set min to be one larger than the guess.
5 - If the guess was too high, set max to be one smaller than the guess.
6 - Go back to step two.*/
/*guess number is 30 and not more than 9 times*/
func main() {
	guess = 30
	min := 1
	max := 300
	time := 9
	a := guesit(min, guess, max, time)
	fmt.Println(a)

}

func guesit(min int, guess int, max int, time int) int {
	for v < time {
		mid = ((min + max) / 2)
		v++
		n := mid != guess
		switch n {
		case mid < guess:
			min = mid
			fmt.Println(min, guess, max, time)
			return guesit(min, guess, max, time)
			fmt.Println("You didn't found it!")

		case mid > guess:
			max = mid
			fmt.Println(min, guess, max, time)
			return guesit(min, guess, max, time)
			fmt.Println("You didn't found it!")

		case mid == guess:
			if mid == guess {
				fmt.Println("You found it!")
				break
			}
		}
	}
	return mid
}
