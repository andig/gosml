package sml

func sml_boolean_parse(buf *sml_buffer) (bool, error) {
	if sml_buf_optional_is_skipped(buf) {
		return false, nil
	}

	sml_debug(buf, "sml_boolean_parse")

	if err := sml_expect(buf, SML_TYPE_BOOLEAN, 1); err != nil {
		return false, err
	}

	if sml_buf_get_current_byte(buf) > 0 {
		sml_buf_update_bytes_read(buf, 1)
		return true, nil
	} else {
		sml_buf_update_bytes_read(buf, 1)
		return false, nil
	}
}
