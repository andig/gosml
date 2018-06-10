package sml

import (
	"fmt"
)

type sml_value struct {
	typ uint8
	data_bytes sml_octet_string
	data_boolean bool
	data_int int64
}

func sml_value_parse(buf *sml_buffer) (sml_value, error) {
/*
	if (sml_buf_optional_is_skipped(buf)) {
		return 0;
	}

	int max = 1;
	int type = sml_buf_get_next_type(buf);
	unsigned char byte = sml_buf_get_current_byte(buf);

	sml_value *value = sml_value_init();
	value->type = type;

	switch (type) {
		case SML_TYPE_OCTET_STRING:
			value->data.bytes = sml_octet_string_parse(buf);
			break;
		case SML_TYPE_BOOLEAN:
			value->data.boolean = sml_boolean_parse(buf);
			break;
		case SML_TYPE_UNSIGNED:
		case SML_TYPE_INTEGER:
			// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
			while (max < ((byte & SML_LENGTH_FIELD) - 1)) {
				max <<= 1;
			}

			value->data.uint8 = sml_number_parse(buf, type, max);
			value->type |= max;
			break;
		default:
			buf->error = 1;
			break;
	}
*/
	value := sml_value{}

	if sml_buf_optional_is_skipped(buf) {
		return value, nil
	}

	sml_debug(buf, "sml_value_parse")

	typefield := sml_buf_get_next_type(buf)
	b := sml_buf_get_current_byte(buf)

	max := 1
	value.typ = typefield

	var err error
	switch (typefield) {
		case SML_TYPE_OCTET_STRING:
			value.data_bytes, err = sml_octet_string_parse(buf)
			if err != nil {
				return value, err
			}
		case SML_TYPE_BOOLEAN:
			value.data_boolean, err = sml_boolean_parse(buf)
			if err != nil {
				return value, err
			}
		case SML_TYPE_UNSIGNED:
			// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
			for max < int((b & SML_LENGTH_FIELD) - 1) {
				max = max << 1
			}

			value.data_int, err = sml_number_parse(buf, typefield, max)
			if err != nil {
				return value, err
			}

			value.typ = value.typ | uint8(max)
		case SML_TYPE_INTEGER:
			// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
			for max < int((b & SML_LENGTH_FIELD) - 1) {
				max = max << 1
			}

			value.data_int, err = sml_number_parse(buf, typefield, max)
			if err != nil {
				return value, err
			}

			value.typ = value.typ | uint8(max)
		default:
			return value, fmt.Errorf("sml: Unexpected type %02x", typefield)
	}

	return value, nil
}
