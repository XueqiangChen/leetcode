package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. 构建邻接表(v是k的后续课程)，k：课程1，v：[课程0，课程2]
	afterMp := map[int][]int{}
	// 2. 每门课程的入度数组，index：某课程，v：该课程的入度
	inDegrees := make([]int, numCourses)
	for _, p := range prerequisites {
		after, first := p[0], p[1]                     // after后修课程，first先修课程
		afterMp[first] = append(afterMp[first], after) // 通过邻接表，构建课程关系依赖
		//因为学after之前必须先修first，所以after的入度+1
		inDegrees[after]++
	}
	// BFS 的queue
	var q []int
	// 初始化q，入度为0的课程进入队列
	for courseIdx, inDegree := range inDegrees {
		if inDegree == 0 {
			q = append(q, courseIdx)
		}
	}
	// 学过的课程数量
	learnedCourseCount := 0
	for len(q) > 0 {
		// 学过的课程数+1
		learnedCourseCount++
		// q内均是入度为0的课程，q.pop()则意味着学完了次learnedCourse课
		learnedCourse := q[0]
		// 同上，共同完成q.pop()动作
		q = q[1:]
		for _, after := range afterMp[learnedCourse] {
			// 因为已经学完after的先修课程，所以after的入度--
			inDegrees[after]--
			// 如果after已经无先修课程，入度为0，则进入队列
			if inDegrees[after] == 0 {
				q = append(q, after)
			}
		}
	}

	// 检查课程是否全部学完
	return learnedCourseCount == numCourses
}

func main() {
	prerequisites := [][]int{
		{
			1, 0,
		},
		{
			0, 1,
		},
	}
	res := canFinish(2, prerequisites)
	fmt.Println(res)
}
