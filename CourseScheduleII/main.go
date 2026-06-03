const (
	NOT_TOOK = 0
	PENDING  = 1
	TOOK     = 2
)

type course struct {
	num           int
	coursesToTake []*course
	state         int
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int]*course)
	for _, p := range prerequisites {
		if c, ok := graph[p[0]]; !ok {
			graph[p[0]] = &course{
				num:           p[0],
				coursesToTake: make([]*course, 1),
			}
			graph[p[0]].coursesToTake[0] = &course{
				num: p[1],
			}
		} else {
			c.coursesToTake = append(
				c.coursesToTake,
				&course{
					num: p[1],
				},
			)
		}
	}
	result := make([]int, numCourses)
	curIndex := 0
	for i := range numCourses {
		curIndex = takeCourse(i, graph, result, curIndex)
		if curIndex == -1 {
			return make([]int, 0)
		}
	}
	return result
}

func takeCourse(courseIndex int, graph map[int]*course, courseOrder []int, curOrderPosition int) int {
	if c, ok := graph[courseIndex]; !ok {
		graph[courseIndex] = &course{
			num:   courseIndex,
			state: TOOK,
		}
		courseOrder[curOrderPosition] = courseIndex
		return curOrderPosition + 1
	} else if c.state == TOOK {
		return curOrderPosition
	} else if c.state == PENDING {
		return -1
	} else {
		c.state = PENDING
		for i := range c.coursesToTake {
			curOrderPosition = takeCourse(
				c.coursesToTake[i].num,
				graph,
				courseOrder,
				curOrderPosition,
			)
			if curOrderPosition == -1 {
				return -1
			}
		}
		courseOrder[curOrderPosition] = courseIndex
		c.state = TOOK
		return curOrderPosition + 1
	}
}