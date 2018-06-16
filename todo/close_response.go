package sml

type CloseRequest CloseResponse;

CloseResponse *CloseResponseParse(Buffer *buf) {
	CloseResponse *msg = CloseResponseInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 1) {
		buf.error = 1;
		goto error;
	}

	msg.globalSignature = OctetStringParse(buf);
	if err != nil {
		return err
	}

	return msg;

error:
	CloseResponseFree(msg);
	return 0;
}
