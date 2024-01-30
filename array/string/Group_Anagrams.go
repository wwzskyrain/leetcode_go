package string

import "sort"

// 49. Group Anagrams
// https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{} //map的初始化方式（非make版本）
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j] //s[i]就是字符的byte值，就是ASCII的值
		})
		sortStr := string(s) //注意这里string与[]byte之间的转换
		mp[sortStr] = append(mp[sortStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		// v就是sring切片，即[]string
		ans = append(ans, v)
	}
	return ans
}
