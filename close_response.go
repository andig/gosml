package sml

type CloseResponse CloseRequest

func CloseResponseParse(buf *Buffer) (CloseResponse, error) {
	msg := CloseResponse{}
	var err error

	if err := Expect(buf, TYPELIST, 1); err != nil {
		return msg, err
	}

	if msg.GlobalSignature, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
