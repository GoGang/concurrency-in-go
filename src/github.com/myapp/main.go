package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

const DUMP_DIR = "/home/bchenebault/wikipedia-articles/"

/*
  Count words occurences in wikipedia dump
*/
func main() {
	runtime.GOMAXPROCS(1)

	// Load wikipedia in memory
	fmt.Println("Loading content in memory")
	inmemory := make(map[string][]string)
	files, _ := ioutil.ReadDir(DUMP_DIR)
	for _, f := range files {
		file, _ := os.Open(DUMP_DIR + f.Name())
		lines := bufio.NewScanner(file)
		for lines.Scan() {
			inmemory[f.Name()] = append(inmemory[f.Name()], lines.Text())
		}
	}

	// Start MAP/REDUCE
	fmt.Println("Start indexing wikipedia dump")
	startTime := time.Now()
	index := make(map[string]int)
	channel := make(chan map[string]int, 100)

	for name, content := range inmemory {
		fmt.Println("Start worker for file", name)
		go countWords(content, channel)
	}

	for i := 0; i < len(files); i++ {
		partialIndex := <-channel
		fmt.Println("get one partial index to merge, num of entries :", len(partialIndex))
		for key, value := range partialIndex {
			index[key] = index[key] + value
		}
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

func countWords(content []string, channel chan map[string]int) {
	words := make(map[string]int)
	for _, title := range content {
		for _, w := range strings.Fields(title) {
			words[w]++
		}
	}
	channel <- words
}
