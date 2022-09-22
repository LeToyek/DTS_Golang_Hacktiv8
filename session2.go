package main

import "fmt"

func MySlice() {
	fruits := []string{"Apple", "Banana", "Mango"}

	secFruits := make([]string, 3)
	secFruits[0] = "Durian"
	secFruits[2] = "Nanas"
	secFruits = append(secFruits, "Apple")
	secFruits = append(secFruits, "Mango")
	secFruits = append(secFruits, fruits...)

	fmt.Println(fruits)
	fmt.Println(secFruits)

	no := copy(fruits, secFruits)
	fmt.Println(no)
	fmt.Println(secFruits[3:5])

}
func Arrays() {
	numbers := [4]int{1, 2, 3, 45}
	balances := [2][3]int{{1, 2, 3}, {2, 3, 4}}

	fmt.Print(numbers)
	fmt.Print(balances)
}
func Loops() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
}
func Conditions() {
	var currentYear = 2021

	if age := currentYear - 2010; age < 17 {
		fmt.Print("Kamu belum cukup umur yaitu (", age, " tahun)")
	} else {
		fmt.Print("Kamu sudah cukup umur")
	}
}

type Int int

func miniQuiz() {
	var result = map[Int]int{}
	a := []Int{1, 2, 2, 3, 4, 5, 5, 6}

	for _, n := range a {
		result[n]++
	}
	for key, v := range result {
		fmt.Printf("Angka %d : duplikasi %d\n", key, v)
	}
}
func miniQuiz2() {
	str := []string{"cal", "cal", "cal", "man", "man", "ta", "ta", "ra", "ra", "ra"}
	temp := map[string]int{}
	var tampung []string
	for _, v := range str {
		temp[v]++
	}
	for key, value := range temp {
		if value >= 3 {
			tampung = append(tampung, key)
		}
	}
	str2 := str
	str2[0] = "toyek"
	fmt.Print(str[0])
	fmt.Println(tampung)
}
