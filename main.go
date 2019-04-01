package main

import (
	"./data"
	"fmt"
)

var url = "https://www.cbr.ru/scripts/XML_daily_eng.asp"

func main() {
	result, _ := data.GetValuteByName(url, "USD")
	fmt.Println(result)
}