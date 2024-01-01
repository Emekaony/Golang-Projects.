package filter

import "fmt"

func filter(pred func(int) bool, values []int) []int {
	answer := []int{}
	for _, value := range values {
		if pred(value) == true {
			answer = append(answer, value)
		}
	}
	return answer
}

func isEven(i int) bool {
	return i%2 == 0
}

func isOdd(i int) bool {
	return i%2 == 1
}

func Run() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Even values: ", filter(isEven, values))
	fmt.Println("Odd values: ", filter(isOdd, values))
}
