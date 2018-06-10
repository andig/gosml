package sml

import (
	"encoding/binary"
	"fmt"
)

const (
	SML_TYPE_NUMBER_8  = 1
	SML_TYPE_NUMBER_16 = 2
	SML_TYPE_NUMBER_32 = 4
	SML_TYPE_NUMBER_64 = 8
)

func sml_u8_parse(buf *sml_buffer) (uint8, error) {
	num, err := sml_number_parse(buf, SML_TYPE_UNSIGNED, SML_TYPE_NUMBER_8)
	return uint8(num), err
}

func sml_u16_parse(buf *sml_buffer) (uint16, error) {
	num, err := sml_number_parse(buf, SML_TYPE_UNSIGNED, SML_TYPE_NUMBER_16)
	return uint16(num), err
}

func sml_u32_parse(buf *sml_buffer) (uint32, error) {
	num, err := sml_number_parse(buf, SML_TYPE_UNSIGNED, SML_TYPE_NUMBER_32)
	return uint32(num), err
}

func sml_u64_parse(buf *sml_buffer) (uint64, error) {
	num, err := sml_number_parse(buf, SML_TYPE_UNSIGNED, SML_TYPE_NUMBER_64)
	return uint64(num), err
}

func sml_i8_parse(buf *sml_buffer) (int8, error) {
	num, err := sml_number_parse(buf, SML_TYPE_INTEGER, SML_TYPE_NUMBER_8)
	return int8(num), err
}

func sml_i16_parse(buf *sml_buffer) (int16, error) {
	num, err := sml_number_parse(buf, SML_TYPE_INTEGER, SML_TYPE_NUMBER_16)
	return int16(num), err
}

func sml_i32_parse(buf *sml_buffer) (int32, error) {
	num, err := sml_number_parse(buf, SML_TYPE_INTEGER, SML_TYPE_NUMBER_32)
	return int32(num), err
}

func sml_i64_parse(buf *sml_buffer) (int64, error) {
	num, err := sml_number_parse(buf, SML_TYPE_INTEGER, SML_TYPE_NUMBER_64)
	return int64(num), err
}

func sml_number_parse(buf *sml_buffer, numtype uint8, max_size int) (int64, error) {
	if skip := sml_buf_optional_is_skipped(buf); skip {
		return 0, nil
	}

	sml_debug(buf, "sml_number_parse")

	typefield := sml_buf_get_next_type(buf)
	if typefield != numtype {
		return 0, fmt.Errorf("sml: Unexpected type %02x (expected %02x)", typefield, numtype)
	}

	length := sml_buf_get_next_length(buf)
	if length < 0 || length > max_size {
		return 0, fmt.Errorf("sml: Invalid length: %d", length)
	}

	np := make([]byte, max_size)
	missing_bytes := max_size - length

	for i := 0; i < length; i++ {
		np[missing_bytes+i] = buf.buf[buf.cursor+i]
	}

	negative_int := typefield == SML_TYPE_INTEGER && (typefield & 128 > 0)
	if negative_int {
		for i := 0; i < missing_bytes; i++ {
			np[i] = 0xFF
		}
	}

	// fmt.Printf("np:  % x\n", np)

	var num int64
	switch max_size {
	case SML_TYPE_NUMBER_8:
		num = int64(np[0])
	case SML_TYPE_NUMBER_16:
		num = int64(binary.BigEndian.Uint16(np))
	case SML_TYPE_NUMBER_32:
		num = int64(binary.BigEndian.Uint32(np))
	case SML_TYPE_NUMBER_64:
		num = int64(binary.BigEndian.Uint64(np))
	default:
		panic("sml: Invalid number type. This should not happen. Please open an issue.")
	}

	sml_buf_update_bytes_read(buf, length)
	// fmt.Printf("num: %d\n", num)

	return num, nil
}
