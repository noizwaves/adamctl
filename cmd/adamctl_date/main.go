package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func run(out io.Writer, t time.Time) {
	fmt.Fprintln(out, t.Format(time.UnixDate))
}

func main() {
	now := time.Now()
	run(os.Stdout, now)
}
