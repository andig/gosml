package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"

	sml "github.com/andig/gosml"
)

const (
	PRINTRANGE = 46
)

func PrintMessage(msg sml.Message) {
	list, ok := msg.MessageBody.Data.(sml.GetListResponse)
	if !ok {
		panic("Could not cast list reponse")
	}

	for _, elem := range list.ValList {
		PrintListEntry(elem)
	}
}

func Octet2Obis(o sml.OctetString) string {
	return fmt.Sprintf("%d-%d:%d.%d.%d*%d", o[0], o[1], o[2], o[3], o[4], o[5])
}

func ListEntry2Float(entry sml.ListEntry) float64 {
	scaler := 0
	if entry.Scaler != 0 {
		scaler = int(entry.Scaler)
	}

	return float64(entry.Value.DataInt) * math.Pow10(scaler)
}

func PrintListEntry(entry sml.ListEntry) {
	obis := Octet2Obis(entry.ObjName)
	fmt.Printf("%-22s", obis)

	if entry.Value.Typ == sml.TYPEOCTETSTRING {
		fmt.Printf("% x\n", entry.Value.DataBytes)
	} else if entry.Value.Typ == sml.TYPEBOOLEAN {
		fmt.Printf("%v\n", entry.Value.DataBoolean)
	} else if ((entry.Value.Typ & sml.TYPEFIELD) == sml.TYPEINTEGER) || ((entry.Value.Typ & sml.TYPEFIELD) == sml.TYPEUNSIGNED) {
		value := ListEntry2Float(entry)

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	for _, f := range os.Args[1:] {
		fmt.Println(f)

		f, err := os.Open(f)
		check(err)

		r := bufio.NewReader(f)

		for {
			buf, err := sml.TransportRead(r)
			if err == io.EOF {
				break
			}
			check(err)

			// parse without escape sequence/ begin/end marker
			messages, err := sml.FileParse(buf[8 : len(buf)-16])

			for _, msg := range messages {
				if msg.MessageBody.Tag == sml.MESSAGEGETLISTRESPONSE {
					PrintMessage(msg)
				}
			}

			if err != nil {
				fmt.Printf("%+v\n", err)
				goto nextfile
			}
		}
	nextfile:
		// return
	}
}
