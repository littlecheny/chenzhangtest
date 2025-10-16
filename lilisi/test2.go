package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @param k int整型
 * @return long长整型
 */
func ans(array []int, k int) int64 {
	// write code here
	ans := 0
	m := make(map[int]int, len(array))
	for i, v := range array {
		m[v] = i
	}
	for i, v := range array {
		if v >= k {
			continue
		} else {
			for n := range k - v {
				if j, ok := m[k-n-v]; ok && j > i {
					ans++
				}
			}
		}
	}
	return int64(ans)
}
