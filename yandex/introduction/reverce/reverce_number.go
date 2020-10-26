package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	d, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	numb := strings.TrimSpace(string(d))
	if len(numb) == 0 {
		return
	}

	res := ""

	if numb[0] == '-' {
		res += "-"
		numb = numb[1:]
	}

	start := true
	for i := len(numb) - 1; i >= 0; i-- {
		if numb[i] == '0' && start && len(numb) != 1 {
			continue
		}
		start = false
		res += string(numb[i])
	}

	fOut, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	if _, err = fOut.WriteString(res); err != nil {
		fmt.Printf(err.Error())
	}
}
