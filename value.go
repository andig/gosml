package sml

import "fmt"

type Value struct {
	Typ         uint8
	DataBytes   OctetString
	DataBoolean bool
	DataInt     int64
}

func ValueParse(buf *Buffer) (Value, error) {
	/*
		if (BufOptionalIsSkipped(buf)) {
			return 0;
		}

		int max = 1;
		int type = BufGetNextType(buf);
		unsigned char byte = BufGetCurrentByte(buf);

		Value *value = ValueInit();
		value->type = type;

		switch (type) {
			case TYPEOCTETSTRING:
				value->data.bytes = OctetStringParse(buf);
				break;
			case TYPEBOOLEAN:
				value->data.boolean = BooleanParse(buf);
				break;
			case TYPEUNSIGNED:
			case TYPEINTEGER:
				// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
				while (max < ((byte & LENGTHFIELD) - 1)) {
					max <<= 1;
				}

				value->data.uint8 = NumberParse(buf, type, max);
				value->type |= max;
				break;
			default:
				buf->error = 1;
				break;
		}
	*/
	value := Value{}

	if BufOptionalIsSkipped(buf) {
		return value, nil
	}

	Debug(buf, "ValueParse")

	typefield := BufGetNextType(buf)
	b := BufGetCurrentByte(buf)

	max := 1
	value.Typ = typefield

	var err error
	switch typefield {
	case TYPEOCTETSTRING:
		value.DataBytes, err = OctetStringParse(buf)
		if err != nil {
			return value, err
		}
	case TYPEBOOLEAN:
		value.DataBoolean, err = BooleanParse(buf)
		if err != nil {
			return value, err
		}
	case TYPEUNSIGNED:
		// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
		for max < int((b&LENGTHFIELD)-1) {
			max = max << 1
		}

		value.DataInt, err = NumberParse(buf, typefield, max)
		if err != nil {
			return value, err
		}

		value.Typ = value.Typ | uint8(max)
	case TYPEINTEGER:
		// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
		for max < int((b&LENGTHFIELD)-1) {
			max = max << 1
		}

		value.DataInt, err = NumberParse(buf, typefield, max)
		if err != nil {
			return value, err
		}

		value.Typ = value.Typ | uint8(max)
	default:
		return value, fmt.Errorf("Unexpected type %02x", typefield)
	}

	return value, nil
}
