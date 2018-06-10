package sml

import (
	"fmt"
)

type sml_octet_string []byte

func sml_octet_string_parse(buf *sml_buffer) (sml_octet_string, error) {
	if skip := sml_buf_optional_is_skipped(buf); skip {
		return nil, nil;
	}

	sml_debug(buf, "sml_octet_str_parse")

	if err := sml_expect_type(buf, SML_TYPE_OCTET_STRING); err != nil {
		return nil, err
	}

	length := sml_buf_get_next_length(buf)
	if length < 0 {
		return nil, fmt.Errorf("sml: Invalid octet string length %d", length)
	}

	str := buf.buf[buf.cursor:buf.cursor+length]
	sml_buf_update_bytes_read(buf, length)

	return str, nil
}

