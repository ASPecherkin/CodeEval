package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

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

func stepwiseWord(word string) (result string) {
	for k := 0; k < len(word); k++ {
		tmp := ""
		for i := 0; i < k; i++ {
			tmp = "*" + tmp
		}
		result = result + tmp + string(word[k]) + " "
	}
	return
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := maxLenWord(strings.Split(scanner.Text(), " "))
		result := stepwiseWord(word)
		fmt.Println(result[:len(result)-1])
	}
}
