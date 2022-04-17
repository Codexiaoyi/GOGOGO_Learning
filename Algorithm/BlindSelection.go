package main

//39
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		if candidates[i] == target {
			res = append(res, []int{target})
			continue
		}
		cs := combinationSum(candidates, target-candidates[i])
		for _, c := range cs {
			temp := make([]int, 0, len(c)+1)
			temp = append(temp, candidates[i])
			temp = append(temp, c...)
			res = append(res, temp)
		}
	}

	return res
}
