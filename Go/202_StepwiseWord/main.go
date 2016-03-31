package main

import "log"
import "bufio"
import "os"
import "strings"
import "sync"
import "fmt"

func maxLenWord(words []string) (longest string) {
	max := 0
	longest = words[0]
	for k := range words {
		if len(words[k]) > max {
			max, longest = len(words[k]), words[k]
		}
	}
	return
}

func generateReuslt(word string) (answer string) {
	for k := 0; k < len(word); k++ {
		tmp := ""
		for i := 0; i < k; i++ {
			tmp = strings.Join(append([]string{"*", tmp}), "")
		}
		answer = strings.Join(append([]string{answer, tmp, string(word[k]), " "}), "")
	}
	return
}

func stepwiseWord(wg *sync.WaitGroup, words chan string, out chan string) {
	defer wg.Done()
	result := ""
	select {
	case word := <-words:
		if word == "done" {
			return
		}
		longest := maxLenWord(strings.Split(word, " "))
		result = generateReuslt(longest)
		out <- result
	}
}

func writeResults(wg *sync.WaitGroup, out chan string) {
	defer wg.Done()
	for i := range out {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	in := make(chan string)
	out := make(chan string)
	go writeResults(&wg, out)
	for scanner.Scan() {
		wg.Add(1)
		go stepwiseWord(&wg, in, out)
		in <- scanner.Text()
	}
	wg.Wait()
}
