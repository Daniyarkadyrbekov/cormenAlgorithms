package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//На вход программе подается большое количество целых чисел.
// Все числа, кроме одного, имеют пару, причем может быть несколько одинаковых пар.
// Найдите число без пары.

func getArrOfInput(data []byte) []int {

	arrStr := strings.Split(string(data), "\n")
	resArr := make([]int, 0)

	for _, intStr := range arrStr {
		if intStr == "" {
			continue
		}
		i, err := strconv.Atoi(intStr)
		if err != nil {
			panic("atoi " + err.Error())
		}
		resArr = append(resArr, i)
	}
	return resArr
}

func findUniq(arr []int) int {
	uniqMap := make(map[int]bool)

	for _, el := range arr {
		uniq, ok := uniqMap[el]
		if !ok {
			uniqMap[el] = true
			continue
		}
		if uniq {
			uniqMap[el] = false
		}
	}

	for k, v := range uniqMap {
		if v {
			return k
		}
	}

	return 0
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("read os stdin" + err.Error())
	}

	arr := getArrOfInput(data)

	_, err = os.Stdout.WriteString(strconv.Itoa(findUniq(arr)))
	if err != nil {
		panic("stdout" + err.Error())
	}
}
