package sml

type CloseResponse CloseRequest

func CloseResponseParse(Buffer *buf) (CloseResponse, error) {
	msg = &CloseResponse{}

	if err := Expect(buf, TYPELIST, 1); err != nil {
		return msg, err
	}

	if msg.GlobalSignature, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
