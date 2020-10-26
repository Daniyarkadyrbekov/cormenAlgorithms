package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	tracksNumber, err := strconv.Atoi(line[:len(line)-1])
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	if tracksNumber == 0 {
		return
	}

	line1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	line2, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		fmt.Printf(err.Error())
		return
	}

	arr1 := strings.Split(line1[:len(line1)-1], " ")
	arr2 := strings.Split(line2, " ")

	if len(arr1) != len(arr2) && len(arr1) != tracksNumber {
		fmt.Printf("tracks number don't matches")
	}

	fOut, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for i := 0; i < tracksNumber; i++ {
		_, err := fOut.WriteString(arr1[i] + " " + arr2[i] + " ")
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	}

}
