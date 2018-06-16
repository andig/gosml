package sml

import (
	"fmt"
)

var DebugEnable bool

func Debug(buf *Buffer, function string) {
	if DebugEnable {
		fmt.Printf("%-22s % x\n", function, buf.Bytes[buf.Cursor:buf.Cursor+30])
	}
}
