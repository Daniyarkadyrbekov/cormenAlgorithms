package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	forMain()
}

func forMain() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	resArr := make([]string, n*n)
	counter := 0
	for scanner.Scan() {
		resArr[counter] = scanner.Text()
		counter++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fOut, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fOut.Close()

	mid := n / 2
	var i, j = mid, mid
	counter = 0
	for {

		counter += 1
		//up
		for st := 1; st <= counter; st++ {
			if i == 0 && j == 0 {
				fOut.WriteString(resArr[i*n+j] + "\n")
				return
			}
			fOut.WriteString(resArr[i*n+j] + "\n")
			i -= 1
		}

		//right
		for st := 1; st <= counter; st++ {
			fOut.WriteString(resArr[i*n+j] + "\n")
			j += 1
		}

		counter += 1
		//down
		for st := 1; st <= counter; st++ {
			fOut.WriteString(resArr[i*n+j] + "\n")
			i += 1
		}

		//left
		for st := 1; st <= counter; st++ {
			fOut.WriteString(resArr[i*n+j] + "\n")
			j -= 1
		}
	}
}
