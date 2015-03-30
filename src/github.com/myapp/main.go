package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	DUMP_DIR = "/home/bchenebault/DEV/github/wordCounter/wikipedia-articles/"
	WORD     = "Theory"
)

type WordCount struct {
	word     string
	numOfOcc int
}

/*
  Count words occurences in wikipedia dump
*/
func main() {
	fmt.Println("Start indexing wikipedia dump")

	startTime := time.Now()
	files, _ := ioutil.ReadDir(DUMP_DIR)

	words := make(map[string]int)
	numOfDocs := 0
	for _, f := range files {
		fmt.Println(DUMP_DIR + f.Name())
		file, _ := os.Open(DUMP_DIR + f.Name())
		lines := bufio.NewScanner(file)
		for lines.Scan() {
			numOfDocs++
			for _, w := range strings.Fields(lines.Text()) {
				words[w]++
			}
		}
	}
	fmt.Println(numOfDocs, "wikipedia articles indexed in : ", time.Since(startTime))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Type a word and press RETURN : ")
		word, _ := reader.ReadString('\n')
		word = strings.Replace(word, "\n", "", -1)
		if words[word] == 0 {
			fmt.Println("No occurence found")
		} else {
			fmt.Println("Found ", words[word], " occurence(s)")
		}
	}

}
