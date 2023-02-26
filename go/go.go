package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/oschwald/maxminddb-golang"
)

func main() {
	// Ask the user for the path to the MMDB file
	scanner := bufio.NewScanner((bufio.NewReader((os.Stdin))))
	fmt.Println("Enter the path to the MMDB file:")
	scanner.Scan()
	filePath := scanner.Text()

	// Open the MMDB file
	db, err := maxminddb.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ask the user for the IP address to look up
	fmt.Println("Enter the IP address to look up:")
	scanner.Scan()
	ipAddress := scanner.Text()

	// Look up the location data for the IP address
	ip := net.ParseIP(ipAddress)
	var record struct {
		City struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"city"`
		Country struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"country"`
		Location struct {
			Latitude  float64 `maxminddb:"latitude"`
			Longitude float64 `maxminddb:"longitude"`
		} `maxminddb:"location"`
	}
	err = db.Lookup(ip, &record)
	if err != nil {
		panic(err)
	}

	// Print the location information
	fmt.Println("City: " + record.City.Names["en"])
	fmt.Println("Country: " + record.Country.Names["en"])
	fmt.Println("Latitude: " + fmt.Sprintf("%f", record.Location.Latitude))
	fmt.Println("Longitude: " + fmt.Sprintf("%f", record.Location.Longitude))
}