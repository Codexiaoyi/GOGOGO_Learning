package daily

import "sort"

//*******************************2021/7/30 171*******************
func titleToNumber(columnTitle string) int {
	res := 0
	multiple := 1
	for i := len(columnTitle); i >= 1; i-- {
		word := columnTitle[i-1] - 'A' + 1
		wordInt := int(word)
		res = res + wordInt*multiple
		multiple *= 26
	}
	return res
}

//*******************************69. x 的平方根*******************
func mySqrt(x int) int {
	count := 1
	for {
		sqrt := count * count
		if sqrt > x {
			return count - 1
		}
		count++
	}
}

//*******************************1337. 矩阵中战斗力最弱的 K 行*******************
func kWeakestRows(mat [][]int, k int) []int {
	type army struct {
		index int
		power int
	}
	armyPower := make([]*army, 0)
	for i := 0; i < len(mat); i++ {
		p := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				break
			}
			p++
		}
		armyPower = append(armyPower, &army{index: i, power: p})
	}
	res := make([]int, k)

	return res
}

//*******************************1337. 矩阵中战斗力最弱的 K 行 2021/8/3*******************
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	sortNums := make([]int, len(nums))
	copy(sortNums, nums)
	//先快排
	quickSort(sortNums, 0, len(sortNums)-1)
	//sort.Ints(sortNums)
	start, end := 0, len(nums)-1
	for {
		//找到最后都没有不同，那就说明原来就是排好序的
		if start == end {
			return 0
		}
		if sortNums[start] != nums[start] {
			break
		}
		start++
	}
	for {
		if start == end {
			return 0
		}
		if sortNums[end] != nums[end] {
			break
		}
		end--
	}
	return end - start + 1
}

func quickSort(arr []int, index_start, index_end int) {
	//定义递归出口
	if index_start >= index_end {
		return
	}
	//快排基准，第一个元素
	num := arr[index_start]
	i, j := index_start, index_end
	for {
		//如果到i,j相等还没找到就算了
		if i >= j {
			break
		}
		//将比基准小的放到基准的左边
		for {
			//i就是需要被换到的位置，j就是当前检测的
			if i >= j {
				break
			}
			if arr[j] < num {
				//当前检测的比基准小就放到i的位置，并把i往前移动，然后去下面的for从前往后
				arr[i] = arr[j]
				i++
				break
			}
			//如果没有找到比基准小的就继续往前找
			j--
		}
		//将比基准大的放到右边
		for {
			if i >= j {
				break
			}
			if arr[i] >= num {
				arr[j] = arr[i]
				j--
				break
			}
			i++
		}
	}
	//全部找完之后，把基准填到坑里
	arr[i] = num
	//左边继续排序
	quickSort(arr, index_start, i)
	//右边继续排序
	quickSort(arr, i, index_end)
}

//*******************************611. 有效三角形的个数 2021/8/4*******************
func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	result := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			result += sort.SearchInts(nums[j+1:], nums[i]+nums[j])
		}
	}
	return result
}
