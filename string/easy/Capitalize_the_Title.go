package easy

import "strings"

// 2129. Capitalize the Title
func capitalizeTitle(title string) string {
	words := strings.Fields(title) //按照空格分开了
	for i := range words {
		words[i] = strings.ToLower(words[i]) //都小写
		if len(words[i]) > 2 {
			words[i] = strings.Title(words[i]) //还有title这种函数库？怪不得deprecated呢
		}
	}
	return strings.Join(words, " ")
}
