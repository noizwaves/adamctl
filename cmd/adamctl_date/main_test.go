package main

import (
	"bytes"
	"testing"
	"time"
)

var mdt = time.FixedZone("MDT", -6*60*60)

func TestRun(t *testing.T) {
	now := time.Date(2023, time.June, 9, 16, 22, 45, 0, mdt)
	var out bytes.Buffer
	err := run(&out, now, "")

	if err != nil {
		t.Fatal(err)
	}

	expected := `Fri Jun  9 16:22:45 MDT 2023

UTC: Fri Jun  9 22:22:45 UTC 2023
Los Angeles: Fri Jun  9 15:22:45 PDT 2023
Denver: Fri Jun  9 16:22:45 MDT 2023
New York: Fri Jun  9 18:22:45 EDT 2023
Turkey: Sat Jun 10 01:22:45 +03 2023
`
	actual := out.String()

	if expected != actual {
		t.Errorf("Expected %v but received %v", expected, actual)
	}
}

func TestRunValidDateValue(t *testing.T) {
	now := time.Date(2023, time.June, 9, 16, 22, 45, 0, mdt)
	var out bytes.Buffer
	run(&out, now, "Sat Jun 17 14:44:25 PDT 2023")

	expected := `Sat Jun 17 14:44:25 PDT 2023

UTC: Sat Jun 17 21:44:25 UTC 2023
Los Angeles: Sat Jun 17 14:44:25 PDT 2023
Denver: Sat Jun 17 15:44:25 MDT 2023
New York: Sat Jun 17 17:44:25 EDT 2023
Turkey: Sun Jun 18 00:44:25 +03 2023
`
	actual := out.String()

	if expected != actual {
		t.Errorf("Expected %v but received %v", expected, actual)
	}
}
