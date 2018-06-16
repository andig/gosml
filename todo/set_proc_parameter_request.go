package sml

type struct {
	octetString *serverId; // optional
	octetString *username; // optional
	octetString *password; // optional
	TreePath *parameterTreePath;
	Tree *parameterTree;
} SetProcParameterRequest;

SetProcParameterRequest *SetProcParameterRequestParse(Buffer *buf) {
	SetProcParameterRequest *msg = SetProcParameterRequestInit();

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

	msg.parameterTree = TreeParse(buf);
	if err != nil {
		return err
	}

	return msg;

error:
	SetProcParameterRequestFree(msg);
	return 0;
}
