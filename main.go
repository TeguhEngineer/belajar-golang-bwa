package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Halo, belajar perkalian")
	// result := calculation.Perkalian(2, 5)
	// fmt.Println(result)

	// Variable
	// var nama string = "teguh afriansyah umur"
	// umur := 20
	// tahun := "tahun"
	// fmt.Println(nama, umur, tahun)

	// IF 1
	// nilai := 80
	// if nilai < 85 {
	// 	fmt.Println("Tidak lulus")
	// } else {
	// 	fmt.Println("Lulus")
	// }

	// IF & ELSE IF
	// nilai := 90
	// var grade string
	// if nilai <= 60 {
	// 	grade = "E"
	// } else if nilai <= 70 {
	// 	grade = "D"
	// } else if nilai < 80 {
	// 	grade = "C"
	// } else if nilai < 90 {
	// 	grade = "B"
	// } else {
	// 	grade = "NULL"
	// }
	// fmt.Println(grade)

	// Switch
	// fmt.Println(nomorhari(1))

	// looping for
	// for i := 1; i <= 100; i++ {
	// 	println("Hallo, Belajar Golang", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	println("Hallo, Belajar Golang", i)
	// 	i++
	// }

	title := "Teguh Afriansyah"

	// for i, letter := range title {
	// 	if i%2 == 0 {
	// 		fmt.Println("letter :", string(letter))
	// 	}
	// }

	for _, letter := range title {
		switch string(letter) {
		case "a", "i", "u", "e", "o":
			fmt.Println("letter :", string(letter))
		}
	}
}

func huruf(abjad string) string {
	return abjad
}

func nomorhari(hari int) string {
	switch hari {
	case 1:
		return "Hari Senin"
	case 2:
		return "Hari Selasa"
	case 3:
		return "Hari Rabu"
	case 4:
		return "Hari Kamis"
	case 5:
		return "Hari Jumat"
	case 6:
		return "Hari Sabtu"
	case 7:
		return "Hari Minggu"
	default:
		return "Tidak ada hari"
	}
}
