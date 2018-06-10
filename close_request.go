package sml

type sml_close_request struct {
	global_signature sml_octet_string
}

func sml_close_request_parse(buf *sml_buffer) (sml_close_request, error) {
	msg := sml_close_request{}
	var err error

	if err := sml_expect(buf, SML_TYPE_LIST, 1); err != nil {
		return msg, err
	}

	if msg.global_signature, err = sml_octet_string_parse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
