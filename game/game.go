package game

import (
	"fmt"
	"log"
	"main/problem"
	"time"
)

type Game struct {
	problems *problem.Problems
	duration *int
	l        *log.Logger
}

func NewGame(l *log.Logger, lines *[][]string, duration *int) *Game {
	pbs := problem.Problems{}
	pbs.ParseLines(*lines)
	return &Game{&pbs, duration, l}
}

func (g *Game) StartGame() (c, qts int) {
	timer := time.NewTimer(time.Duration(*g.duration) * time.Second)
	correctAnswers := 0
	answerChannel := make(chan string)
loop:
	for i, problem := range *g.problems {
		fmt.Printf("Question number %d: %s = ", i, problem.Question)
		go func() {
			var answer string
			_, err := fmt.Scanf("%s", &answer)
			if err != nil {
				g.l.Fatal(err)
			}
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			break loop
		case answer := <-answerChannel:
			if answer == problem.Answer {
				correctAnswers++
			}
		}
	}
	return correctAnswers, len(*g.problems)
}
