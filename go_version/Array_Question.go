package main

import (
	"fmt"
)

//Return Sum of elements in a array
func SumArray(array []int) int {
	var s int
	for _, item := range array {
		s = s + item
	}
	return s
}

//Search a list for a given value
func SearchValue(data []int, value int) bool {
	for _, val := range data {
		if val == value {
			return true
		}
	}
	return false
}

//Binary Search only for sorted list
//1. 计算中间位置索引将索引处的值与待查找值大小做比较
//2. 改变范围
func BinarySearch(data []int, value int) bool {
	var mid int
	low := 0
	high := len(data)
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else {
			if data[mid] > value {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}

//Rotating a list by K positions
func RotateLeftArray(data []int, k int) {
	n := len(data)
	reverseArray(data, 0, k-1)
	reverseArray(data, k, n-1)
	reverseArray(data, 0, n-1)
}

func RotateRightArray(data []int, k int) {
	n := len(data)
	reverseArray(data, 0, n-1-k)
	reverseArray(data, n-k, n-1)
	reverseArray(data, 0, n-1)
}

func reverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

//Find the largest sum contiguous subarray
func MaxSubArraySum(data []int) int {
	var maxSoFar = 0
	var maxEndingHere = 0
	for i := 0; i < len(data); i++ {
		maxEndingHere = data[i] + maxEndingHere
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	arr := []int{-1, -3, -2, -5, -7}
	s := SumArray(a)
	b := SearchValue(a, 4)
	c := BinarySearch(a, 3)
	RotateRightArray(a, 4)
	fmt.Println(s, b, c)
	fmt.Println(a)
	fmt.Println("最大值为:", MaxSubArraySum(arr))
}
