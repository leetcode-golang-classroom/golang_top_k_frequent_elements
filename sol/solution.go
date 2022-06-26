package sol

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num] += 1
	}
	nLen := len(nums)
	bucket := make([][]int, nLen)
	for key, value := range freq {
		bucket[value-1] = append(bucket[value-1], key)
	}
	result := []int{}
	for n := len(bucket) - 1; n >= 0; n-- {
		list := bucket[n]
		if len(list) > 0 {
			bLen := len(list)
			for idx := 0; idx < bLen; idx++ {
				result = append(result, list[idx])
				if len(result) == k {
					return result
				}
			}
		}
	}
	return result
}
