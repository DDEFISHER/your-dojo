package main

import "fmt"
import "os"
import "bufio"
import "strings"

func AskAllQuestions() {

	mock_quiz1 := Question{"What color is the sky?", "blue"}
	mock_quiz2 := Question{"Who casts the one ring in the fire of mount doom", "Frod o"}
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

	r := strings.NewReplacer(" ", "")

	if strings.ToLower(r.Replace(input)) == strings.ToLower(r.Replace(answer)) {
		fmt.Println("\033[1;32mCorrect\033[0m")
	} else {
		fmt.Println("\033[0;31mWrong\033[0m")
	}
}
