package main
import "fmt"

func main() {
	input := []int{1, 3, 5, 7, 9}
	k :=2 

	// real roate
	k = k % len(input)

	// start 1
	// roate k
	for i:=1; i<=k; i++ {
		input = rotateOnce(input)
	}

	fmt.Println(input)
}


// rotate 1 left 
func rotateOnce(input []int) []int {
	// save 1st
	first := input[0]

	// i starts at 1, keep len same 
	for i:=1; i<len(input); i++ {
		input[i-1] = input[i]		
	}

	input[len(input)-1] = first 
		
	return input
}
