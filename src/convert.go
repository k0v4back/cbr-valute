package src

import (
	"log"
)

func Convert(from, to string, count int) (result float64) {
	if from == "RUB"{
		result := ConvertFromRubles(from, to, count)
		return float64(int(result*100))/100
	}
	if to == "RUB"{
		result := ConvertByRubles(from, to, count)
		return float64(int(result*100))/100
	}

	valuteFromForOne, err := GetValuteByName(from)
	valuteToForOne, err := GetValuteByName(to)
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
	if valuteFromForOne == 0.0 || valuteToForOne == 0.0 {
		return 0.0
	}

	transform := valuteFromForOne * float64(count) / valuteToForOne

	return float64(int(transform*100))/100
}

func ConvertByRubles(from, _ string, count int) (result float64) {
	valuteFromForOne, err := GetValuteByName(from)
	valuteToForOne := 1.0
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
	if valuteFromForOne == 0.0 {
		return 0.0
	}

	transform := valuteFromForOne * float64(count) / valuteToForOne

	return transform
}

func ConvertFromRubles(_, to string, count int) (result float64) {
	valuteFromForOne := 1.0
	valuteToForOne, err := GetValuteByName(to)
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
	if valuteToForOne == 0.0 {
		return 0.0
	}

	transform := valuteFromForOne * float64(count) / valuteToForOne

	return transform
}
