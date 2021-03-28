package print

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Question struct {
	Country string
	Capital []string
	Comment string
}

type Result struct {
	Answer  string
	Country string
	Capital string
	OK      bool
}

type Classification struct {
	Result   []Result
	Position int
	Correct  int
}

var qtn Question
var rst Classification

// showing the question to the user
func showQuestion(q Question) {
	if q.Comment != "" {
		fmt.Println("You didn't finish the game ")
	}

	fmt.Printf("what is the Capital of %s?.\n", q.Country)
	for i, c := range q.Capital {
		fmt.Printf("%d) %s.\n", i+1, c)
	}
	fmt.Println("Select the number please ")

}

// showing the result of the Game
func ShowResult(r Classification) {
	fmt.Println("Result:")
	fmt.Printf("you have %d answer(s) correct .\n", r.Correct)
	for _, c := range r.Result {
		fmt.Printf(" The capital of %s is %s , your answer was %s, and this is %t.\n", c.Country, c.Capital, c.Answer, c.OK)
	}
	fmt.Printf("You were better than %d%% of all quizzers.\n", r.Position)
}

// handle the response from server side
func HandleResponse(h *http.Response, err error) {

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(h.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(responseData), &qtn)
	if len(qtn.Capital) > 1 {
		showQuestion(qtn)
	} else {
		json.Unmarshal([]byte(responseData), &rst)
		ShowResult(rst)
	}
}
