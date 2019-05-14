// https://leetcode-cn.com/problems/word-ladder/
// 127. 单词接龙
package cpt6

type WordInfo struct {
	Val  string
	Step int
}

func diff(s1 string, s2 string) (count int) {
	count = len(s1)
	for idx, val := range s1 {
		if val == int32(s2[idx]) {
			count--
		}
	}
	return
}

// v1
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var (
		queue     []WordInfo      // wordList队列
		lsize     = len(wordList) // wordList大小
		isVisited = make([]bool, lsize)
	)
	if beginWord == endWord {
		return 2
	}

	// 先验证列表中是否有数据
	hasEndWord := false
	for i := 0; i < lsize; i++ {
		isVisited[i] = false
		if wordList[i] == endWord {
			hasEndWord = true
		}
	}
	if !hasEndWord {
		return 0
	}

	queue = append(queue, WordInfo{Val: beginWord, Step: 1})

	for {
		var (
			qsize = len(queue)
			currW WordInfo
		)
		if qsize == 0 {
			return 0
		}
		currW = queue[0]
		queue = queue[1:]
		for i := 0; i < lsize; i++ {
			if isVisited[i] {
				continue
			}
			if diff(currW.Val, wordList[i]) == 1 {
				if endWord == wordList[i] {
					return currW.Step + 1
				}
				queue = append(queue, WordInfo{Val: wordList[i], Step: currW.Step + 1})
				isVisited[i] = true
			}
		}
	}
}
