package sml

import (
	"fmt"
	"math"
)

var DebugEnable bool

func Debug(buf *Buffer, function string) {
	if DebugEnable {
		fmt.Printf("%-22s % x\n", function, buf.Buf[buf.Cursor:buf.Cursor+30])
	}
}

func PrintMessage(msg Message) {
	if msg.MessageBody.Tag == MESSAGEGETLISTRESPONSE {
		list, ok := msg.MessageBody.Data.(GetListResponse)
		if !ok {
			panic("sml: Could not cast list reponse")
		}

		for _, elem := range list.ValList {
			PrintListEntry(elem)
		}
	}
}

func PrintListEntry(entry ListEntry) {
	obis := fmt.Sprintf("%d-%d:%d.%d.%d*%d",
		entry.ObjName[0], entry.ObjName[1],
		entry.ObjName[2], entry.ObjName[3],
		entry.ObjName[4], entry.ObjName[5])

	fmt.Printf("%-22s", obis)

	if entry.Value.Typ == TYPEOCTETSTRING {
		fmt.Printf("% x\n", entry.Value.DataBytes)
	} else if entry.Value.Typ == TYPEBOOLEAN {
		// fmt.Println("foo")
		fmt.Printf("%v\n", entry.Value.DataBoolean)
	} else if ((entry.Value.Typ & TYPEFIELD) == TYPEINTEGER) ||
		((entry.Value.Typ & TYPEFIELD) == TYPEUNSIGNED) {
		scaler := 1
		if entry.Scaler != 0 {
			scaler = int(entry.Scaler)
		}

		value := float64(entry.Value.DataInt) * math.Pow10(scaler)

		unit := ""
		switch entry.Unit {
		case 0x1B:
			unit = "W"
		case 0x1E:
			unit = "Wh"
		}

		fmt.Printf("%12.1f %-3s\n", value, unit)
	}
}
