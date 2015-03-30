package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DUMP_DIR = "/home/bchenebault/DEV/github/wordCounter/wikipedia-articles/"
	WORD     = "Theory"
)

/*
  Count words occurences in wikipedia dump
*/
func main() {
	fmt.Println("Start indexing wikipedia dump")

	startTime := time.Now()
	files, _ := ioutil.ReadDir(DUMP_DIR)
	index := make(map[string]int)

	wordsChannel := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(len(files))

	// Start counter
	go func(wc chan string) {
		for {
			index[<-wc]++
		}
	}(wordsChannel)

	// Start readers
	for _, f := range files {
		go countWords(f, wordsChannel, wg)
	}

	fmt.Println("All wikipedia articles indexed in : ", time.Since(startTime))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Type a word and press RETURN : ")
		word, _ := reader.ReadString('\n')
		word = strings.Replace(word, "\n", "", -1)
		if index[word] == 0 {
			fmt.Println("No occurence found")
		} else {
			fmt.Println("Found ", index[word], " occurence(s)")
		}
	}
}

func countWords(f os.FileInfo, wc chan string, wg sync.WaitGroup) {
	fmt.Println(DUMP_DIR + f.Name())
	file, _ := os.Open(DUMP_DIR + f.Name())
	lines := bufio.NewScanner(file)
	for lines.Scan() {
		for _, w := range strings.Fields(lines.Text()) {
			wc <- w
		}
	}
	wg.Done()
}
