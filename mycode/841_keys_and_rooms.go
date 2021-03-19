package mycode

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)

	canVisitAllRoomsDFS(rooms, 0, visited)

	for _, f := range visited {
		if !f {
			return false
		}
	}
	return true
}

func canVisitAllRoomsDFS(rooms [][]int, cur int, visited []bool) {
	if cur < 0 || cur >= len(rooms) || visited[cur] {
		return
	}

	visited[cur] = true
	for _, r := range rooms[cur] {
		canVisitAllRoomsDFS(rooms, r, visited)
	}
}
