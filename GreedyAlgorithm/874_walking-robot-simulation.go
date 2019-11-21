/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/walking-robot-simulation/
 */
package GreedyAlgorithm

import "strconv"

/**
 * Time complexity: O(n + k), n is the length of commands, k is the length of obstacles
 * Space complexity: O(k)
 */
func robotSim(commands []int, obstacles [][]int) int {
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	currX, currY := 0, 0
	index := 0
	maxDistance := 0
	obstaclesSet := make(map[string]int)
	for _, obstacle := range obstacles {
		key := strconv.Itoa(obstacle[0]) + "-" + strconv.Itoa(obstacle[1])
		obstaclesSet[key] = 1
	}
	for i := 0; i < len(commands); i++ {
		currCommand := commands[i]
		if currCommand == -1 {
			index = (index + 1) % 4
		} else if currCommand == -2 {
			index = (index + 3) % 4
		} else {
			for j := 0; j < currCommand; j++ {
				nextX := currX + dx[index]
				nextY := currY + dy[index]
				nextKey := strconv.Itoa(nextX) + "-" + strconv.Itoa(nextY)
				if _, exist := obstaclesSet[nextKey]; !exist {
					currX = nextX
					currY = nextY
					currDistance := currX*currX + currY*currY
					if currDistance > maxDistance {
						maxDistance = currDistance
					}
				}
			}
		}
	}
	return maxDistance
}
