package main

//slice --> dynamic
//most used construct in go
// + useful methods

func main() {
	//uninitialized slice is nill
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	//
	// // var nums = make([]int, 2, 5)
	// var nums = make([]int, 0, 5)
	// //capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)
	// nums = append(nums, 10)
	// nums = append(nums, 11)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//Copy funtion
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//
	//Slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[1])

	//
	//slice
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}
	// fmt.Println(slices.Equal(nums1, nums2))

	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)
}
