package main

import "fmt"

func Debts(pays map[string]int) map[string]map[string]float64 {
	sum := 0
	for _, v := range pays {
		sum += v
	}

	var avg = float64(sum) / float64(len(pays))

	fmt.Printf("avg = %v\n", avg)

	dolzhniki := make(map[string]float64)
	zaemshiki := make(map[string]float64)

	for k, v := range pays {
		if float64(v) > avg {
			zaemshiki[k] = float64(v) - avg
		} else {
			dolzhniki[k] = avg - float64(v)
		}
	}

	fmt.Printf("zaemshiki = %v\n", zaemshiki)
	fmt.Printf("dolzhniki = %v\n", dolzhniki)

	res := make(map[string]map[string]float64)

	for zaemName, zaem := range zaemshiki {

		if len(res[zaemName]) == 0 {
			res[zaemName] = make(map[string]float64)
		}

		for dolzhName, dolg := range dolzhniki {
			if zaem < 0.00001 {
				break
			}
			if zaem < dolg {
				res[zaemName][dolzhName] = zaem
				dolzhniki[dolzhName] = dolg - zaem
				break
			} else {
				res[zaemName][dolzhName] = dolg
				zaem = zaem - dolg
			}
		}
	}

	return res
}

func main() {
	//var debts = map[string]int{
	//	"Темир":   100,
	//	"Данияр":  31,
	//	"Мустафа": 2010,
	//}

	var debts = map[string]int{
		"Ангелина": 0,
		"Еламан":   47914,
		"Темир":    47401,
		"Данияр":   45042,
		"Мустафа":  1000,
		"Даниэль":  0,
		"Ильяс":    24710,
	}

	fmt.Printf("result = %v\n", Debts(debts))
}
