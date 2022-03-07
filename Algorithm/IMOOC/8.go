package imooc

import "strconv"

//递归与回溯问题

//93
func restoreIpAddresses(s string) []string {
	res := []string{}
	ips := []string{}
	if len(s) < 4 {
		return res
	}
	restoreIpAddresses_back(s, 0, ips, &res)
	return res
}

func restoreIpAddresses_back(s string, index int, ips []string, res *[]string) {
	if len(s) == index && len(ips) == 4 {
		newRes := ips[0] + "." + ips[1] + "." + ips[2] + "." + ips[3]
		*res = append(*res, newRes)
	}

	for i := 1; i <= 3; i++ {
		if i+index > len(s) {
			return
		}
		ip := s[index : index+i]
		if !isIp(ip) {
			continue
		}
		ips = append(ips, ip)
		restoreIpAddresses_back(s, index+i, ips, res)
		ips = ips[:len(ips)-1]
	}
}

func isIp(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return i >= 0 && i <= 255
}

//131
// func partition(s string) [][]string {
// 	res := [][]string{}
// 	path := []string{}
// 	partition_back(s, 0, path, &res)
// 	return res
// }

// func partition_back(s string, index int, path []string, res *[][]string) {
// 	if index == len(s) {
// 		newP := make([]string, index)
// 		copy(newP, path)
// 		*res = append(*res, newP)
// 	}

// 	for i := 0; i < len(s); i++ {
// 		if i+index > len(s)-1 {
// 			return
// 		}
// 		p := s[index : index+i]
// 		if !isBack(p) {
// 			continue
// 		}
// 		path = append(path, p)
// 		partition_back(s, index+i, path, res)
// 		path = path[:len(path)-1]
// 	}
// }

// func isBack(s string) bool {
// 	start, end := 0, len(s)-1
// 	for start < end {
// 		if s[start] != s[end] {
// 			return false
// 		}
// 		start++
// 		end--
// 	}
// 	return true
// }

//46
//47
