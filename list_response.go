package sml

type sml_get_list_response struct {
	client_id sml_octet_string
	server_id sml_octet_string
	list_name sml_octet_string
	act_sensor_time sml_time
	val_list []sml_list_entry
	list_signature sml_octet_string
	act_gateway_time sml_time
}

func sml_get_list_response_parse(buf *sml_buffer) (sml_get_list_response, error) {
	list := sml_get_list_response{}
	var err error

	if err := sml_expect(buf, SML_TYPE_LIST, 7); err != nil {
		return list, err
	}

	if list.client_id, err = sml_octet_string_parse(buf); err != nil {
		return list, err
	}

	if list.server_id, err = sml_octet_string_parse(buf); err != nil {
		return list, err
	}

	if list.list_name, err = sml_octet_string_parse(buf); err != nil {
		return list, err
	}

	if list.act_sensor_time, err = sml_time_parse(buf); err != nil {
		return list, err
	}

	if list.val_list, err = sml_list_parse(buf); err != nil {
		return list, err
	}

	if list.list_signature, err = sml_octet_string_parse(buf); err != nil {
		return list, err
	}

	if list.act_gateway_time, err = sml_time_parse(buf); err != nil {
		return list, err
	}

	return list, nil
}
