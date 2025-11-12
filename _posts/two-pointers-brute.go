package main

func main(){
	result := twoPairSum([]int{-3, -1, 0, 2, 5}, 4)
	println(result[0])
	println(result[1])

	result = twoPairSumTwoPointers([]int{-3, -1, 0, 2, 5}, 4)
	println(result[0])
	println(result[1])
}
func twoPairSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++{
			if numbers[j] + numbers[i] == target{
				return []int{i, j}
			}
		}
	}
	return []int{};
}

func twoPairSumTwoPointers(numbers []int, target int) [] int {
	left := 0;
	right := len(numbers) - 1

	for left < right{
		sum := numbers[left] + numbers[right]

		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left, right}
		}
	}

	return []int{};
}