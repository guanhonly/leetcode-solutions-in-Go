package DFSAndBFS

func findSubsequences(nums []int) [][]int {
    var res [][]int
    var dfs func(startIndex int, tmp []int)
    dfs = func(startIndex int, tmp []int) {
        if len(tmp) > 1 {
            ttmp := make([]int, len(tmp))
            copy(ttmp, tmp)
            res = append(res, ttmp)
        }
        added := make(map[int]bool)
        for i:=startIndex; i<len(nums); i++ {
            if added[nums[i]] {
                continue
            }
            if len(tmp) == 0 || nums[i] >= tmp[len(tmp)-1] {
                added[nums[i]] = true
                //tmp = append(tmp, nums[i])
                dfs(i+1, append(tmp, nums[i]))
                //tmp = tmp[:len(tmp)-1]
            }
        }
    }
    var path []int
    dfs(0, path)
    return res
}

func Find(nums []int) [][]int {
    return findSubsequences(nums)
}
