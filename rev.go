// revcalc

package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func print(a int) {

	fmt.Print(" ")
	fmt.Print(a)
	fmt.Println("")

}

func handle_stack(a string, stack int) int {

	switch a {

	case "p":
		print(stack)
	case "c":
		stack = 0
	}
	return stack

}

func main() {
	fmt.Println("Reverse Polish Calculator with Stack:: ")
	fmt.Print("")
	var input string
	var stack int

	for input != "exit" && input != "e" {

		fmt.Print(" ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		split := strings.Fields(text)

		// Take the first value, second and the operator from third entry.

		if len(split) >= 3 {

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

			if len(split) == 4 {
				stack = handle_stack(split[3], stack)
			}

		} //end if

		if len(split) != 0 {
			stack = handle_stack(split[0], stack)
		}

	} // end for

}
