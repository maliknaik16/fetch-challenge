package main

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func calculatePoints(receipt Receipt) int {
	points := 0

	for _, ch := range receipt.Retailer {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			points += 1
		}
	}

	total, _ := strconv.ParseFloat(receipt.Total, 64)

	if math.Round(total*100)/100 == math.Round(total) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += int(math.Floor(float64(len(receipt.Items)/2))) * 5

	for _, value := range receipt.Items {
		if len(strings.TrimSpace(value.ShortDescription))%3 == 0 {

			price, _ := strconv.ParseFloat(value.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	var layout = "2006-01-02"
	date, err := time.Parse(layout, receipt.PurchaseDate)

	if err != nil {
		panic(err)
	}

	if date.Day()%2 == 1 {
		points += 6
	}

	var timeLayout = "15:04"
	timeDate, err := time.Parse(timeLayout, receipt.PurchaseTime)

	if err != nil {
		panic(err)
	}

	if timeDate.Hour() >= 14 && timeDate.Hour() <= 16 {
		points += 10
	}

	return points
}
