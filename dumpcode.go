package dumpcode

import (
	. "fmt"
	"strconv"
)

func PrintChar(c byte) {
	r := (rune)(c)

	if strconv.IsPrint(r) {
		Printf("%c", r)
	} else {
		Printf(".")
	}
}

func DumpCode(buffer []byte, buffer_len int) {
	i := 0

	for ; i < buffer_len; i++ {
		if i%16 == 0 {
			Printf("0x%08x  ", &(buffer[i]))
		}

		Printf("%02x ", buffer[i])
		if i%16-15 == 0 {
			Printf("  ")
			for j := i - 15; j <= i; j++ {
				PrintChar(buffer[j])
			}
			Printf("\n")
		}
	}

	if i%16 != 0 {
		spaces := (buffer_len-i+16-i%16)*3 + 2

		for j := 0; j < spaces; j++ {
			Printf(" ")
		}

		for j := i - i%16; j < buffer_len; j++ {
			PrintChar(buffer[j])
		}
	}
	Printf("\n")
}
