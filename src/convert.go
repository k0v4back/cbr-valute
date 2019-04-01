package src

import (
	"log"
)

func Convert(from, to string, count int) (result float64, error string) {
	valuteFromForOne, err := GetValuteByName(from)
	valuteToForOne, err := GetValuteByName(to)
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
	if valuteFromForOne == 0.0 || valuteToForOne == 0.0 {
		return 0.0, "This currency not found, read the list of possible currencies"
	}

	transform := valuteFromForOne * float64(count) / valuteToForOne

	return float64(int(transform*100))/100, "OK"
}
