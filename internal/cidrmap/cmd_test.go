package cidrmap

import (
	"bytes"
	"net"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestCidrmapRun(t *testing.T) {
	var out bytes.Buffer

	inputs := []net.IP{
		net.ParseIP("192.168.0.255"),
		net.ParseIP("192.168.1.0"),
	}

	mappings := &[]Mapping{
		{
			cidr:  *parseCidr(t, "192.168.0.0/24"),
			value: "foo",
		}, {
			cidr:  *parseCidr(t, "192.168.1.0/24"),
			value: "bar",
		},
	}

	format, _ := template.New("test-format").Parse("{{.Value}}")

	err := Run(inputs, mappings, format, &out)

	assert.NoError(t, err)

	assert.Equal(t, "foo\nbar\n", out.String())
}

func parseCidr(t *testing.T, value string) *net.IPNet {
	t.Helper()

	_, cidr, err := net.ParseCIDR(value)
	assert.NoError(t, err)

	return cidr
}
