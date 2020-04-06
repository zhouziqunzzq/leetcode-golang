package mycode

func romanToInt(s string) int {
	i := 0
	ans := 0
	ss := []rune(s)
	for i = 0; i < len(ss); i++ {
		switch ss[i] {
		case 'M':
			ans += 1000
		case 'D':
			ans += 500
		case 'C':
			if i+1 < len(ss) {
				switch ss[i+1] {
				case 'M':
					ans += 900
					i++
				case 'D':
					ans += 400
					i++
				default:
					ans += 100
				}
			} else {
				ans += 100
			}
		case 'L':
			ans += 50
		case 'X':
			if i+1 < len(ss) {
				switch ss[i+1] {
				case 'C':
					ans += 90
					i++
				case 'L':
					ans += 40
					i++
				default:
					ans += 10
				}
			} else {
				ans += 10
			}
		case 'V':
			ans += 5
		case 'I':
			if i+1 < len(ss) {
				switch ss[i+1] {
				case 'X':
					ans += 9
					i++
				case 'V':
					ans += 4
					i++
				default:
					ans += 1
				}
			} else {
				ans += 1
			}
		}
	}
	return ans
}
