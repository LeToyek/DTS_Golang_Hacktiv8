package main

import (
	"fmt"

	"github.com/LeToyek/DTS_Golang_Hacktiv8/model"
)

var students []model.User

func init() {
	students = []model.User{
		{
			Id:        1,
			Nama:      "Maulana",
			Alamat:    "Jalan Gayam, Malang",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin menjadi backend dev",
		},
		{
			Id:        2,
			Nama:      "Arif",
			Alamat:    "Jalan Ijen, Malang",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin bisa membuat REST API",
		},
		{
			Id:        3,
			Nama:      "Wijaya",
			Alamat:    "Jalan Mahoni, Gentan, Sukoharjo",
			Pekerjaan: "junior backend dev",
			Alasan:    "Ingin menambah ilmu baru",
		},
	}
}
func ShowAllStudents() {
	for _, s := range students {
		fmt.Println("-----------------------")
		fmt.Println(s.Id)
		fmt.Println(s.Nama)
		fmt.Println(s.Alamat)
		fmt.Println(s.Pekerjaan)
		fmt.Println(s.Alasan)
	}
}
func FindStudent(studentID uint8) {
	for _, s := range students {
		if s.Id == studentID {
			fmt.Println(s.Id)
			fmt.Println(s.Nama)
			fmt.Println(s.Alamat)
			fmt.Println(s.Pekerjaan)
			fmt.Println(s.Alasan)
			return
		}
	}
	fmt.Printf("Student dengan id %s tidak ada pada database\n", studentID)
}
