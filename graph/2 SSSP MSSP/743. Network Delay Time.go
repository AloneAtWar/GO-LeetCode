package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 13:52

 * @Note:

 **/

const MAX = 100000

func networkDelayTime(times [][]int, n int, k int) int {
	distances := make([]int, n)
	for i := 0; i < n; i++ {
		distances[i] = MAX
	}
	distances[k-1] = 0
	for i := 0; i < n; i++ {
		for _, time := range times {
			from, to, cost := time[0]-1, time[1]-1, time[2]
			newCost := distances[from] + cost
			if distances[to] > newCost {
				distances[to] = newCost
			}
		}
	}
	max := 0
	for i, distance := range distances {
		if distance == MAX {
			return -1
		} else if distance > distances[max] {
			max = i
		}
	}
	return distances[max]
}
