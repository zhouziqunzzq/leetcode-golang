package mycode

func canFormArray(arr []int, pieces [][]int) bool {
	headToId := make([]int, 110)
	for i := range headToId {
		headToId[i] = -1
	}
	for i, p := range pieces {
		headToId[p[0]] = i
	}

	// match pieces to arr naively
	arrP := 0
	for arrP < len(arr) {
		id := headToId[arr[arrP]]
		if id == -1 {
			return false
		}

		p := pieces[id]
		for i, v := range p {
			if v != arr[arrP] {
				return false
			}
			arrP++
			if arrP == len(arr) && i != len(p)-1 {
				return false
			}
		}
	}

	return true
}
