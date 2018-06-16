package sml

type struct {
	octetString *serverId;
	Time *actTime;
	u32 *regPeriod;
	TreePath *parameterTreePath;
	Time *valTime;
	u64 *status;
	Sequence *periodList;
	octetString *rawdata;
	Signature *periodSignature;
} GetProfileListResponse;


static void * PeriodEntryParse2( Buffer * buf ) {
	return PeriodEntryParse( buf );
}

static void PeriodEntryFree2( void * p ) {
	PeriodEntryFree( p );
}

GetProfileListResponse *GetProfileListResponseParse(Buffer *buf) {
	GetProfileListResponse *msg = GetProfileListResponseInit();

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

	msg.actTime = TimeParse(buf);
	if err != nil {
		return err
	}

	msg.regPeriod = U32Parse(buf);
	if err != nil {
		return err
	}

	msg.parameterTreePath = TreePathParse(buf);
	if err != nil {
		return err
	}

	msg.valTime = TimeParse(buf);
	if err != nil {
		return err
	}

	msg.status = U64Parse(buf);
	if err != nil {
		return err
	}

	msg.periodList = SequenceParse(buf, PeriodEntryParse2, PeriodEntryFree2);
	if err != nil {
		return err
	}

	msg.rawdata = OctetStringParse(buf);
	if err != nil {
		return err
	}

	msg.periodSignature = SignatureParse(buf);
	if err != nil {
		return err
	}

	return msg;

error:
	buf.error = 1;
	GetProfileListResponseFree(msg);
	return 0;
}
