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

	printResult(resultMap)
}

func scanFile(p Params) (map[int]int, error) {
	file, err := os.Open(p.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make(map[int]int)
	tree := &Tree{}

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		diff := p.n - number
		if tree.exists(diff) {
			result[number] = diff
		}
		tree.insert(number)
	}

	return result, nil
}

func printResult(resultMap map[int]int) {
	fmt.Printf("Total pairs: %v\n", len(resultMap))
	for i, v := range resultMap {
		fmt.Printf("%v + %v = %v \n", i, v, i+v)
	}
}
