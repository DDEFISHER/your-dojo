package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math/rand"
import "time"

func AskAllQuestions() {

	mock_quiz1 := Question{"What color is the sky?", "blue"}
	mock_quiz2 := Question{"Who casts the one ring in the fire of mount doom", "Frod o"}
	mock_quiz3 := Question{"Who is Mr. Robot", "Elliot"}
	mock_quiz4 := Question{"How would you declare 'x' a variable of type int with the value 3", "x := 3"}
	mock_questions := Questions{mock_quiz1, mock_quiz2, mock_quiz3, mock_quiz4}
	amount := len(mock_questions)

	placed_map := make(map[int]int)

	rand.Seed(time.Now().Unix())
	random_questions := make([]Question, amount)

	for _, value := range mock_questions {
		r := rand.Intn(amount)
		for placed_map[r] == 1 {
			r = rand.Intn(amount)
		}
		placed_map[r] = 1
		random_questions[r] = value
	}

	for _, value := range random_questions {
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
