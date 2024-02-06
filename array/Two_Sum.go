package array

func twoSum(nums []int, target int) []int {
	n := len(nums)
	mp := make(map[int]int, n)
	for i, v := range nums {
		if idx, ok := mp[target-v]; ok {
			return []int{idx, i}
		} else {
			mp[v] = i
		}
	}
	return nil
}
