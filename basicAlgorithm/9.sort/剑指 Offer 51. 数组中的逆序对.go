package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 18:40

 * @Note:	https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
			剑指 Offer 51. 数组中的逆序对
			剑指 Offer 51. 数组中的逆序对
 **/

// 归并排序
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) >> 1
	leftSize := mergeSort(nums, left, mid)
	rightSize := mergeSort(nums, mid+1, right)
	mergeSize := 0
	temp := make([]int, right-left+1)
	i := 0
	leftI, rightI := left, mid+1
	for leftI <= mid && rightI <= right {
		if nums[leftI] <= nums[rightI] {
			temp[i] = nums[leftI]
			i++
			leftI++
		} else {
			mergeSize += mid - leftI + 1
			temp[i] = nums[rightI]
			i++
			rightI++
		}
	}
	for leftI <= mid {
		temp[i] = nums[leftI]
		leftI++
		i++
	}
	for rightI <= right {
		temp[i] = nums[rightI]
		rightI++
		i++
	}
	for j := 0; j < right-left+1; j++ {
		nums[left+j] = temp[j]
	}
	return leftSize + rightSize + mergeSize
}
