package sml

func BooleanParse(buf *Buffer) (bool, error) {
	if BufOptionalIsSkipped(buf) {
		return false, nil
	}

	Debug(buf, "BooleanParse")

	if err := Expect(buf, TYPEBOOLEAN, 1); err != nil {
		return false, err
	}

	if BufGetCurrentByte(buf) > 0 {
		BufUpdateBytesRead(buf, 1)
		return true, nil
	} else {
		BufUpdateBytesRead(buf, 1)
		return false, nil
	}
}
