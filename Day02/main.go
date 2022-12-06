package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Throw int

const (
	Rock Throw = iota + 1
	Paper
	Scissors
)

type Outcome int

const (
	Lose Outcome = iota * 3
	Draw
	Win
)

func DetermineOutcome(opponent Throw, mine Throw) Outcome {
	if mine == Rock && opponent == Paper {
		return Lose
	} else if mine == Scissors && opponent == Rock {
		return Lose
	} else if mine == Paper && opponent == Scissors {
		return Lose
	} else if mine == Paper && opponent == Rock {
		return Win
	} else if mine == Scissors && opponent == Paper {
		return Win
	} else if mine == Rock && opponent == Scissors {
		return Win
	} else {
		return Draw
	}
}

func DetermineThrow(opponent Throw, result Outcome) Throw {
	if opponent == Rock && result == Win {
		return Paper
	} else if opponent == Paper && result == Win {
		return Scissors
	} else if opponent == Scissors && result == Win {
		return Rock
	} else if opponent == Rock && result == Lose {
		return Scissors
	} else if opponent == Paper && result == Lose {
		return Rock
	} else if opponent == Scissors && result == Lose {
		return Paper
	} else {
		return opponent
	}
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	totalIfCypherOutcome, totalIfCypherThrow := 0, 0
	for scan.Scan() {
		text := scan.Text()

		opponentThrow := Throw(int(text[0] - 64))
		cypher := int((text[2]) - 87)

		outcomeIfCypherIsThrow := DetermineOutcome(opponentThrow, Throw(cypher))
		throwIfCypherIsOutcome := DetermineThrow(opponentThrow, Outcome(3*(cypher-1)))

		totalIfCypherOutcome += int(outcomeIfCypherIsThrow) + int(Throw(cypher))
		totalIfCypherThrow += int(Outcome(3*(cypher-1))) + int(throwIfCypherIsOutcome)
	}
	fmt.Println(fmt.Sprint(totalIfCypherOutcome))
	fmt.Println(fmt.Sprint(totalIfCypherThrow))
}
