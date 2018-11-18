package main
import "fmt"

func main() {

	// input
	input := []int{1, 3, 5, 7, 9}

	// left rotate once
	k :=1 

	arr := rotate(input, k)
	
	fmt.Println(arr)
}


// rotate left much easier, just mod = k % len, new index = (i+mod) % len 
func rotate(input []int, k int) []int {
	var arr []int

	// the real move 
	mod := k % len(input)

	// we loop original index
	for i:=0; i<len(input); i++ {
		// NOTE: the core idea is we keep origial index
		// then arr[original index + move] == the element we need
		// +mod is good, mod can do  
		index := (i+mod) % len(input) 

		arr = append(arr, input[index])
	}

	return arr 
}
