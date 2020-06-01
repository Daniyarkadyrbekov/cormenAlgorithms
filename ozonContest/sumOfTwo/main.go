package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func parseBytes(f *os.File) int {

	br := bufio.NewReader(f)

	sum := 0
	sumIsSet := false
	sumsArr := make(map[int]struct{}, 0)

	negative := false
	digit := make([]byte, 0)

	for {
		b, err := br.ReadByte()
		if err != nil && err != io.EOF {
			panic("reader err: " + err.Error())
		}
		if err == io.EOF {
			return 0
		}

		if !isSep(b) {
			if string(b) == "-" {
				negative = !negative
				continue
			}
			digit = append(digit, b)
			continue
		}
		if len(digit) == 0 {
			continue
		}
		digitInt, err := strconv.Atoi(string(digit))
		if err != nil {
			panic("get digitINt err " + err.Error())
		}
		digit = digit[:0]
		if !sumIsSet {
			sum = digitInt
			sumIsSet = true
			if negative {
				sum = -sum
				negative = false
			}
			continue
		}

		strInt := digitInt

		if negative {
			strInt = -strInt
			negative = false
		}

		fmt.Printf("sum = %d readedInt = %v\n", sum, strInt)
		_, ok := sumsArr[strInt]
		if ok {
			return 1
		}
		sumsArr[sum-strInt] = struct{}{}
	}

	return 0
}

func isSep(b byte) bool {

	str := string(b)

	return str == " " || str == "\t" || str == "\n"
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic("open file" + err.Error())
	}

	result := parseBytes(file)

	if err := ioutil.WriteFile("./output.txt", []byte(strconv.Itoa(result)), 0666); err != nil {
		panic("error writing out err:" + err.Error())
	}
}
