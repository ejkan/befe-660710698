package main

import ("fmt")
// var email string = "nuanlaong_j@su.ac.th"

func main(){
	// var name string = "jirawat"
	var age int = 24
	email := "nuanlaong_j@su.ac.th"
	gpa := 2.99
	firstname, surname := "Jirawat", "Nuanlaong"
	fmt.Printf("Name : %s %s, Age : %d, Email : %s, GPA : %.2f\n", firstname, surname, age, email, gpa)
}