package main

import (
	"fmt"
)

func main() {
	maxRedigit(1000)
	maxRedigit(0)
	maxRedigit(123)
	maxRedigit(879)
}

func maxRedigit(num int) {
	// Check if its 3 digit positive int, if true, rearrange the digit from the biggest num
	if num >= 100 && num <= 999 {
		nums := make([]int, 0, 3)
		for num > 0 {
			nums = append(nums, num%10) // Get each digit with mod & push the remainder into the slice e.g = 792 % 10 = 2
			num /= 10                   // e.g = 792 / 10 = 79, 79 / 10 = 7
		}

		sortedNums := sortNums(nums)
		fmt.Printf("%d\n", sortedNums)
	} else {
		fmt.Println(nil)
	}
}

// Sorting by comparing 2 int and swapping them
func sortNums(nums []int) int {
	var N = len(nums)
	var x = 0
	var y = 1
	for y < N {
		var firstNumber = nums[x]
		var secondNumber = nums[y]
		if firstNumber < secondNumber {
			nums[x] = secondNumber
			nums[y] = firstNumber
		}
		x++
		y++
	}
	// Turn back each 3 digit into one integer
	return nums[0]*100 + nums[1]*10 + nums[2]
}
