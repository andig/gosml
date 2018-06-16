package sml

type struct {
    octetString *serverId; // optional
    octetString *username; // optional
    octetString *password; // optional
    TreePath *parameterTreePath;
    octetString *attribute; // optional
} GetProcParameterRequest;


GetProcParameterRequest *GetProcParameterRequestParse(Buffer *buf) {
	GetProcParameterRequest *msg = GetProcParameterRequestInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 5) {
		buf.error = 1;
		goto error;
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

	msg.parameterTreePath = TreePathParse(buf);
	if err != nil {
		return err
	}

	msg.attribute = OctetStringParse(buf);
	if err != nil {
		return err
	}

	return msg;

error:
	GetProcParameterRequestFree(msg);
	return 0;
}
