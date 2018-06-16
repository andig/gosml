package sml

type AttentionResponse struct
{
	octetString *serverId;
	octetString *attentionNumber;
	octetString *attentionMessage; // optional
	Tree *attentionDetails;	 // optional
}


AttentionResponse *AttentionResponseParse(Buffer *buf){
	AttentionResponse *msg = AttentionResponseInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 4) {
		buf.error = 1;
		goto error;
	}

	msg.serverId = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.attentionNumber = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.attentionMessage = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.attentionDetails = TreeParse(buf);
	if err != nil {
		return err
	}

	return msg;

	error:
		AttentionResponseFree(msg);
		return 0;
}
