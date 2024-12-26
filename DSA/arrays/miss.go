package main

import "sort"

func missingNumber(nums []int) int {

	n := len(nums)
	sort.Ints(nums);
	testSum := (n * (n + 1) * 2) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return testSum - sum
}
// more optimised can be done via sorting
func missingNumberOpti(nums []int) int {
	sort.Ints(nums)
	if(nums[0] !=0){
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[i-1] + 1{
			return nums[i-1] - 1
		}
	}
	return len(nums)
}


/*
nums = [3,0,1]
 ans = 2

 n * (n+1) /2
 3 * 4 /2 = 6

 3 -1 = 2 
*/
