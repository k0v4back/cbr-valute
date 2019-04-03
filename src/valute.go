package src

import (
	"encoding/xml"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type xmlResult struct {
	Name   string   `xml:"name,attr"`
	Valute []valute `xml:"Valute"`
}

type valute struct {
	ID       string  `xml:"ID,attr"`
	NumCode  int64   `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  float64 `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    string  `xml:"Value"`
}

type NowRate struct {
	ID      string
	NumCode int64
	CharCode string
	Name    string
	Value   float64
}

var UrlParse = "https://www.cbr.ru/scripts/XML_daily_eng.asp"

var NowRates map[string]NowRate

func GetAllValute(url string) (map[string]NowRate, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf(err.Error())
	}

	var data xmlResult

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&data)
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}

	NowRates = make(map[string]NowRate)

	for _, el := range data.Valute {
		value, _ := strconv.ParseFloat(strings.Replace(el.Value, ",", ".", 1), 8)

		NowRates[el.CharCode] = NowRate{
			ID:      el.ID,
			NumCode: el.NumCode,
			CharCode: el.CharCode,
			Name:    el.Name,
			Value:   value / el.Nominal,
		}
	}

	return NowRates, nil
}


func GetValuteByName (name string) (float64, error) {
	allValet, err := GetAllValute(UrlParse)
	if err != nil {
		return 0.0, err
	}

	for _, v := range allValet {
		if v.CharCode == name {
			return v.Value, nil
		}
	}

	return 0.0, nil
}
