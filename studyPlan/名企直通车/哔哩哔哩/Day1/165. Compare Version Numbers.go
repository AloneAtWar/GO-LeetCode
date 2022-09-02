package main

import (
	"strconv"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 20:06

 * @Note:	https://leetcode.cn/problems/compare-version-numbers/?plan=zhitongche&plan_progress=zmq4dzv
			165. Compare Version Numbers
			165. 比较版本号
 **/

func compareVersion(version1 string, version2 string) int {
	split1 := strings.Split(version1, ".")
	split2 := strings.Split(version2, ".")
	n, m := len(split1), len(split2)
	i := 0
	for i < n && i < m {
		atoi1, _ := strconv.Atoi(split1[i])
		atoi2, _ := strconv.Atoi(split2[i])
		if atoi1 > atoi2 {
			return 1
		} else if atoi1 < atoi2 {
			return -1
		} else {
			i++
		}
	}
	if n == m {
		return 0
	}
	if i == n {
		for i < m {
			atoi, _ := strconv.Atoi(split2[i])
			if atoi != 0 {
				return -1
			}
			i++
		}
		return 0
	} else {
		for i < n {
			atoi, _ := strconv.Atoi(split1[i])
			if atoi != 0 {
				return 1
			}
			i++
		}
		return 0
	}
}
