package main

import (
	"./src"
	"fmt"
)

func main() {
	fmt.Println(src.Convert("USD", "RUB", 1))
	//fmt.Println(src.GetValuteByName("USD"))
}