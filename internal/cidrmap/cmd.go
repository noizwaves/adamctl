package cidrmap

import (
	"fmt"
	"io"
	"net"
)

type Mapping struct {
	cidr  net.IPNet
	value string
}

type Mappings = []Mapping

type Inputs = []net.IP

func Run(inputs Inputs, mappings *Mappings, out io.Writer) error {
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

func NewMapping(cidr net.IPNet, value string) Mapping {
	return Mapping{cidr, value}
}
