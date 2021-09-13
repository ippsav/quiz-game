package main

import (
	"flag"
	"fmt"
	"log"
	"main/csv"
	"main/game"
	"os"
)

func main() {
	fileName := flag.String("csv", "quiz.csv", "A CSV file containing the problems")
	duration := flag.Int("duration", 30, "Max duration for the quizz")
	flag.Parse()
	l := log.New(os.Stdout, "quizz-app: ", log.LstdFlags)
	fmt.Println("Quizz program started")
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Couldn t read the specified file %s \n", *fileName)
	}
	cFile := csv.NewCSVFile(file, l)
	lines := cFile.ParseCSVFile()
	// Init Game
	game := game.NewGame(l, &lines, duration)
	c, qts := game.StartGame()

	fmt.Printf("\nYou have %d correct asnwers out of %d\n", c, qts)
}
