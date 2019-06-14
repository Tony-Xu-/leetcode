func merge(nums1 []int, m int, nums2 []int, n int)  {
    idx1 := m - 1
    idx2 := n - 1
    idxTarget := m + n - 1
    for {
        if (idx1 < 0 && idx2 < 0) || idxTarget < 0 {
            return
        }
        currNum := 0
        if idx1 < 0 {
            //已经遍历完成原始nums1
            currNum = nums2[idx2]
            nums1[idxTarget] = currNum
            idxTarget--
            idx2--
            continue
        }
        if idx2 < 0 {
            //已经遍历完成原始nums2
            currNum = nums1[idx1]
            nums1[idxTarget] = currNum
            idxTarget--
            idx1--
            continue
        }
        if nums1[idx1] >= nums2[idx2] {
            currNum = nums1[idx1]
            idx1--
        } else {
            currNum = nums2[idx2]
            idx2--
        }
        nums1[idxTarget] = currNum
        idxTarget--
    }
}
