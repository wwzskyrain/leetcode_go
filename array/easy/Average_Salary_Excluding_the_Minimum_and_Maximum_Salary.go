package easy

// 1491. Average Salary Excluding the Minimum and Maximum Salary
func average(salary []int) float64 {
	minS, maxS, sum := salary[0], salary[0], 0
	for _, s := range salary {
		if s > maxS {
			maxS = s
		}
		if s < minS {
			minS = s
		}
		sum += s
	}
	sum -= maxS + minS
	return float64(sum) / float64(len(salary)-2)
}
