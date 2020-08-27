package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, d := range nums {
		another := target - d
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[d] = i
	}

	return nil
}
