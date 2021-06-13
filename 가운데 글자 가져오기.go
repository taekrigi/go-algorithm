package main

// https://programmers.co.kr/learn/courses/30/lessons/12903
func solution(s string) string {
    isEven := isEven(len(s))
    index := len(s) / 2
    
    if isEven {
        return s[index - 1:index + 1]
    }
    
    return s[index:index + 1]
}

func isEven(len int) bool {
    return len % 2 == 0
}