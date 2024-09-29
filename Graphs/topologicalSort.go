package graphs

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.
//
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
//
//
//
// Example 1:
//
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0. So it is possible.

func canFinishParticularCourse(
	courseNum int,
	visiting *map[int]bool,
	visited *map[int]bool,
	prerequisites [][]int,
) bool {
	if _, ok := (*visiting)[courseNum]; ok {
		// we found a loop
		return false
	}

	if _, ok := (*visited)[courseNum]; ok {
		return true
	}

	(*visiting)[courseNum] = true

	for _, prereq := range prerequisites {
		// prereq[0] is the prerequisite for courseNum
		if prereq[1] == courseNum {
			// find out if we can actually complete prereq[0]
			if !canFinishParticularCourse(prereq[0], visiting, visited, prerequisites) {
				return false
			}

			(*visited)[courseNum] = true
		}
	}

	delete(*visiting, courseNum)

	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	visiting := map[int]bool{}

	// all the courses that we can complete
	visited := map[int]bool{}

	for course := range numCourses {
		// find if the course has some prerequisite
		if !canFinishParticularCourse(course, &visiting, &visited, prerequisites) {
            return false
        }

		visited[course] = true
	}

	return len(visited) == numCourses
}
