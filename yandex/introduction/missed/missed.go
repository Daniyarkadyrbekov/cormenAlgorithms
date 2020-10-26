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
	numb, err := strconv.Atoi(line[:len(line)-1])
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	line, err = reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		fmt.Printf(err.Error())
		return
	}

	arr := make([]bool, numb)

	numbArr := strings.Split(line, " ")
	for _, nStr := range numbArr {
		n, err := strconv.Atoi(nStr)
		if err != nil {
			fmt.Print(err.Error())
		}
		arr[n-1] = true
	}

	fOut, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	del := " "
	for i := range arr {
		if !arr[i] {
			if _, err := fOut.WriteString(strconv.Itoa(i+1) + del); err != nil {
				fmt.Print(err.Error())
			}
			del = ""
		}
	}
}
