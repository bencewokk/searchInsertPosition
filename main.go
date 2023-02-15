func searchInsert(nums []int, target int) int {

	var start, middle, end int
	start = 0
	end = len(nums)
	middle = start + end/2

	var cache, cache2 int
	for {
		fmt.Println(start, middle, end, cache)
		middle = (start + end) / 2

		cache = nums[middle]
		if cache == cache2 {
			if target < cache {
				return middle
			} else {
				return middle+1
			}
		}
		cache2 = cache
        
		if cache == target {
			return middle
		} else if cache > target {
			end = middle
		} else {
			start = middle
		}
	}
}
