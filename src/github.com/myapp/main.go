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

/*
  Count words occurences in wikipedia dump
*/
func main() {
	fmt.Println("Start indexing wikipedia dump")

	startTime := time.Now()
	files, _ := ioutil.ReadDir(DUMP_DIR)
	index := make(map[string]int)

	maps := make(chan map[string]int)

	/**** Reduce (wait for incoming maps) ****/
	go func(maps chan map[string]int) {
		fmt.Println("ffsdfds")
		for i := 0; ; i++ {
			comingMap := <-maps
			fmt.Println(i, "map received")
			for k, v := range comingMap {
				index[k] = index[k] + v
			}
		}
	}(maps)

	/**** Map ****/
	for _, f := range files {
		go countWords(f, maps)
	}

	fmt.Println("All wikipedia articles indexed in : ", time.Since(startTime))

	// Prompt
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

func countWords(f os.FileInfo, maps chan map[string]int) {
	//fmt.Println(DUMP_DIR + f.Name())
	words := make(map[string]int)

	file, _ := os.Open(DUMP_DIR + f.Name())
	lines := bufio.NewScanner(file)
	for lines.Scan() {
		for _, w := range strings.Fields(lines.Text()) {
			words[w]++
		}
	}
	//fmt.Println("Traitment for file", f.Name(), "is finishedr")
	maps <- words
	fmt.Println(f.Name(), "SENT")
}
