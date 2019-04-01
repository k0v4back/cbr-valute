package main

import (
	"testing"
	"./src"
)

var url = "https://www.cbr.ru/scripts/XML_daily_eng.asp"
var valute = []string{"AUD", "BYN", "BRL", "USD", "KZT"}

func TestGetAllValute(t *testing.T) {
	result, err := src.GetAllValute(url)
	if result == nil || err != nil {
		t.Error(
			"Error in func GetAllValute",
		)
	}
}

func TestGetValuteByName(t *testing.T) {
	for _, v := range valute {
		result, err := src.GetValuteByName(v)
		if err != nil {
			t.Error(
				"Can not get valute - ", v,
			)
		}
		if result == 0.0 {
			t.Error(
				"Can not get valute",
			)
		}
	}
}