package hard

// 42. Trapping Rain Water
func trap(height []int) (ans int) {
	length := len(height)
	if length < 3 {
		return 0
	}
	lMax, rMax := height[0], height[length-1]
	l, r := 0, length-1
	ret := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > lMax {
				lMax = height[l]
			} else {
				ret += lMax - height[l]
			}
			l++
		} else {
			if height[r] > rMax {
				rMax = height[r]
			} else {
				ret += rMax - height[r]
			}
			r--
		}
	}
	return ret
}
