package main

func twoSum(nums []int, target int) []int {
	notes := make(map[int]int)

	for i, num := range nums {
		deferent := target - num

		if indexDeferent, found := notes[deferent]; found {
			return []int{indexDeferent, i}
		}
		notes[num] = i
	}
	return nil
}