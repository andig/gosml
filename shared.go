package sml

import "fmt"

const (
	MESSAGEEND = 0x00

	TYPEFIELD   = 0x70
	LENGTHFIELD = 0x0F
	ANOTHERTL   = 0x80

	TYPEOCTETSTRING = 0x00
	TYPEBOOLEAN     = 0x40
	TYPEINTEGER     = 0x50
	TYPEUNSIGNED    = 0x60
	TYPELIST        = 0x70

	OPTIONALSKIPPED = 0x01
)

type Buffer struct {
	Bytes  []byte
	Cursor int
}

func BufGetCurrentByte(buf *Buffer) byte {
	return buf.Bytes[buf.Cursor]
}

func BufUpdateBytesRead(buf *Buffer, delta int) {
	buf.Cursor += delta
}

func Expect(buf *Buffer, expectedType uint8, expectedLength int) error {
	if err := ExpectType(buf, expectedType); err != nil {
		return err
	}

	if length := BufGetNextLength(buf); length != expectedLength {
		return fmt.Errorf("Invalid length: %d (expected %d)", length, expectedLength)
	}

	return nil
}

func ExpectType(buf *Buffer, expectedType uint8) error {
	if typefield := BufGetNextType(buf); typefield != expectedType {
		return fmt.Errorf("Unexpected type %02x (expected %02x)", typefield, expectedType)
	}

	return nil
}

func BufGetNextType(buf *Buffer) uint8 {
	return BufGetCurrentByte(buf) & TYPEFIELD
}

func BufGetNextLength(buf *Buffer) int {
	var length uint8
	var list int

	b := BufGetCurrentByte(buf)

	// not a list
	if b&TYPEFIELD != TYPELIST {
		list = -1
	}

	for {
		b := BufGetCurrentByte(buf)

		length = length << 4
		length = length | (b & LENGTHFIELD)

		if b&ANOTHERTL != ANOTHERTL {
			break
		}

		// another TL field used
		BufUpdateBytesRead(buf, 1)

		// not a list
		if list != 0 {
			list--
		}
	}

	BufUpdateBytesRead(buf, 1)

	return int(length) + list
}

func BufOptionalIsSkipped(buf *Buffer) bool {
	if BufGetCurrentByte(buf) == OPTIONALSKIPPED {
		BufUpdateBytesRead(buf, 1)
		return true
	}

	return false
}
