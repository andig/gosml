package sml

import (
	"fmt"
)

const (
	SML_MESSAGE_OPEN_REQUEST                = 0x00000100
	SML_MESSAGE_OPEN_RESPONSE               = 0x00000101
	SML_MESSAGE_CLOSE_REQUEST               = 0x00000200
	SML_MESSAGE_CLOSE_RESPONSE              = 0x00000201
	SML_MESSAGE_GET_PROFILE_PACK_REQUEST    = 0x00000300
	SML_MESSAGE_GET_PROFILE_PACK_RESPONSE   = 0x00000301
	SML_MESSAGE_GET_PROFILE_LIST_REQUEST    = 0x00000400
	SML_MESSAGE_GET_PROFILE_LIST_RESPONSE   = 0x00000401
	SML_MESSAGE_GET_PROC_PARAMETER_REQUEST  = 0x00000500
	SML_MESSAGE_GET_PROC_PARAMETER_RESPONSE = 0x00000501
	SML_MESSAGE_SET_PROC_PARAMETER_REQUEST  = 0x00000600
	SML_MESSAGE_SET_PROC_PARAMETER_RESPONSE = 0x00000601 // This doesn't exist in the spec
	SML_MESSAGE_GET_LIST_REQUEST            = 0x00000700
	SML_MESSAGE_GET_LIST_RESPONSE           = 0x00000701
	SML_MESSAGE_ATTENTION_RESPONSE          = 0x0000FF01
)

type sml_message struct {
	transaction_id sml_octet_string
	group_id uint8
	abort_on_error uint8
	message_body sml_message_body
	crc uint16
}

type sml_message_body struct {
	tag uint32
	data sml_message_body_data
}

type sml_message_body_data interface{ }

func sml_message_body_parse(buf *sml_buffer) (sml_message_body, error) {
	body := sml_message_body{}
	var err error

	if err := sml_expect(buf, SML_TYPE_LIST, 2); err != nil {
		return body, err
	}

	if body.tag, err = sml_u32_parse(buf); err != nil {
		return body, err
	}

	switch body.tag {
	case SML_MESSAGE_OPEN_REQUEST:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_OPEN_REQUEST")
		// msg_body->data = sml_open_request_parse(buf);
	case SML_MESSAGE_OPEN_RESPONSE:
		body.data, err = sml_open_response_parse(buf)
		return body, err
	case SML_MESSAGE_CLOSE_REQUEST:
		body.data,err = sml_close_request_parse(buf)
		return body, err
	case SML_MESSAGE_CLOSE_RESPONSE:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_CLOSE_RESPONSE")
		// msg_body->data = sml_close_response_parse(buf);
	case SML_MESSAGE_GET_PROFILE_PACK_REQUEST:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_GET_PROFILE_PACK_REQUEST")
		// msg_body->data = sml_get_profile_pack_request_parse(buf);
	case SML_MESSAGE_GET_PROFILE_PACK_RESPONSE:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_GET_PROFILE_PACK_RESPONSE")
		// msg_body->data = sml_get_profile_pack_response_parse(buf);
	case SML_MESSAGE_GET_PROFILE_LIST_REQUEST:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_GET_PROFILE_LIST_REQUEST")
		// msg_body->data = sml_get_profile_list_request_parse(buf);
	case SML_MESSAGE_GET_PROFILE_LIST_RESPONSE:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_GET_PROFILE_LIST_RESPONSE")
		// msg_body->data = sml_get_profile_list_response_parse(buf);
	case SML_MESSAGE_GET_PROC_PARAMETER_REQUEST:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_GET_PROC_PARAMETER_REQUEST")
		// msg_body->data = sml_get_proc_parameter_request_parse(buf);
	case SML_MESSAGE_GET_PROC_PARAMETER_RESPONSE:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_GET_PROC_PARAMETER_RESPONSE")
		// msg_body->data = sml_get_proc_parameter_response_parse(buf);
	case SML_MESSAGE_SET_PROC_PARAMETER_REQUEST:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_SET_PROC_PARAMETER_REQUEST")
		// msg_body->data = sml_set_proc_parameter_request_parse(buf);
	case SML_MESSAGE_GET_LIST_REQUEST:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_GET_LIST_REQUEST")
		// msg_body->data = sml_get_list_request_parse(buf);
	case SML_MESSAGE_GET_LIST_RESPONSE:
		body.data, err = sml_get_list_response_parse(buf)
		return body, err
	case SML_MESSAGE_ATTENTION_RESPONSE:
		return body, fmt.Errorf("sml: Unimplemented message type SML_MESSAGE_ATTENTION_RESPONSE")
		// msg_body->data = sml_attention_response_parse(buf);
	default:
		return body, fmt.Errorf("sml: Invalid message type: % x", body.tag)
	}

	return body, nil
}

func sml_message_parse(buf *sml_buffer) (sml_message, error) {
	sml_debug(buf, "sml_message_parse")

	msg := sml_message{}
	var err error

	if err := sml_expect(buf, SML_TYPE_LIST, 6); err != nil {
		return msg, err
	}

	if msg.transaction_id, err = sml_octet_string_parse(buf); err != nil {
		return msg, err
	}

	if msg.group_id, err = sml_u8_parse(buf); err != nil {
		return msg, err
	}

	if msg.abort_on_error, err = sml_u8_parse(buf); err != nil {
		return msg, err
	}

	if msg.message_body, err = sml_message_body_parse(buf); err != nil {
		return msg, err
	}

	if msg.crc, err = sml_u16_parse(buf); err != nil {
		return msg, err
	}

	if sml_buf_get_current_byte(buf) == SML_MESSAGE_END {
		sml_buf_update_bytes_read(buf, 1)
	}

	return msg, nil
}
