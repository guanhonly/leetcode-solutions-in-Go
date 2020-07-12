package Trie
/**
 哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。
 像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。
 在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。
 假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。

示例：

输入：
dictionary = ["looked","just","like","her","brother"]
sentence = "jesslookedjustliketimherbrother"
输出： 7
解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。

链接：https://leetcode-cn.com/problems/re-space-lcci
**/

// The main problem is how to find the word from the dictionary efficiently,
// Trie is a good approach, Rabin Karp algorithm is also good.
type Trie2 struct {
    next [26]*Trie2
    isEnd bool
}

func (this *Trie2) insert(s string) {
    currPos := this
    for i:=len(s)-1; i>=0; i-- {
        t := s[i] - 'a'
        if currPos.next[t] == nil {
            currPos.next[t] = &Trie2{
                next: [26]*Trie2{},
            }
        }
        currPos = currPos.next[t]
    }
    currPos.isEnd = true
}

func respace(dictionary []string, sentence string) int {
    n, inf := len(sentence), 0x3f3f3f3f
    root := &Trie2{
        next: [26]*Trie2{},
    }
    for _, word := range dictionary {
        root.insert(word)
    }
    dp := make([]int, n+1)
    for i:=1; i<=n; i++ {
        dp[i] = inf
    }
    for i:=1; i<=n; i++ {
        dp[i] = dp[i-1] + 1
        currPos := root
        for j:=i; j>=1; j-- {
            t := sentence[j-1] - 'a'
            if currPos.next[t] == nil {
                break
            } else if currPos.next[t].isEnd {
                dp[i] = min(dp[i], dp[j-1])
            }
            if dp[i] == 0 {
                break
            }
            currPos = currPos.next[t]
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}