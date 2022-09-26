package main

import "fmt"

func main() {
	choice := 0
	for 0 <= choice && choice < 3 {
		showMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			ShowAllStudents()
			break
		case 2:
			fmt.Println("Silahkan isi id murid yang ingin dicari")
			var search uint8
			fmt.Scanln(&search)
			FindStudent(search)
			break
		}
	}
}
func showMenu() {
	fmt.Println("=====================================")
	fmt.Println("Silahkan pilih menu")
	fmt.Println("1. Lihat seluruh data siswa")
	fmt.Println("2. Cari data siswa berdasarkan ID")
	fmt.Println("3. Keluar")
}
