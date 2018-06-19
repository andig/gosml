package sml

type GetListRequest struct {
	ClientID OctetString
	ServerID OctetString // optional
	Username OctetString // optional
	Password OctetString // optional
	ListName OctetString // optional
}

func GetListRequestParse(buf Buffer) (GetListRequest, error) {
	msg := &GetListRequest{}
	var err error

	if err := Expect(buf, TYPELIST, 5); err != nil {
		return msg, err
	}

	if msg.ClientID, err = OctetStringParse(buf); err != nil {
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

	if msg.ListName, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	return msg, nil
}
