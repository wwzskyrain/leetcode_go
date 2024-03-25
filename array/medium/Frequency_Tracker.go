package medium

// 2671. Frequency Tracker
type FrequencyTracker struct {
	counter      map[int]int
	frequencyMap map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		counter:      make(map[int]int),
		frequencyMap: make(map[int]int),
	}
}

// 1. 要用map的默认值
// 2.map的值可以直接操作自加自减
func (this *FrequencyTracker) Add(number int) {
	c, _ := this.counter[number]
	this.counter[number]++
	this.frequencyMap[c]-- //尽情的
	this.frequencyMap[c+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	v := this.counter[number]
	if v == 0 {
		return
	}
	this.counter[number]--
	this.frequencyMap[v]--
	this.frequencyMap[v-1]++

}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	v, _ := this.frequencyMap[frequency]
	return v > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
