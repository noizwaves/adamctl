package date

import (
	"bytes"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mdt = time.FixedZone("MDT", -6*60*60)

// from https://stackoverflow.com/a/29497680
var ansiRegex = regexp.MustCompile("[\u001b\u009b][[()#;?]*(?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[0-9A-ORZcf-nqry=><]")

func stripAnsi(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

func TestRun(t *testing.T) {
	now := time.Date(2023, time.June, 9, 16, 22, 45, 0, mdt)
	var out bytes.Buffer

	err := Run(&out, now, "", "")
	if err != nil {
		t.Fatal(err)
	}

	actual := stripAnsi(out.String())

	expected := "" +
		" The time in various places                        \n" +
		" PLACE        OFFSET  DATE                         \n" +
		" Raw          -6:00   Fri Jun  9 16:22:45 MDT 2023 \n" +
		" Local        -7:00   Fri Jun  9 15:22:45 PDT 2023 \n" +
		" UTC          +0:00   Fri Jun  9 22:22:45 UTC 2023 \n" +
		" Los Angeles  -7:00   Fri Jun  9 15:22:45 PDT 2023 \n" +
		" Denver       -6:00   Fri Jun  9 16:22:45 MDT 2023 \n" +
		" New York     -4:00   Fri Jun  9 18:22:45 EDT 2023 \n" +
		" Turkey       +3:00   Sat Jun 10 01:22:45 +03 2023 \n"

	assert.Equal(t, expected, actual)
}

func TestRunValidDateValue(t *testing.T) {
	now := time.Date(2023, time.June, 9, 16, 22, 45, 0, mdt)
	var out bytes.Buffer

	err := Run(&out, now, "Sat Jun 17 14:44:25 PDT 2023", "")
	if err != nil {
		t.Fatal(err)
	}

	actual := stripAnsi(out.String())

	expected := "" +
		" The time in various places                        \n" +
		" PLACE        OFFSET  DATE                         \n" +
		" Raw          -7:00   Sat Jun 17 14:44:25 PDT 2023 \n" +
		" Local        -7:00   Sat Jun 17 14:44:25 PDT 2023 \n" +
		" UTC          +0:00   Sat Jun 17 21:44:25 UTC 2023 \n" +
		" Los Angeles  -7:00   Sat Jun 17 14:44:25 PDT 2023 \n" +
		" Denver       -6:00   Sat Jun 17 15:44:25 MDT 2023 \n" +
		" New York     -4:00   Sat Jun 17 17:44:25 EDT 2023 \n" +
		" Turkey       +3:00   Sun Jun 18 00:44:25 +03 2023 \n"

	assert.Equal(t, expected, actual)
}

func TestRunSpecifyTimezone(t *testing.T) {
	now := time.Date(2023, time.June, 9, 16, 22, 45, 0, mdt)
	var out bytes.Buffer

	err := Run(&out, now, "Sat Jun 17 14:44:25 2023", "UTC")
	if err != nil {
		t.Fatal(err)
	}

	actual := stripAnsi(out.String())

	expected := "" +
		" The time in various places                        \n" +
		" PLACE        OFFSET  DATE                         \n" +
		" Raw          +0:00   Sat Jun 17 14:44:25 UTC 2023 \n" +
		" Local        -7:00   Sat Jun 17 07:44:25 PDT 2023 \n" +
		" UTC          +0:00   Sat Jun 17 14:44:25 UTC 2023 \n" +
		" Los Angeles  -7:00   Sat Jun 17 07:44:25 PDT 2023 \n" +
		" Denver       -6:00   Sat Jun 17 08:44:25 MDT 2023 \n" +
		" New York     -4:00   Sat Jun 17 10:44:25 EDT 2023 \n" +
		" Turkey       +3:00   Sat Jun 17 17:44:25 +03 2023 \n"

	assert.Equal(t, expected, actual)
}
