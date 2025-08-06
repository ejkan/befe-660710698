package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}

func (s *Student) IsHonour() bool{
	return s.GPA >= 3.50
}

func (s *Student) Validate() error{
	if s.Name == "" {
		return errors.New("name is required for this field")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 1-4")
	}
	return nil
}

func main(){
	// var st Student = Student{
		// ID: "660710698",
		// Name: "Jirawat",
		// Email: "nuanlaong_j@su.ac.th",
		// Year: 3,
		// GPA: 3.29,
	// }
	students := []Student{
		{
			ID: "1",
			Name: "Ji",
			Email: "nuanlaong_j@su.ac.th",
			Year: 3,
			GPA: 3.29,
		},
		{
			ID: "2",
			Name: "",
			Email: "ji_n@su.ac.th",
			Year: 2,
			GPA: 3.45,
		},
	}
	newStudent := Student{
		ID: "3",
		Name: "Alice",
		Email: "jane_a@su.ac.th",
		Year: 1,
		GPA: 4.00,
	}
	students = append(students, newStudent)
	for i, student:= range students{
		fmt.Printf("Student %d Honourable : %v\n", i+1, student.IsHonour())
		fmt.Printf("Student %d Fields validation : %v\n", i+1, student.Validate())
		fmt.Printf("\n")
	}
}