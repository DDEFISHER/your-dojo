package main

import "fmt"
import "os"
import "bufio"

func AskAllQuestions() {

	mock_quiz1 := Question{"What color is the sky?", "blue"}
	mock_quiz2 := Question{"Who casts the one ring in the fire of mount doom", "Frodo"}
	mock_questions := Questions{mock_quiz1, mock_quiz2}

	for _, value := range mock_questions {
		fmt.Println(value.Question)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		printResult(scanner.Text(), value.Answer)
	}

}

// TODO: Check if terminal and if it supports color
func printResult(input, answer string) {
	if input == answer {
		fmt.Println("\033[1;32mCorrect\033[0m")
	} else {
		fmt.Println("\033[0;31mWrong\033[0m")
	}
}
