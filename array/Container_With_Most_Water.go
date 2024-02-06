package array

// 11. Container With Most Water
// 11. 盛最多水的容器
//
// 双指针法: 面积只有两个变量，i和j相向时，这个变量在减小。点到为止
func maxArea(height []int) int {
	ret := 0
	i, j := 0, len(height)-1
	for i < j {
		if height[i] < height[j] {
			ret = max(ret, height[i]*(j-i))
			i++
		} else {
			ret = max(ret, height[j]*(j-i))
			j--
		}
	}
	return ret
}
