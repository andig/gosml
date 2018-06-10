package sml

import (
	"fmt"
)

const (
	SML_MESSAGE_END = 0x00

	SML_TYPE_FIELD   = 0x70
	SML_LENGTH_FIELD = 0x0F
	SML_ANOTHER_TL   = 0x80

	SML_TYPE_OCTET_STRING = 0x00
	SML_TYPE_BOOLEAN      = 0x40
	SML_TYPE_INTEGER      = 0x50
	SML_TYPE_UNSIGNED     = 0x60
	SML_TYPE_LIST         = 0x70

	SML_OPTIONAL_SKIPPED = 0x01
)

type sml_buffer struct {
	buf []byte
	cursor int
}

func sml_buf_get_current_byte(buf *sml_buffer) byte {
	return buf.buf[buf.cursor]
}

func sml_buf_update_bytes_read(buf *sml_buffer, delta int) {
	buf.cursor += delta
}

func sml_expect(buf *sml_buffer, expected_type uint8, expected_length int) error {
	if err := sml_expect_type(buf, expected_type); err != nil {
		return err
	}

	if length := sml_buf_get_next_length(buf); length != expected_length {
		return fmt.Errorf("sml: Invalid length: %d (expected %d)", length, expected_length)
	}

	return nil
}

func sml_expect_type(buf *sml_buffer, expected_type uint8) error {
	if typefield := sml_buf_get_next_type(buf); typefield != expected_type {
		return fmt.Errorf("sml: Unexpected type %02x (expected %02x)", typefield, expected_type)
	}

	return nil
}

func sml_buf_get_next_type(buf *sml_buffer) uint8 {
	return sml_buf_get_current_byte(buf) & SML_TYPE_FIELD
}

func sml_buf_get_next_length(buf *sml_buffer) int {
	var length uint8
	var list int

	b := sml_buf_get_current_byte(buf)

	// not a list
	if b&SML_TYPE_FIELD != SML_TYPE_LIST {
		list = -1
	}

	for {
		b := sml_buf_get_current_byte(buf)

		length = length << 4
		length = length | (b & SML_LENGTH_FIELD)

		if b&SML_ANOTHER_TL != SML_ANOTHER_TL {
			break
		}

		// another TL field used
		sml_buf_update_bytes_read(buf, 1)

		// not a list
		if list != 0 {
			list--
		}
	}

	sml_buf_update_bytes_read(buf, 1)

	return int(length) + list
}

func sml_buf_optional_is_skipped(buf *sml_buffer) bool {
	if sml_buf_get_current_byte(buf) == SML_OPTIONAL_SKIPPED {
		sml_buf_update_bytes_read(buf, 1)
		return true
	}

	return false
}
