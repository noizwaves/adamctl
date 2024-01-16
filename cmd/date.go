package cmd

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
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

func location(name string) *time.Location {
	l, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}

	return l
}

func parseDate(s string) (time.Time, error) {
	return dateparse.ParseLocal(s)
}

func formatOffset(offset int) string {
	d := time.Second * time.Duration(offset)
	return fmt.Sprintf("%+2.f:00", math.Round(d.Hours()))
}

func run(out io.Writer, current time.Time, value string, tz string) error {
	t := current
	if value != "" {
		var err error
		if tz != "" {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				return err
			}

			t, err = dateparse.ParseIn(value, loc)
			if err != nil {
				return err
			}
		} else {
			t, err = dateparse.ParseLocal(value)
			if err != nil {
				return err
			}
		}
	}

	places, err := getPlaces()
	if err != nil {
		return err
	}

	tab := table.NewWriter()
	tab.SetTitle("The time in various places")
	tab.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	tab.AppendHeader(table.Row{"Place", "Offset", "Date"})

	_, offsetRaw := t.Zone()
	tab.AppendRow(table.Row{"Raw", formatOffset(offsetRaw), t.Format(time.UnixDate)})

	tLocal := t.In(time.Local)
	_, offsetLocal := tLocal.Zone()
	tab.AppendRow(table.Row{"Local", formatOffset(offsetLocal), tLocal.Format(time.UnixDate)})

	tUtc := t.In(time.UTC)
	_, offsetUtc := tUtc.Zone()
	tab.AppendRow(table.Row{"UTC", formatOffset(offsetUtc), tUtc.Format(time.UnixDate)})

	for _, p := range *places {
		tIn := t.In(p.Location)
		_, offset := tIn.Zone()
		tab.AppendRow(table.Row{p.Name, formatOffset(offset), tIn.Format(time.UnixDate)})
	}

	tab.SetOutputMirror(out)
	tab.Render()

	return nil
}

func getValueArgument(stdin *os.File, args []string) string {
	value := ""
	if len(args) > 0 {
		value = args[0]
	} else if stat, _ := stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err == nil {
			value = strings.TrimSpace(line)
		}
	}
	return value
}

var dateCmd = &cobra.Command{
	Use:   "date [value]",
	Short: "Print information about current date",
	Long: `A general purpose date parser and printer. Shows useful information about the date. By default shows current date.

Optionally override date value used via argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		value := getValueArgument(os.Stdin, args)
		tz, err := cmd.Flags().GetString("tz")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)

			return
		}

		err = run(os.Stdout, now, value, tz)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	},
}

func init() {
	dateCmd.Flags().String("tz", "", "Contextually parse dates in this timezone instead")

	rootCmd.AddCommand(dateCmd)
}
