package problem0001

func twoSum(nums []int, target int) []int {

	// create an map to store item and index
	index := make(map[int]int, len(nums))
	// loop all the items in there and store the item as key, index as value into map. caculate the result with target
	for i, b := range nums {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}
		index[b] = i
	}
	return nil
}
