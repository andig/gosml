package sml

import(
)

type sml_open_response struct {
	codepage sml_octet_string
	client_id sml_octet_string
	req_file_id sml_octet_string
	server_id sml_octet_string
	ref_time sml_time
	sml_version uint8
}

func sml_open_response_parse(buf *sml_buffer) (sml_open_response, error) {
	sml_debug(buf, "OPEN_REPONSE")

	var err error
	msg := sml_open_response{}

	if err := sml_expect(buf, SML_TYPE_LIST, 6); err != nil {
		return msg, err
	}

	if msg.codepage, err = sml_octet_string_parse(buf); err != nil {
		return msg, err
	}

	if msg.client_id, err = sml_octet_string_parse(buf); err != nil {
		return msg, err
	}

	if msg.req_file_id, err = sml_octet_string_parse(buf); err != nil {
		return msg, err
	}

	if msg.server_id, err = sml_octet_string_parse(buf); err != nil {
		return msg, err
	}

	if msg.ref_time, err = sml_time_parse(buf); err != nil {
		return msg, err
	}

	if msg.sml_version, err = sml_u8_parse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
