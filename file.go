package sml

func sml_file_parse(bytes []byte) ([]sml_message, error) {
	buf := &sml_buffer{}
	buf.buf = make([]byte, SML_MAX_FILE_SIZE)
	copy(buf.buf, bytes)

	messages := make([]sml_message, 0)

	for buf.cursor < len(buf.buf) {
		if sml_buf_get_current_byte(buf) == SML_MESSAGE_END {
			// reading trailing zeroed bytes
			sml_buf_update_bytes_read(buf, 1)
			continue;
		}

		msg, err := sml_message_parse(buf)
		if err != nil {
			return messages, err
		}

		messages = append(messages, msg)
	}

	return messages, nil
}
