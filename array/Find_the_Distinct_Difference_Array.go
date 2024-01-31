package array

// 2670. Find the Distinct Difference Array
// 第一次，独立的code，文不加点，没有语法错误，跑过所有case
func distinctDifferenceArray(nums []int) []int {
	ans := make([]int, 0)
	rMap := make(map[int]int)
	for _, v := range nums {
		rMap[v]++
	}
	lMap := make(map[int]int)
	for _, v := range nums {
		rMap[v]--
		if rMap[v] == 0 {
			// 这个是核心之一，删除map中的key
			delete(rMap, v)
		}
		lMap[v]++
		ans = append(ans, len(lMap)-len(rMap))
	}
	return ans
}

// 乐扣官网的写法，好丑啊
func distinctDifferenceArray1(nums []int) []int {
	st := map[int]struct{}{} //这个map的v竟然是一个空struct，就完全是一个站位吧。
	sufCnt := make([]int, len(nums)+1)
	for i := len(nums) - 1; i > 0; i-- {
		st[nums[i]] = struct{}{}
		sufCnt[i] = len(st)
	}
	var res []int
	st = map[int]struct{}{} //st这个变量有用了，上面把“右边不同的数的个数”已经存放到sufCnt了，所以st就没用了。
	for i := 0; i < len(nums); i++ {
		st[nums[i]] = struct{}{}
		res = append(res, len(st)-sufCnt[i+1])
	}
	return res
}
