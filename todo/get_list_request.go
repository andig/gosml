package sml

type struct {
	octetString *clientId;
	octetString *serverId;    // optional
	octetString *username;		// optional
	octetString *password;		// optional
	octetString *listName; 	// optional
} GetListRequest;


GetListRequest *GetListRequestParse(Buffer *buf) {
	GetListRequest *msg = (GetListRequest *) malloc(sizeof(GetListRequest));
	*msg = ( GetListRequest ) {
		.clientId = NULL,
		.serverId = NULL,
		.username = NULL,
		.password = NULL,
		.listName = NULL
	};

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 5) {
		buf.error = 1;
		goto error;
	}

	msg.clientId = OctetStringParse(buf);
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

	msg.listName = OctetStringParse(buf);
	if err != nil {
		return err
	}

	return msg;

error:
	GetListRequestFree(msg);
	return 0;
}


void GetListRequestFree(GetListRequest *msg) {
	if (msg) {
		OctetStringFree(msg.clientId);
		OctetStringFree(msg.serverId);
		OctetStringFree(msg.listName);
		OctetStringFree(msg.username);
		OctetStringFree(msg.password);
		free(msg);
	}
}

