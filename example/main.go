package main

import (
	"github.com/chanshik/dumpcode"
	"io/ioutil"
)

func main() {
	if data, e := ioutil.ReadFile("./dumpcode/dumpcode.go"); e == nil {
		dumpcode.DumpCode(([]byte)(data), len(data))
	}
}
