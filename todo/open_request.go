package sml

type struct {
    octetString *codepage; // optional
	octetString *clientId;
	octetString *reqFileId;
	octetString *serverId; // optional
	octetString *username; // optional
	octetString *password; // optional
	u8 *Version; // optional
} OpenRequest;


OpenRequest *OpenRequestParse(Buffer *buf) {
	OpenRequest *msg = OpenRequestInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 7) {
		buf.error = 1;
		goto error;
	}

	msg.codepage = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.clientId = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.reqFileId = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.serverId = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.username = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.password = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.Version = U8Parse(buf);
	if err != nil {
		return err
	}

	return msg;

error:
	OpenRequestFree(msg);
	return 0;
}
