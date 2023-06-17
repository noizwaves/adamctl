package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type place struct {
	Name     string
	Location *time.Location
}

func getPlaces() (*[]place, error) {
	desired := [](struct {
		Display  string
		Timezone string
	}){
		{"Los Angeles", "America/Los_Angeles"},
		{"Denver", "America/Denver"},
		{"New York", "America/New_York"},
		{"Turkey", "Turkey"},
	}

	cities := make([]place, 0)
	for _, city := range desired {
		location, err := time.LoadLocation(city.Timezone)
		if err != nil {
			return nil, err
		}
		cities = append(cities, place{city.Display, location})
	}

	return &cities, nil
}

func run(out io.Writer, t time.Time) {
	places, err := getPlaces()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", *places)
	fmt.Fprintf(out, "%v\n\n", t.Format(time.UnixDate))
	fmt.Fprintf(out, "UTC: %v\n", t.UTC().Format(time.UnixDate))
	for _, p := range *places {
		fmt.Fprintf(out, "%s: %s\n", p.Name, t.In(p.Location).Format(time.UnixDate))
	}
}

func main() {
	now := time.Now()
	run(os.Stdout, now)
}
