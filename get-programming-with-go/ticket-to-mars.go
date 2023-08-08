package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type TripType string

const (
	RoundTrip TripType = "Round-trip"
	OneWay    TripType = "One-way"
)

func main() {
	printTable()

}

func getSpaceline() string {
	spacelineNames := [3]string{"Virgin Galactic", "SpaceX", "SpaceAdventures"}
	return spacelineNames[rand.Intn(len(spacelineNames))]
}

func getSpeedAndDuration() (speed int, duration int) {
	const numberOfSecsInDay = float64(86400) // 24 * 60 * 60
	const distance = float64(62100000)       // km
	speed = int(rand.Intn(15) + 16)          //km/s
	duration = int(math.Round(distance / (float64(speed) * numberOfSecsInDay)))
	return speed, duration
}

func getTripType() TripType {
	tripTypes := make([]TripType, 0)
	tripTypes = append(tripTypes, RoundTrip, OneWay)
	return tripTypes[rand.Intn(len(tripTypes))]
}

func getPrice(speed int, tripType TripType) int {
	price := rand.Intn(15) + 36 // million dollars
	if tripType == RoundTrip {
		price *= 2
	}
	return price
}

func printTable() {
	// Header
	fmt.Printf("%-15v %-4v %-10v %-5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("%v", strings.Repeat("=", 37)+"\n")

	// Rows
	for i := 0; i < 10; i++ {
		spaceline := getSpaceline()
		speed, duration := getSpeedAndDuration()
		tripType := getTripType()
		price := getPrice(speed, tripType)

		fmt.Printf("%-15v %4v %-10v $%4v\n", spaceline, duration, tripType, price)
	}
}
