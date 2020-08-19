package mycode

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	for i := 1; candies > 0; candies -= i - 1 {
		if i <= candies {
			ans[(i-1)%num_people] += i
		} else {
			ans[(i-1)%num_people] += candies
		}
		i++
	}
	return ans
}
