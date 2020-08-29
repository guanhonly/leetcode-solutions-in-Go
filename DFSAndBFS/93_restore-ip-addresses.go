package DFSAndBFS

const SEG_COUNT = 4

var segments []int
var ans []string

func restoreIpAddresses(s string) []string {
	segments = make([]int, SEG_COUNT)
	ans = make([]string, 0)
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, startIndex, segID int) {
	if segID == SEG_COUNT {
		if startIndex == len(s) {
			var ip string
			for i := 0; i < SEG_COUNT; i++ {
				ip += strconv.Itoa(segments[i])
				if i < SEG_COUNT-1 {
					ip += "."
				}
			}
			ans = append(ans, ip)
		}
		return
	}
	if startIndex == len(s) {
		return
	}

	if s[startIndex] == '0' {
		segments[segID] = 0
		dfs(s, startIndex+1, segID+1)
	}

	addr := 0
	for endIndex := startIndex; endIndex < len(s); endIndex++ {
		addr = addr*10 + int(s[endIndex]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[segID] = addr
			dfs(s, endIndex+1, segID+1)
		} else {
			break
		}
	}
}
