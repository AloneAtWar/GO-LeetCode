package main

import "math"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 23:29

 * @Note:	https://leetcode.cn/problems/min-cost-to-connect-all-points/
			1584. Min Cost to Connect All Points
			1584. 连接所有点的最小费用
 **/

func minCostConnectPoints(points [][]int) (result int) {
	n := len(points)
	visited := make([]bool, n)
	distances := make([]int, n)
	visited[0] = true
	distances[0] = 0
	for i, point := range points[1:] {
		distances[i+1] = calculationDistance(points[0], point)
	}
	i := 1
	for i < n {
		minDistance, index := math.MaxInt64, 0
		for j, distance := range distances {
			if !visited[j] && distance < minDistance {
				minDistance = distance
				index = j
			}
		}
		visited[index] = true
		result += minDistance
		for j, distance := range distances {
			if !visited[j] {
				// 上面的写法将会变成 最短路径
				//newDistance:=calculationDistance(points[index],points[j])+minDistance
				newDistance := calculationDistance(points[index], points[j])
				if newDistance < distance {
					distances[j] = newDistance
				}
			}
		}
	}
	return result
}

func calculationDistance(point1 []int, point2 []int) int {
	return abs(point1[0]-point2[0]) + abs(point1[1]-point2[1])
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
