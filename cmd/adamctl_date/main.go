package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func run(out io.Writer, t time.Time) {
	losAngeles, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	denver, err := time.LoadLocation("America/Denver")
	if err != nil {
		panic(err)
	}
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	turkey, err := time.LoadLocation("Turkey")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(out, "%v\n\n", t.Format(time.UnixDate))

	fmt.Fprintf(out, "UTC: %v\n", t.UTC().Format(time.UnixDate))
	fmt.Fprintf(out, "Los Angeles: %v\n", t.In(losAngeles).Format(time.UnixDate))
	fmt.Fprintf(out, "Denver: %v\n", t.In(denver).Format(time.UnixDate))
	fmt.Fprintf(out, "New York: %v\n", t.In(newYork).Format(time.UnixDate))
	fmt.Fprintf(out, "Turkey: %v\n", t.In(turkey).Format(time.UnixDate))
}

func main() {
	now := time.Now()
	run(os.Stdout, now)
}
