package medium

import "testing"

func TestFrequencyTracker(t *testing.T) {
	obj := Constructor()
	obj.Add(3)
	obj.DeleteOne(3)
	param_3 := obj.HasFrequency(2)
	println(param_3)
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
