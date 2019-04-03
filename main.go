package main

import (
	"fmt"
	"github.com/k0v4back/cbr-valute/src"
)

func main() {
	fmt.Println(src.Convert("USD", "RUB", 1))
}