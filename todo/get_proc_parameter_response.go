package sml


type struct {
	octetString *serverId;
	TreePath *parameterTreePath;
	Tree *parameterTree;
} GetProcParameterResponse;


GetProcParameterResponse *GetProcParameterResponseParse(Buffer *buf) {
	GetProcParameterResponse *msg = GetProcParameterResponseInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 3) {
		buf.error = 1;
		goto error;
	}

	msg.serverId = OctetStringParse(buf);
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
	GetProcParameterResponseFree(msg);
	return 0;
}
