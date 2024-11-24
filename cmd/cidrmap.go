package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"text/template"

	"github.com/noizwaves/adamctl/internal/cidrmap"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

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

func loadMappings(path string, raw string) (*cidrmap.Mappings, error) {
	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		return parseMappings(string(data))
	}

	return parseMappings(raw)
}

func parseMappings(raw string) (*cidrmap.Mappings, error) {
	var dto mappingDto

	err := yaml.Unmarshal([]byte(raw), &dto)
	if err != nil {
		return nil, err
	}

	output := make(cidrmap.Mappings, 0)
	for k, v := range dto {
		_, cidr, err := net.ParseCIDR(k)
		if err != nil {
			return nil, err
		}
		output = append(output, cidrmap.NewMapping(*cidr, v))
	}

	return &output, nil
}

func loadFormat(cmd *cobra.Command) (*template.Template, error) {
	raw, err := cmd.Flags().GetString("format")
	if err != nil {
		return nil, err
	}

	t, err := template.New("output").Parse(raw)
	if err != nil {
		return nil, err
	}

	return t, nil
}

var cidrmapCmd = &cobra.Command{
	Use:          "cidrmap [value]",
	Short:        "Map IP Address to Value",
	Long:         `Map IP Addresses to values based upon a configured mapping`,
	SilenceUsage: true,
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

		format, err := loadFormat(cmd)
		if err != nil {
			return err
		}

		err = cidrmap.Run(input, mappings, format, os.Stdout)
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
	cidrmapCmd.Flags().String("format", "{{.IP}}: {{.Value}}", "Output format string as a Go template (available fields: IP, Value)")
}
