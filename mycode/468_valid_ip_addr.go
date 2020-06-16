package mycode

import "strconv"

func validIPAddress(IP string) string {
	// buf state
	buf := make([]rune, 0, 4)
	chInBuf := false

	// global state
	groupCnt := 0
	dotFlag := false
	colonFlag := false

	for _, c := range []rune(IP) {
		if c >= '0' && c <= '9' {
			buf = append(buf, c)
		} else if (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
			buf = append(buf, c)
			chInBuf = true
		} else if c == '.' {
			// assuming IPv4
			// check separator consistency
			dotFlag = true
			if colonFlag {
				return "Neither"
			}
			// check buf length
			if len(buf) == 0 || len(buf) > 3 {
				return "Neither"
			}
			// check leading zeros
			if buf[0] == '0' && len(buf) > 1 {
				return "Neither"
			}
			// check buf content
			if chInBuf {
				return "Neither"
			}
			if i, _ := strconv.Atoi(string(buf)); i < 0 || i > 255 {
				return "Neither"
			}

			// group check passed
			groupCnt++
			// resetting buf and state
			buf = make([]rune, 0, 4)
			chInBuf = false
		} else if c == ':' {
			// assuming IPv6
			// check separator consistency
			colonFlag = true
			if dotFlag {
				return "Neither"
			}
			// check buf length (including check leading zeros)
			if len(buf) == 0 || len(buf) > 4 {
				return "Neither"
			}

			// group check passed
			groupCnt++
			// resetting buf and state
			buf = make([]rune, 0, 4)
			chInBuf = false
		} else {
			return "Neither"
		}
	}

	// post-process last group
	if dotFlag && colonFlag {
		return "Neither"
	} else if dotFlag {
		// assuming IPv4
		// check buf length
		if len(buf) == 0 || len(buf) > 3 {
			return "Neither"
		}
		// check leading zeros
		if buf[0] == '0' && len(buf) > 1 {
			return "Neither"
		}
		// check buf content
		if chInBuf {
			return "Neither"
		}
		if i, _ := strconv.Atoi(string(buf)); i < 0 || i > 255 {
			return "Neither"
		}

		// group check passed
		groupCnt++
		// check group count
		if groupCnt == 4 {
			return "IPv4"
		} else {
			return "Neither"
		}
	} else {
		// assuming IPv6
		// check buf length (including check leading zeros)
		if len(buf) == 0 || len(buf) > 4 {
			return "Neither"
		}

		// group check passed
		groupCnt++
		// check group count
		if groupCnt == 8 {
			return "IPv6"
		} else {
			return "Neither"
		}
	}
}
