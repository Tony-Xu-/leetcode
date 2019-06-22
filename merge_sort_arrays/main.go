package main

func main() {
	nums1 := []int{2,0}
	nums2 := []int{1}
	m,n:=1,1

	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if len(nums1) == 0 || len(nums2) == 0 || n==0 || len(nums1) < m+n {
		return
	}
	// 从nums1[m+n]开始倒序赋值
	idx1, idx2 := m-1, n-1
	if m < 0 {
		idx1 = 0
	}
	if n < 0 {
		idx2 = 0
	}
	if m == 0 {
		nums1 = nums2[:n]
		return
	}
	if  n == 0 {
		return
	}
	for targetIndex:=m+n-1; targetIndex>=0; targetIndex-- {
		if idx1<0 && idx2<0 {
			return
		}
		if (idx2 <0 && idx1 >= 0) {
			nums1[targetIndex] = nums1[idx1]
			idx1--
			continue
		}
		if (idx1 < 0 && idx2 >= 0) {
			nums1[targetIndex] = nums2[idx2]
			idx2--
			continue
		}
		if nums1[idx1] > nums2[idx2] {
			nums1[targetIndex] = nums1[idx1]
			idx1--
			continue
		}
		if nums1[idx1] <= nums2[idx2] {
			nums1[targetIndex] = nums2[idx2]
			idx2--
			continue
		}
	}
}