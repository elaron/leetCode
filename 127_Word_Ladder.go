package main

import (
	"fmt"
)

func mainBFS()  {
	wordlist := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	ret := ladderLength("hit", "cog", wordlist)
	//wordlist := []string{"cog"}
	//ret := ladderLength("hog", "cog", wordlist)
	fmt.Print(ret)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m_set, visit := make(map[string]bool, 0), make(map[string]bool, 0)
	for _, word := range wordList {
		m_set[word] = true
	}
	if !m_set[endWord] {
		return 0
	}
	beginQueue := []string{beginWord}
	endQueue := []string{endWord}
	round := 0
	for len(beginQueue) > 0 && len(endQueue) > 0 {
		round++
		if len(beginQueue) > len(endQueue) {
			beginQueue, endQueue = endQueue, beginQueue
		}
		nextQ := make([]string, 0)
		for _, word := range beginQueue {
			str := []byte(word)
			for i := 0; i < len(str); i++ {
				oriC :=  str[i]
				for c:='a'; c <= 'z'; c++ {
					str[i] = byte(c)
					target := string(str)
					if !m_set[target] {
						continue
					}
					if contains(endQueue, target) {
						return round+1
					}
					if m_set[target] && !visit[target] {
						visit[target] = true
						nextQ = append(nextQ, target)
					}
				}
				str[i] = oriC
			}
		}
		beginQueue = nextQ
	}
	return 0
}

func contains(list []string, target string) bool {
	for _, word := range list {
		if target == word {
			return true
		}
	}
	return false
}
