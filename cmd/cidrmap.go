package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Mapping struct {
	cidr  net.IPNet
	value string
}

type Mappings = []Mapping

func cidrmapRun(inputs []net.IP, mappings *Mappings, out io.Writer) error {
	for _, input := range inputs {
		err := checkAddress(input, mappings, out)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkAddress(input net.IP, mappings *Mappings, out io.Writer) error {
	for _, mapping := range *mappings {
		if mapping.cidr.Contains(input) {
			fmt.Fprintf(out, "%s\n", mapping.value)
			return nil
		}
	}

	return fmt.Errorf("no mapping found for %s\n", input)
}

func parseInput(inputs []string) ([]net.IP, error) {
	output := make([]net.IP, 0)
	for _, input := range inputs {
		result := net.ParseIP(input)
		if result == nil {
			return nil, fmt.Errorf("unable to parse input as IP address")
		}
		output = append(output, result)
	}

	return output, nil
}

func loadInputs(stdin *os.File, args []string) ([]net.IP, error) {
	if len(args) > 0 {
		return parseInput(args)
	} else if stat, _ := stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		inputs := make([]string, 0)
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			inputs = append(inputs, scanner.Text())
		}

		return parseInput(inputs)
	} else {
		return nil, fmt.Errorf("no input addresses provided")
	}
}

type mappingDto = map[string]string

func loadMappings(path string, raw string) (*Mappings, error) {
	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		return parseMappings(string(data))
	}

	return parseMappings(raw)
}

func parseMappings(raw string) (*Mappings, error) {
	var dto mappingDto

	err := yaml.Unmarshal([]byte(raw), &dto)
	if err != nil {
		return nil, err
	}

	output := make(Mappings, 0)
	for k, v := range dto {
		_, cidr, err := net.ParseCIDR(k)
		if err != nil {
			return nil, err
		}
		output = append(output, Mapping{cidr: *cidr, value: v})
	}

	return &output, nil
}

var cidrmapCmd = &cobra.Command{
	Use:   "cidrmap [value]",
	Short: "Map IP Address to Value",
	Long:  `Map IP Addresses to values based upon a configured mapping`,
	// Args:  cobra.Arg,
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := loadInputs(os.Stdin, args)
		if err != nil {
			return err
		}

		rawString, err := cmd.Flags().GetString("mapping")
		if err != nil {
			return err
		}
		mappingPath, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}

		mappings, err := loadMappings(mappingPath, rawString)
		if err != nil {
			return err
		}

		err = cidrmapRun(input, mappings, os.Stdout)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cidrmapCmd)

	cidrmapCmd.Flags().String("mapping", "", "YAML formatted CIDR-to-value mapping")
	cidrmapCmd.Flags().String("path", "", "Path to mappings YAML file")
}
