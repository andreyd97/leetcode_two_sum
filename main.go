package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var final []int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				final = append(final, i)
				final = append(final, j)
				return final
			}

		}

	}
	return nil
}

func main() {
	a := []int{3, 2, 4}
	catch := 6
	fmt.Println(twoSum(a, catch))
}

