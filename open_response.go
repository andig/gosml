package sml

type OpenResponse struct {
	Codepage  OctetString
	ClientID  OctetString
	ReqFileID OctetString
	ServerID  OctetString
	RefTime   Time
	Version   uint8
}

func OpenResponseParse(buf *Buffer) (OpenResponse, error) {
	msg := OpenResponse{}
	var err error

	if err := Expect(buf, TYPELIST, 6); err != nil {
		return msg, err
	}

	if msg.Codepage, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.ClientID, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.ReqFileID, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.ServerID, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.RefTime, err = TimeParse(buf); err != nil {
		return msg, err
	}

	if msg.Version, err = U8Parse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
