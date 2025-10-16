package main

import . "nc_tools"

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @return ListNode类
 */
func formatList(head *ListNode) *ListNode {
	// write code here
	nums := make([]int, 0)
	ans := head
	for ans != nil {
		nums = append(nums, ans.Val)
		ans = ans.Next
	}
	nums1 := make([]int, len(nums))
	mid := len(nums) / 2
	for i := range nums {
		if i%2 == 0 {
			nums1[mid-i/2] = nums[i]
		} else {
			nums1[mid+i/2] = nums[i]
		}
	}
	ans = head
	for i := range nums1 {
		ans.Val = nums1[i]
		ans = ans.Next
	}
	return head

}
