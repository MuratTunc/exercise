package main

import "fmt"

func checkCapitalExist(countires map[string]string, capital string) bool {

	for _, value := range countires {
		if value == capital {
			return true
		}
	}
	return false

}
func main() {

	countries := map[string]string{
		"Turkey":  "Ankara",
		"Germany": "Berlin",
		"Japan":   "Tokyo",
	}

	capital, exist := countries["Germany"]

	if exist {
		fmt.Println("Berlin exist", capital)
	}

	fmt.Println(checkCapitalExist(countries, "Tokyo"))

}
