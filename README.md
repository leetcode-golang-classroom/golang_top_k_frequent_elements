# golang_top_k_frequent_elements

Given an integer array `nums` and an integer `k`, return *the*`k` *most frequent elements*. You may return the answer in **any order**.

## Examples

**Example 1:**

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]

```

**Constraints:**

- `1 <= nums.length <= 105`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

**Follow up:** Your algorithm's time complexity must be better than `O(n log n)`, where n is the array's size.

## 解析

題目給了一個整數陣列 nums 還有一個正整數 k

要求寫一個演算法來找出前 k 個最常出現的整數

直覺的作法是用一個 HashMap 來紀錄每個元素出現的次數

然後對這個 HashMap 的 key 做 sort 但這樣做會是 O(nlogn)

要比 O(nlogn) 還要優化

這邊可以透過 Bucket Sort 的方式來做處理

概念是把所有元素依照值做分類到不同的 bucket

因為這邊 每個元素最多出現 nums 陣列長度 n

所以剛好會是 bucket n 到 0

如下圖
![](https://i.imgur.com/pnDijwU.png)

因為最多 n 個

所以只到找 n 次 就可以 

時間複雜度 O(n)

## 程式碼
```go
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

```
## 困難點

1. 需要透過 Bucket Sort 來做優化

## Solve Point

- [x]  建立 HashMap 紀錄 freq
- [x]  建立 bucket 來做 bucketSort