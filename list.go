package sml

type sml_list_entry struct {
	obj_name sml_octet_string
	status int64
	val_time sml_time
	unit uint8
	scaler int8
	value sml_value
	value_signature sml_octet_string
}

func sml_list_entry_parse(buf *sml_buffer) (sml_list_entry, error) {
	sml_debug(buf, "sml_list_entry_parse")

	elem := sml_list_entry{}
	var err error

	if err := sml_expect(buf, SML_TYPE_LIST, 7); err != nil {
		return elem, err
	}

	if elem.obj_name, err = sml_octet_string_parse(buf); err != nil {
		return elem, err
	}

	if elem.status, err = sml_status_parse(buf); err != nil {
		return elem, err
	}

	if elem.val_time, err = sml_time_parse(buf); err != nil {
		return elem, err
	}

	if elem.unit, err = sml_u8_parse(buf); err != nil {
		return elem, err
	}

	if elem.scaler, err = sml_i8_parse(buf); err != nil {
		return elem, err
	}

	if elem.value, err = sml_value_parse(buf); err != nil {
		return elem, err
	}

	if elem.value_signature, err = sml_octet_string_parse(buf); err != nil {
		return elem, err
	}

	return elem, nil
}

func sml_list_parse(buf *sml_buffer) ([]sml_list_entry, error) {
	if sml_buf_optional_is_skipped(buf) {
		return nil, nil
	}

	sml_debug(buf, "sml_list_parse")

	if err := sml_expect_type(buf, SML_TYPE_LIST); err != nil {
		return nil, err
	}

	list := make([]sml_list_entry, 0)

	elems := sml_buf_get_next_length(buf)

	for elems > 0 {
		elem, err := sml_list_entry_parse(buf)
		if err != nil {
			return nil, err
		}
		list = append(list, elem)
		elems--
	}

	return list, nil
}
