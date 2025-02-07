package main

func main() {
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	last := len(nums1) - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[last] = nums1[i]
			i--
		} else {
			nums1[last] = nums2[j]
			j--
		}
		last--
	}

	if i <= 0 && j >= 0 {
		for x := 0; x <= j; x++ {
			if x < n {
				nums1[x] = nums2[x]
			}
		}
		// nums1[0 : j+1] = nums2[0 : j+1]
	}
}
