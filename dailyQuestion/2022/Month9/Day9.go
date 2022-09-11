package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/11 12:13

 * @Note:	https://leetcode.cn/problems/crawler-log-folder/
			1598. Crawler Log Folder
			1598. 文件夹操作日志搜集器
 **/

func minOperations(logs []string) (result int) {
	for _, log := range logs {
		if log == "./" {
			continue
		} else if log == "../" {
			if result != 0 {
				result--
			}
		} else {
			result++
		}
	}
	return result
}
