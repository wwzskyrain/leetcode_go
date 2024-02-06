package array

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums { // 这里要用元素值，一定不能省略_,哈，哈哈
		set[num] = true
	}
	longestStreak := 0
	for num := range set {
		if !set[num-1] {
			//这里可不是map的contains的写法，这里还是取值，不记得了，在上面map的value就是bool类型呀，小子
			//还要注意算法，这里只走一个方向就好了
			curNum := num
			curStreak := 1
			for set[curNum+1] {
				curNum++
				curStreak++
			}
			longestStreak = max(longestStreak, curStreak)
		}
	}
	return longestStreak
}
