package sml

func BooleanParse(buf *Buffer) (bool, error) {
	if BufOptionalIsSkipped(buf) {
		return false, nil
	}

	Debug(buf, "BooleanParse")

	if err := Expect(buf, TYPEBOOLEAN, 1); err != nil {
		return false, err
	}

	b := BufGetCurrentByte(buf)
	BufUpdateBytesRead(buf, 1)
	return b > 0, nil
}
