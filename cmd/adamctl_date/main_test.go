package main

import (
	"bytes"
	"testing"
	"time"
)

var mdt = time.FixedZone("MDT", -7*60*60)

func TestRun(t *testing.T) {
	now := time.Date(2023, time.June, 9, 16, 22, 45, 0, mdt)
	var out bytes.Buffer
	run(&out, now)

	expected := "Fri Jun  9 16:22:45 MDT 2023\n"
	actual := out.String()

	if expected != actual {
		t.Errorf("Expected %v but received %v", expected, actual)
	}
}
