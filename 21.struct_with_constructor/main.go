package main

import "fmt"

type student struct {
	rollNo  int
	fname   string
	lname   string
	grade   float32
	stander int
}

func newStudent(rollNo, stander int, fname, lname string, grade float32) *student {
	stu := student{
		rollNo:  rollNo,
		fname:   fname,
		lname:   lname,
		grade:   grade,
		stander: stander,
	}
	return &stu
}

func main() {
	// 🔥 - struct
	ankit := newStudent(1, 10, "Ankit", "Jodhani", 5.5)
	fmt.Println(*ankit)
	// fmt.Println(*&ankit.fname)
}
