package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

var bufferedReader = bufio.NewReader(os.Stdin)
var csvPtr = flag.String("csvFile", "problems.csv", "csv file in format: 'question,answer'")
var timePtr = flag.Int("timelimit", 25, "Time limit for quiz(in seconds)")
var shuffle = flag.String("shuffle", "No", "Ask if you want the quiz to be shuffled: 'Yes or No'")
var rowData [][]string

func main() {
	flag.Parse()
	file, err := os.Open(*csvPtr)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rowData, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problem := parseData(rowData)

	timer := time.NewTimer(time.Duration(*timePtr) * time.Second)
	fmt.Printf("You have %d seconds to answer this quiz\n", *timePtr)

	var correct int
	for index, prob := range problem {
		fmt.Printf("Problem #%d: %s= \n", index+1, prob.question)
		answerChannel := make(chan string)
		go func() {
			var answerEntered string
			answerEntered, err = bufferedReader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			answerEntered = strings.Replace(answerEntered, "\n\r", "", -1)
			answerEntered = strings.TrimSpace(answerEntered)
			//////OR///////
			//	var answerEntered string
			//	fmt.Scanln(&answerEntered)
			answerChannel <- answerEntered
		}()

		select {
		case <-timer.C:
			fmt.Printf("You got %d correct out of %d", correct, len(problem))
			return
		case answer := <-answerChannel:
			if answer == prob.answer {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Wrong!")
			}
		}
	}
	fmt.Printf("You got %d correct out of %d", correct, len(problem))
}

func parseData(rowData [][]string) []problem {
	filledData := make([]problem, len(rowData))
	for index, data := range rowData {
		if strings.EqualFold(*shuffle, "Yes") {

			i := genRand(index)
			filledData[i] = problem{
				question: data[0],
				answer:   data[1],
			}
		} else {
			filledData[index] = problem{
				question: data[0],
				answer:   data[1],
			}
		}
	}
	return filledData
}
func genRand(a int) (c int) {
	rand.Seed(time.Now().Unix())
	b := rand.Perm(len(rowData))
	for range b {
		c = b[a]
	}
	return c
}
