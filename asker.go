package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math/rand"
import "time"

func AskAllQuestions() {

	mock_questions := LoadYaml()
	amount := len(mock_questions.List)

	placed_map := make(map[int]int)

	rand.Seed(time.Now().Unix())
	random_questions := make([]Question, amount)

	for _, value := range mock_questions.List {
		r := rand.Intn(amount)
		for placed_map[r] == 1 {
			r = rand.Intn(amount)
		}
		placed_map[r] = 1
		random_questions[r] = value
	}

	for _, value := range random_questions {
		fmt.Println(value["question"])
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		printResult(scanner.Text(), value["answer"])
	}

}

// TODO: Check if terminal and if it supports color
func printResult(input, answer string) {

	r := strings.NewReplacer(" ", "")

	if strings.ToLower(r.Replace(input)) == strings.ToLower(r.Replace(answer)) {
		fmt.Println("\033[1;32mCorrect\033[0m\n")
	} else {
		fmt.Println("\033[0;31mWrong\033[0m\n")
	}
}
