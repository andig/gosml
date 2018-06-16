package sml

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

const (
	TYPENUMBER_8  = 1
	TYPENUMBER_16 = 2
	TYPENUMBER_32 = 4
	TYPENUMBER_64 = 8
)

func U8Parse(buf *Buffer) (uint8, error) {
	num, err := NumberParse(buf, TYPEUNSIGNED, TYPENUMBER_8)
	return uint8(num), err
}

func U16Parse(buf *Buffer) (uint16, error) {
	num, err := NumberParse(buf, TYPEUNSIGNED, TYPENUMBER_16)
	return uint16(num), err
}

func U32Parse(buf *Buffer) (uint32, error) {
	num, err := NumberParse(buf, TYPEUNSIGNED, TYPENUMBER_32)
	return uint32(num), err
}

func U64Parse(buf *Buffer) (uint64, error) {
	num, err := NumberParse(buf, TYPEUNSIGNED, TYPENUMBER_64)
	return uint64(num), err
}

func I8Parse(buf *Buffer) (int8, error) {
	num, err := NumberParse(buf, TYPEINTEGER, TYPENUMBER_8)
	return int8(num), err
}

func I16Parse(buf *Buffer) (int16, error) {
	num, err := NumberParse(buf, TYPEINTEGER, TYPENUMBER_16)
	return int16(num), err
}

func I32Parse(buf *Buffer) (int32, error) {
	num, err := NumberParse(buf, TYPEINTEGER, TYPENUMBER_32)
	return int32(num), err
}

func I64Parse(buf *Buffer) (int64, error) {
	num, err := NumberParse(buf, TYPEINTEGER, TYPENUMBER_64)
	return int64(num), err
}

func NumberParse(buf *Buffer, numtype uint8, maxSize int) (int64, error) {
	if skip := BufOptionalIsSkipped(buf); skip {
		return 0, nil
	}

	Debug(buf, "NumberParse")

	typefield := BufGetNextType(buf)
	if typefield != numtype {
		return 0, errors.Errorf("Unexpected type %02x (expected %02x)", typefield, numtype)
	}

	length := BufGetNextLength(buf)
	if length < 0 || length > maxSize {
		return 0, errors.Errorf("Invalid length: %d", length)
	}

	np := make([]byte, maxSize)
	missingBytes := maxSize - length

	for i := 0; i < length; i++ {
		np[missingBytes+i] = buf.Bytes[buf.Cursor+i]
	}

	negativeInt := typefield == TYPEINTEGER && (typefield&128 > 0)
	if negativeInt {
		for i := 0; i < missingBytes; i++ {
			np[i] = 0xFF
		}
	}

	var num int64
	switch maxSize {
	case TYPENUMBER_8:
		num = int64(np[0])
	case TYPENUMBER_16:
		num = int64(binary.BigEndian.Uint16(np))
	case TYPENUMBER_32:
		num = int64(binary.BigEndian.Uint32(np))
	case TYPENUMBER_64:
		num = int64(binary.BigEndian.Uint64(np))
	default:
		return num, errors.Errorf("Invalid number type size %02x", maxSize)
	}

	BufUpdateBytesRead(buf, length)
	// fmt.Printf("num: %d\n", num)

	return num, nil
}
