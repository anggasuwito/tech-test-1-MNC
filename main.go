package main

import "fmt"

func main() {
	var result string

	//result = NumberOne(4, []string{"abcd", "acbd", "aaab", "acbd"})
	//fmt.Println("1 : " + result)
	//result = NumberOne(5, []string{"pisang", "goreng", "enak", "sekali", "rasanya"})
	//fmt.Println("2 : " + result)
	//result = NumberOne(11, []string{"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"})
	//fmt.Println("3 : " + result)

	result = NumberTwo(700649, 800000)
	fmt.Println("1 : " + result)
	result = NumberTwo(657650, 600000)
	fmt.Println("2 : " + result)
	result = NumberTwo(575650, 580000)
	fmt.Println("3 : " + result)
}
