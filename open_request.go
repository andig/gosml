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

func OpenRequestParse(buf Buffer) (OpenRequest, error) {
	msg := &OpenRequest{}
	var err error

	if err := Expect(buf, TYPELIST, 7); err != nil {
		return msg, err
	}

	if msg.Codepage = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.ClientID = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.ReqFileID = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.ServerID = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.Username = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.Password = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.Version = U8Parse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
