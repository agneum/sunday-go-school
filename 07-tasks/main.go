package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Params struct {
	path string
	n    int
}

var target Params

func init() {
	var targetNumber string

	flag.StringVar(&target.path, "path", "", "File path")
	flag.StringVar(&targetNumber, "N", "", "N")
	flag.Parse()

	if target.path == "" || targetNumber == "" {
		fmt.Println("Flags are missing")
		return
	}
	target.n, _ = strconv.Atoi(targetNumber)
}

func main() {
	resultMap, err := scanFile(target)

	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Total pairs: %v\n", len(resultMap))
	for i, v := range resultMap {
		fmt.Printf("%v + %v = %v \n", i, v, i+v)
	}
}

func scanFile(p Params) (map[int]int, error) {
	result := make(map[int]int)

	file, err := os.Open(p.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		diff := p.n - number
		for i := range lines {
			if lines[i] == diff {
				result[number] = diff
			}
		}
		lines = append(lines, number)
	}

	return result, nil
}
