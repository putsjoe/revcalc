// revcalc

package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
	fmt.Print("")
	var input string
	var stack int

	for input != "exit" && input != "e" {

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		split := strings.Fields(text)

		// Take the first value, second and the operator from third entry.

		if len(split) == 3 {

			var first, _ = strconv.Atoi(split[0])
			var sec, _ = strconv.Atoi(split[1])

			switch split[2] {
			case "*":
				var stk int = first * sec
				stack = stack + stk
			case "/":
				var stk int = first / sec
				stack = stack + stk
			case "-":
				var stk int = first - sec
				stack = stack + stk
			case "+":
				var stk int = first + sec
				stack = (stack + stk)

			default:
				fmt.Println("Unknown")

			} // end switch

		} //end for

		if len(split) != 0 {
			// Turn it into a case statement
			if split[0] == "p" {
				fmt.Println(stack)
			}
			if split[0] == "c" {
				stack = 0
			}

		}

	}

	/*
	   fmt.Scanln(&input) << Think the issue at the moment is this line.
	   The doc says that it takes each string seperated by space is taken as an argument.
	*/
}
