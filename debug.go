package sml

import (
	"fmt"
	"math"
)

var sml_debug_enable bool

func sml_debug(buf *sml_buffer, function string) {
	if sml_debug_enable {
		fmt.Printf("%-22s % x\n", function, buf.buf[buf.cursor:buf.cursor+30])
	}
}

func Sml_print_message(msg sml_message) {
	if msg.message_body.tag == SML_MESSAGE_GET_LIST_RESPONSE {
		list, ok := msg.message_body.data.(sml_get_list_response)
		if !ok {
			panic("sml: Could not cast list reponse")
		}

		for _, elem := range list.val_list {
			Sml_print_list_entry(elem)
		}
	}
}

func Sml_print_list_entry(entry sml_list_entry) {
	obis := fmt.Sprintf("%d-%d:%d.%d.%d*%d",
		entry.obj_name[0], entry.obj_name[1],
		entry.obj_name[2], entry.obj_name[3],
		entry.obj_name[4], entry.obj_name[5])

	fmt.Printf("%-22s", obis)

	if (entry.value.typ == SML_TYPE_OCTET_STRING) {
		fmt.Printf("% x\n", entry.value.data_bytes)
	} else if (entry.value.typ == SML_TYPE_BOOLEAN) {
		// fmt.Println("foo")
		fmt.Printf("%v\n", entry.value.data_boolean)
	} else if (((entry.value.typ & SML_TYPE_FIELD) == SML_TYPE_INTEGER) ||
			((entry.value.typ & SML_TYPE_FIELD) == SML_TYPE_UNSIGNED)) {
		scaler := 1
		if entry.scaler != 0 {
			scaler = int(entry.scaler)
		}

		value := float64(entry.value.data_int) * math.Pow10(scaler)

		unit := ""
		switch (entry.unit) {
		case 0x1B:
			unit ="W"
		case 0x1E:
			unit ="Wh"
		}

		fmt.Printf("%12.1f %-3s\n", value, unit)
	}
}
