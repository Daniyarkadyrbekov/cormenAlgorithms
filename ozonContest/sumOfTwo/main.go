package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Дано целое положительное число "target". Также дана последовательность из целых положительных чисел.
//Необходимо записать в выходной файл "1", если в последовательности есть два числа сумма,
// которых равна значению "target" или "0" если таких нет

func parseBytes(b []byte) int {
	strInput := strings.Split(string(b), "\n")
	if len(strInput) < 2 {
		fmt.Printf("strInput = %v len = %v\n", strInput, len(strInput))
		panic("strInput< 2")
	}

	sum, err := strconv.Atoi(strInput[0])
	if err != nil {
		panic("sum read err:" + err.Error())
	}

	sumsArr := make(map[int]struct{}, 0)

	strs := strings.Split(strInput[len(strInput)-1], " ")
	for _, str := range strs {
		strInt, err := strconv.Atoi(str)
		if err != nil {
			panic("strArr read err:" + err.Error())
		}

		_, ok := sumsArr[strInt]
		if ok {
			return 1
		}
		sumsArr[sum-strInt] = struct{}{}
	}

	return 0
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic("open file" + err.Error())
	}

	b, err := ioutil.ReadAll(file)

	//err = file.Close()
	//if err != nil {
	//	panic("close err: " + err.Error())
	//}

	result := parseBytes(b)

	f, err := os.Create("./output.txt")
	if err != nil {
		panic("create err: " + err.Error())
	}

	_, err = f.Write([]byte(strconv.Itoa(result)))
	if err != nil {
		panic("write err: " + err.Error())
	}

	//err = f.Close()
	//if err != nil {
	//	panic("close err: " + err.Error())
	//}
}
