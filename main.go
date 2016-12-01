package main

import "fmt"

func main() {

	mock_quiz1 := Question{"What color is the sky?", "blue"}
	mock_quiz2 := Question{"Who casts the one ring in the fire of mount doom", "Frodo"}
	mock_questions := Questions{mock_quiz1, mock_quiz2}
	fmt.Println(mock_questions)
}
