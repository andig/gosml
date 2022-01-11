package sml

import "fmt"

type OctetString []byte

func OctetStringParse(buf *Buffer) (OctetString, error) {
	if skip := BufOptionalIsSkipped(buf); skip {
		return nil, nil
	}

	Debug(buf, "OctetStrParse")

	if err := ExpectType(buf, TYPEOCTETSTRING); err != nil {
		return nil, err
	}

	length := BufGetNextLength(buf)
	if length < 0 {
		return nil, fmt.Errorf("Invalid octet string length %d", length)
	}

	str := buf.Bytes[buf.Cursor : buf.Cursor+length]
	BufUpdateBytesRead(buf, length)

	return str, nil
}
