package leetcode

func extra_mn_merge(nums1 []int, m int, nums2 []int, n int) {
	ans := make([]int, m+n)
	i := 0
	j := 0
	p := 0
	for i < m || j < n {
		if i >= m {
			ans[p] = nums2[j]
			p++
			j++
			continue
		}
		if j >= n {
			ans[p] = nums1[i]
			p++
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			ans[p] = nums1[i]
			i++
			p++
		} else {
			ans[p] = nums2[j]
			j++
			p++
		}
	}
	copy(nums1, ans)
}

func lt88_merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	p := m + n - 1
	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[p] = nums2[j]
			p--
			j--
			continue
		}
		if j < 0 {
			nums1[p] = nums1[i]
			i--
			p--
			continue
		}
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			p--
			i--
		} else {
			nums1[p] = nums2[j]
			p--
			j--
		}
	}
}
