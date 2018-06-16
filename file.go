package sml

func FileParse(bytes []byte) ([]Message, error) {
	buf := &Buffer{}
	buf.Buf = make([]byte, MAXFILESIZE)
	copy(buf.Buf, bytes)

	messages := make([]Message, 0)

	for buf.Cursor < len(buf.Buf) {
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
