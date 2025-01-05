package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for i := 0; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	results := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if max <= candies[i] + extraCandies {
			results[i] = true
		}
	}

	return results
}