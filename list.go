package sml

type ListEntry struct {
	ObjName        OctetString
	Status         int64
	ValTime        Time
	Unit           uint8
	Scaler         int8
	Value          Value
	ValueSignature OctetString
}

func ListEntryParse(buf *Buffer) (ListEntry, error) {
	Debug(buf, "ListEntryParse")

	elem := ListEntry{}
	var err error

	if err := Expect(buf, TYPELIST, 7); err != nil {
		return elem, err
	}

	if elem.ObjName, err = OctetStringParse(buf); err != nil {
		return elem, err
	}

	if elem.Status, err = StatusParse(buf); err != nil {
		return elem, err
	}

	if elem.ValTime, err = TimeParse(buf); err != nil {
		return elem, err
	}

	if elem.Unit, err = U8Parse(buf); err != nil {
		return elem, err
	}

	if elem.Scaler, err = I8Parse(buf); err != nil {
		return elem, err
	}

	if elem.Value, err = ValueParse(buf); err != nil {
		return elem, err
	}

	if elem.ValueSignature, err = OctetStringParse(buf); err != nil {
		return elem, err
	}

	return elem, nil
}

func ListParse(buf *Buffer) ([]ListEntry, error) {
	if BufOptionalIsSkipped(buf) {
		return nil, nil
	}

	Debug(buf, "ListParse")

	if err := ExpectType(buf, TYPELIST); err != nil {
		return nil, err
	}

	list := make([]ListEntry, 0)

	elems := BufGetNextLength(buf)

	for elems > 0 {
		elem, err := ListEntryParse(buf)
		if err != nil {
			return nil, err
		}
		list = append(list, elem)
		elems--
	}

	return list, nil
}
