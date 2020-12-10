package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var grp []int

func main() {

	// get number of to-be-workers

	var employees = make(map[string]int)
	if len(os.Args) != 2 {
		fmt.Println("Need #workers!")
		os.Exit(1)
	}

	noWorkers, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	employees = map[string]int{
		"Robert": 30,
		"John": 475,
		"Elly": 1022,
		"Bob": 99,
	}

	ch := make(chan string, len(employees))
	ch1 := make(chan int, len(employees))

	for j := 0; j < noWorkers; j++ {
		go getData(j, ch, ch1, len(employees))

	}

	for i, k := range employees {
		ch <- i
		ch1 <- k
	}

	time.Sleep(1e9)
	printUniqueValue(grp)

}

func getData(j int, ch chan string, ch1 chan int, mapLength int) {

	var input string
	var input1 int
	var i int
	for i < mapLength {

		// fmt.Println(grp)

		input = <-ch
		input1 = <-ch1
		grp = append(grp, j)

		res := input1 / 7
		rem := input1 % 7

		switch rem {
		case 0:
			fmt.Println(input, "has worked", res, "weeks in the company")
		case 1:
			fmt.Println(input, "has worked", res, "weeks and", rem, "day in the company")
		default:
			fmt.Println(input, "has worked", res, "weeks in the company and", rem, "days in the company")
		}

		i++

	}

}

func printUniqueValue(arr []int) {
	//Create a   dictionary of values for each element
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	// fmt.Println(dict)
	fmt.Println()
	fmt.Println("Info:")
	fmt.Printf("Workers count: %d\n", len(dict))
	for name, val := range dict {

		if val == 1 {
			fmt.Printf("Worker#%d -> %d element processed\n", name, val)
			continue 
		}

		fmt.Printf("Worker#%d -> %d elements processed\n", name, val)
		
	}
}
