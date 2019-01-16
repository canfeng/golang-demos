package basic_sort

import "sort"

//冒泡排序

//冒泡排序算法的运作如下：
//
//1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
//2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
//3. 针对所有的元素重复以上的步骤，除了最后一个。
//4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

//时间复杂度：O(n^2)
func BubbleSort(data sort.IntSlice) {
	n := data.Len()
	for i := n; i > 0; i-- {
		isChanged := false
		for j := 0; j < i-1; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}
