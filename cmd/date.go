package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/noizwaves/adamctl/internal/date"
	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date [value]",
	Short: "Print information about current date",
	Long: `A general purpose date parser and printer. Shows useful information about the date. By default shows current date.

Optionally override date value used via argument.`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		value := getValueArgument(os.Stdin, args)
		tz, err := cmd.Flags().GetString("tz")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)

			return
		}

		err = date.Run(os.Stdout, now, value, tz)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	},
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

func init() {
	dateCmd.Flags().String("tz", "", "contextually parse dates in this timezone instead")

	rootCmd.AddCommand(dateCmd)
}
