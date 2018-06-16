package sml

func FileParse(bytes []byte) ([]Message, error) {
	buf := &Buffer{}
	buf.Bytes = make([]byte, MAXFILESIZE)
	copy(buf.Bytes, bytes)

	messages := make([]Message, 0)

	for buf.Cursor < len(buf.Bytes) {
		if BufGetCurrentByte(buf) == MESSAGEEND {
			// reading trailing zeroed bytes
			BufUpdateBytesRead(buf, 1)
			continue
		}

		msg, err := MessageParse(buf)
		if err != nil {
			return messages, err
		}

		messages = append(messages, msg)
	}

	return messages, nil
}
