package main

import (
	"sort"
)

// https://programmers.co.kr/learn/courses/30/lessons/42748?language=go
func solution(array []int, commands [][]int) []int {
    var result []int

	for i := 0; i < len(commands); i++ {
        sortedArray := deepCopy(array[commands[i][0] - 1:commands[i][1]])
        sort.Ints(sortedArray)
        
        result = append(result, sortedArray[commands[i][2] - 1]) 
	}

    return result
}

func deepCopy(array []int) []int {
    copiedArray := make([]int, len(array))
    copy(copiedArray, array)
    return copiedArray
}