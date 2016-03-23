package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "time"
import "sort"

type Stamp struct {
	timestamp int64
	original string
}

type TimeSlice []Stamp

func (p TimeSlice) Len() int {
	return len(p)
}

func (p TimeSlice) Less(i, j int) bool {
	if p[i].timestamp < p[j].timestamp {
		return true
	}
	return false
}

func (p TimeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Reduce(test string) TimeSlice {
	result := strings.Split(test, " ")
	s := make(TimeSlice, 0, 5)
	for k := range result {
		t, err := time.Parse("15:04:05", result[k])
		s = append(s, Stamp{timestamp: t.Unix(), original: result[k]})
		if err != nil {
			fmt.Println(err)
		}
	}
    return s
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		answer := Reduce(scanner.Text())
		sort.Sort(sort.Reverse(answer))
		for k, v := range answer {
			if k != len(answer)-1 {
			fmt.Printf("%s ",v.original)
		    } else {
		    	fmt.Println(v.original)
		    }
		} 
	}
}
