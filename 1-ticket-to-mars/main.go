package main

import (
	"fmt"
	"math/rand"
)

const distanceToMars = 62100000

var spacelines = []string{
	"Virgin Galactic",
	"SpaceX",
	"Space Adventures",
}

func main() {
	printHeaders()
	for i := 0; i <= 10; i++ {
		var (
			spaceline string
			ticket    string
		)
		speed := rand.Intn(15) + 16
		price := speed + 20
		days := distanceToMars / (speed * 60 * 60 * 24)
		sl := rand.Intn(3)
		switch sl {
		case 0:
			spaceline = spacelines[0]
		case 1:
			spaceline = spacelines[1]
		case 2:
			spaceline = spacelines[2]
		}
		tt := rand.Intn(2)
		switch tt {
		case 0:
			ticket = "Roundtrip"
			price = price * 2
		case 1:
			ticket = "One-way"
		}
		printRow(spaceline, ticket, days, price)
	}
}

func printHeaders() {
	headers := []string{
		"Spaceline",
		"Days",
		"Trip Type",
		"Price",
	}
	s := fmt.Sprintf("%-15v %4v %8v %4v", headers[0], headers[1], headers[2], headers[3])
	div := ""
	for i := 0; i < len(s); i++ {
		div += "="
	}
	fmt.Println(s)
	fmt.Println(div)
}

func printRow(spaceline, trip string, days, price int) {
	r := fmt.Sprintf("%-15v %4v %8v $%4v", spaceline, days, trip, price)
	fmt.Println(r)
}
