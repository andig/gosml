package sml

type struct ObjReqEntryListEntry {
	ObjReqEntry *objectListEntry;

	// list specific
	struct ObjReqEntryListEntry *next;
} ObjReqEntryList;

type struct {
	octetString *serverId;	// optional
	octetString *username; 	//  optional
	octetString *password; 	//  optional
	Boolean *withRawdata;  // optional
	Time *beginTime;		// optional
	Time *endTime;			// optional
	TreePath *parameterTreePath;
	ObjReqEntryList *objectList; // optional
	Tree *dasDetails;		// optional
} GetProfilePackRequest;


GetProfilePackRequest *GetProfilePackRequestParse(Buffer *buf) {
	GetProfilePackRequest *msg = GetProfilePackRequestInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 9) {
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

	msg.withRawdata = BooleanParse(buf);
	if err != nil {
		return err
	}

	msg.beginTime = TimeParse(buf);
	if err != nil {
		return err
	}

	msg.endTime = TimeParse(buf);
	if err != nil {
		return err
	}

	msg.parameterTreePath = TreePathParse(buf);
	if err != nil {
		return err
	}

	if (!BufOptionalIsSkipped(buf)) {
		if (BufGetNextType(buf) != TYPELIST) {
			buf.error = 1;
			goto error;
		}
		int i, len = BufGetNextLength(buf);
		ObjReqEntryList *last = 0, *n = 0;
		for (i = len; i > 0; i--) {
			n = (ObjReqEntryList *) malloc(sizeof(ObjReqEntryList));
			*n = ( ObjReqEntryList ) {
				.objectListEntry = NULL,
				.next = NULL
			};
			n.objectListEntry = ObjReqEntryParse(buf);
			if err != nil {
				return err
			}

			if (msg.objectList == 0) {
				msg.objectList = n;
				last = msg.objectList;
			}
			else {
				last.next = n;
				last = n;
			}
		}
	}

	msg.dasDetails = TreeParse(buf);
	if err != nil {
		return err
	}

	return msg;

error:
	GetProfilePackRequestFree(msg);
	return 0;
}
