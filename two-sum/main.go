package twosum

func twoSum(nums []int, target int) []int {
	var sum []int
	for i1, v1 := range nums {
		for i2, v2 := range nums {
			if i1 != i2 {
				if v1+v2 == target {
					sum = []int{i1, i2}
				}
			}

		}
	}
	return sum
}
