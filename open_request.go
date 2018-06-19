package sml

type OpenRequest struct {
	Codepage  OctetString // optional
	ClientID  OctetString
	ReqFileID OctetString
	ServerID  OctetString // optional
	Username  OctetString // optional
	Password  OctetString // optional
	Version   uint8       // optional
}

func OpenRequestParse(buf *Buffer) (OpenRequest, error) {
	msg := OpenRequest{}
	var err error

	if err := Expect(buf, TYPELIST, 7); err != nil {
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

	if msg.Username, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.Password, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.Version, err = U8Parse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
