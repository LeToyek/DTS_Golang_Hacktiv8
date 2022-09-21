package main

import "fmt"

func main() {
	greet("Toyek", "Melbourne,Australia")
	result := greet2("Hello", []string{"Jona", "Joni", "Jupri"})
	fmt.Println(result)
	luas, keliling := circle(8)
	fmt.Println("Luas = ", luas)
	fmt.Println("Keliling = ", keliling)
	studentList := GetStudents("johan", "bambang", "supriyanto", "cipto")
	for i, v := range studentList {
		temp := fmt.Sprintf("Student %d", i+1)
		fmt.Println(v[temp])

	}
}
