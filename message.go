package sml

import (
	//	"fmt"

	"errors"
	"fmt"
)

const (
	MESSAGEOPENREQUEST              = 0x00000100
	MESSAGEOPENRESPONSE             = 0x00000101
	MESSAGECLOSEREQUEST             = 0x00000200
	MESSAGECLOSERESPONSE            = 0x00000201
	MESSAGEGETPROFILEPACKREQUEST    = 0x00000300
	MESSAGEGETPROFILEPACKRESPONSE   = 0x00000301
	MESSAGEGETPROFILELISTREQUEST    = 0x00000400
	MESSAGEGETPROFILELISTRESPONSE   = 0x00000401
	MESSAGEGETPROCPARAMETERREQUEST  = 0x00000500
	MESSAGEGETPROCPARAMETERRESPONSE = 0x00000501
	MESSAGESETPROCPARAMETERREQUEST  = 0x00000600
	MESSAGESETPROCPARAMETERRESPONSE = 0x00000601 // This doesn't exist in the spec
	MESSAGEGETLISTREQUEST           = 0x00000700
	MESSAGEGETLISTRESPONSE          = 0x00000701
	MESSAGEATTENTIONRESPONSE        = 0x0000FF01
)

type Message struct {
	TransactionID OctetString
	GroupID       uint8
	AbortOnError  uint8
	MessageBody   MessageBody
	Crc           uint16
}

type MessageBody struct {
	Tag  uint32
	Data MessageBodyData
}

type MessageBodyData interface{}

func MessageBodyParse(buf *Buffer) (MessageBody, error) {
	body := MessageBody{}
	var err error

	if err := Expect(buf, TYPELIST, 2); err != nil {
		return body, err
	}

	if body.Tag, err = U32Parse(buf); err != nil {
		return body, err
	}

	switch body.Tag {
	case MESSAGEOPENREQUEST:
		body.Data, err = OpenRequestParse(buf)
		return body, err
	case MESSAGEOPENRESPONSE:
		body.Data, err = OpenResponseParse(buf)
		return body, err
	case MESSAGECLOSEREQUEST:
		body.Data, err = CloseRequestParse(buf)
		return body, err
	case MESSAGECLOSERESPONSE:
		body.Data, err = CloseResponseParse(buf)
		return body, err
	case MESSAGEGETPROFILEPACKREQUEST:
		return body, fmt.Errorf("Unimplemented message type MESSAGEGETPROFILEPACKREQUEST")
		// msgBody->data = GetProfilePackRequestParse(buf);
	case MESSAGEGETPROFILEPACKRESPONSE:
		return body, fmt.Errorf("Unimplemented message type MESSAGEGETPROFILEPACKRESPONSE")
		// msgBody->data = GetProfilePackResponseParse(buf);
	case MESSAGEGETPROFILELISTREQUEST:
		return body, fmt.Errorf("Unimplemented message type MESSAGEGETPROFILELISTREQUEST")
		// msgBody->data = GetProfileListRequestParse(buf);
	case MESSAGEGETPROFILELISTRESPONSE:
		return body, fmt.Errorf("Unimplemented message type MESSAGEGETPROFILELISTRESPONSE")
		// msgBody->data = GetProfileListResponseParse(buf);
	case MESSAGEGETPROCPARAMETERREQUEST:
		return body, fmt.Errorf("Unimplemented message type MESSAGEGETPROCPARAMETERREQUEST")
		// msgBody->data = GetProcParameterRequestParse(buf);
	case MESSAGEGETPROCPARAMETERRESPONSE:
		return body, fmt.Errorf("Unimplemented message type MESSAGEGETPROCPARAMETERRESPONSE")
		// msgBody->data = GetProcParameterResponseParse(buf);
	case MESSAGESETPROCPARAMETERREQUEST:
		return body, fmt.Errorf("Unimplemented message type MESSAGESETPROCPARAMETERREQUEST")
		// msgBody->data = SetProcParameterRequestParse(buf);
	case MESSAGEGETLISTREQUEST:
		body.Data, err = GetListRequestParse(buf)
		return body, err
	case MESSAGEGETLISTRESPONSE:
		body.Data, err = GetListResponseParse(buf)
		return body, err
	case MESSAGEATTENTIONRESPONSE:
		return body, fmt.Errorf("Unimplemented message type MESSAGEATTENTIONRESPONSE")
		// msgBody->data = AttentionResponseParse(buf);
	}

	return body, fmt.Errorf("Invalid message type: % x", body.Tag)
}

func MessageParse(buf *Buffer, validate ...bool) (Message, error) {
	Debug(buf, "MessageParse")

	msg := Message{}
	var err error

	crcStart := buf.Cursor

	if err := Expect(buf, TYPELIST, 6); err != nil {
		return msg, err
	}

	if msg.TransactionID, err = OctetStringParse(buf); err != nil {
		return msg, err
	}

	if msg.GroupID, err = U8Parse(buf); err != nil {
		return msg, err
	}

	if msg.AbortOnError, err = U8Parse(buf); err != nil {
		return msg, err
	}

	if msg.MessageBody, err = MessageBodyParse(buf); err != nil {
		return msg, err
	}

	crcEnd := buf.Cursor

	if msg.Crc, err = U16Parse(buf); err != nil {
		return msg, err
	}

	if len(validate) > 0 && validate[0] {
		//		fmt.Println(buf.Cursor)
		crc := Crc16Calculate(buf.Bytes[crcStart:crcEnd], crcEnd-crcStart)
		//		fmt.Printf("%04x-%04x\n", crc, msg.Crc)

		if crc != msg.Crc {
			err := errors.New("Crc error")
			return msg, err
		}
	}

	if BufGetCurrentByte(buf) == MESSAGEEND {
		BufUpdateBytesRead(buf, 1)
	}

	return msg, nil
}
