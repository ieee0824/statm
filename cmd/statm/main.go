package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ieee0824/statm"
)

var (
	pid = flag.Int("p", 0, "pid")
)

func main() {
	flag.Parse()

	for {
		s := statm.New(*pid)

		if s == nil {
			break
		}
		fmt.Println(s)
		time.Sleep(10 * time.Millisecond)
	}
}
