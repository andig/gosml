package sml

type GetListResponse struct {
	ClientId       OctetString
	ServerId       OctetString
	ListName       OctetString
	ActSensorTime  Time
	ValList        []ListEntry
	ListSignature  OctetString
	ActGatewayTime Time
}

func GetListResponseParse(buf *Buffer) (GetListResponse, error) {
	list := GetListResponse{}
	var err error

	if err := Expect(buf, TYPELIST, 7); err != nil {
		return list, err
	}

	if list.ClientId, err = OctetStringParse(buf); err != nil {
		return list, err
	}

	if list.ServerId, err = OctetStringParse(buf); err != nil {
		return list, err
	}

	if list.ListName, err = OctetStringParse(buf); err != nil {
		return list, err
	}

	if list.ActSensorTime, err = TimeParse(buf); err != nil {
		return list, err
	}

	if list.ValList, err = ListParse(buf); err != nil {
		return list, err
	}

	if list.ListSignature, err = OctetStringParse(buf); err != nil {
		return list, err
	}

	if list.ActGatewayTime, err = TimeParse(buf); err != nil {
		return list, err
	}

	return list, nil
}
