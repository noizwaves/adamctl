package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type cityTimezone struct {
	Name     string
	Location *time.Location
}

func getCities() (*[]cityTimezone, error) {
	desired := [](struct {
		Display  string
		Timezone string
	}){
		{"Los Angeles", "America/Los_Angeles"},
		{"Denver", "America/Denver"},
		{"New York", "America/New_York"},
		{"Turkey", "Turkey"},
	}

	cities := make([]cityTimezone, 0)
	for _, city := range desired {
		location, err := time.LoadLocation(city.Timezone)
		if err != nil {
			return nil, err
		}
		cities = append(cities, cityTimezone{city.Display, location})
	}

	return &cities, nil
}

func run(out io.Writer, t time.Time) {
	cities, err := getCities()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", *cities)
	fmt.Fprintf(out, "%v\n\n", t.Format(time.UnixDate))
	fmt.Fprintf(out, "UTC: %v\n", t.UTC().Format(time.UnixDate))
	for _, c := range *cities {
		fmt.Fprintf(out, "%s: %s\n", c.Name, t.In(c.Location).Format(time.UnixDate))
	}
}

func main() {
	now := time.Now()
	run(os.Stdout, now)
}
