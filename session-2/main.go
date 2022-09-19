package main

import "fmt"

func main() {
	mySlice()
}
func mySlice() {
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
func arrays() {
	numbers := [4]int{1, 2, 3, 45}
	balances := [2][3]int{{1, 2, 3}, {2, 3, 4}}

	fmt.Print(numbers)
	fmt.Print(balances)
}
func loops() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
}
func conditions() {
	var currentYear = 2021

	if age := currentYear - 2010; age < 17 {
		fmt.Print("Kamu belum cukup umur yaitu (", age, " tahun)")
	} else {
		fmt.Print("Kamu sudah cukup umur")
	}
}
