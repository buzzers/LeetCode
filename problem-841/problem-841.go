package problem_841

func canVisitAllRooms(rooms [][]int) bool {
	// 深度遍历查看是否是连通图
	visited := make([]bool, len(rooms))
	return visit(rooms, 0, visited, 0) == len(rooms)
}

func visit(rooms [][]int, id int, visited []bool, n int) int {
	visited[id] = true
	for _, room := range rooms[id] {
		if !visited[room] {
			n = visit(rooms, room, visited, n)
		}
	}
	return n + 1
}
