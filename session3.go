package main

import (
	"fmt"
	"math"
	"strings"
)

// function
func greet(name string, address string) {
	fmt.Println("Hello my name is " + name)
	fmt.Println("I live in " + address)
}
func greet2(msg string, names []string) string {
	joinStr := strings.Join(names, " ")

	result := fmt.Sprintf("%s %s", msg, joinStr)
	return result
}
func circle(r float64) (float64, float64) {
	luas := math.Pow(r, 2) * math.Pi
	keliling := math.Pi * r * 2
	return luas, keliling
}
func GetStudents(name ...string) []map[string]string {
	var result []map[string]string
	for i, v := range name {
		key := fmt.Sprintf("Student %d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}
