package cidrmap

import (
	"fmt"
	"io"
	"net"
	"text/template"
)

type Mapping struct {
	cidr  net.IPNet
	value string
}

type Mappings = []Mapping

type Inputs = []net.IP

func Run(inputs Inputs, mappings *Mappings, format *template.Template, out io.Writer) error {
	for _, input := range inputs {
		err := checkAddress(input, mappings, format, out)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkAddress(input net.IP, mappings *Mappings, format *template.Template, out io.Writer) error {
	for _, mapping := range *mappings {
		if mapping.cidr.Contains(input) {
			data := struct {
				IP    string
				Value string
			}{input.String(), mapping.value}

			err := format.Execute(out, data)
			if err != nil {
				return err
			}

			fmt.Fprintln(out)

			return nil
		}
	}

	return fmt.Errorf("no mapping found for %s\n", input)
}

func NewMapping(cidr net.IPNet, value string) Mapping {
	return Mapping{cidr, value}
}
