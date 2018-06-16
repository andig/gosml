package sml


type struct {
	octetString *serverId;
	Time *actTime; // specified by germans (current time was meant) ;)
	u32 *regPeriod;
	TreePath *parameterTreePath;
	Sequence *headerList; 			// list of ProfObjHeaderEntry
	Sequence *periodList;			// list of ProfObjPeriodEntry
	octetString *rawdata;  			// optional
	Signature *profileSignature; 	// optional

} GetProfilePackResponse;

type struct {
	octetString *objName;
	Unit *unit;
	i8 *scaler;
} ProfObjHeaderEntry;

type struct {
	Time *valTime;
	u64 *status;
	Sequence *valueList;
	Signature *periodSignature;
} ProfObjPeriodEntry;

type struct {
	Value *value;
	Signature *valueSignature;
} ValueEntry;

static void * ProfObjHeaderEntryParse(Buffer *buf);
static void * ValueEntryParse(Buffer *buf);
static void ValueEntryFree( void * p );
static void ValueEntryWrite( void * p, Buffer *buf);

static void ProfObjHeaderEntryFree( void * p ) {
	ProfObjHeaderEntry * entry = p;
	if (entry) {
		OctetStringFree(entry.objName);
		UnitFree(entry.unit);
		NumberFree(entry.scaler);

		free(entry);
	}
}

static void * ProfObjPeriodEntryParse(Buffer *buf) {
	ProfObjPeriodEntry *entry = ProfObjPeriodEntryInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 4) {
		buf.error = 1;
		goto error;
	}


	entry.valTime = TimeParse(buf);
	if (BufHasErrors(buf)) goto error;
	entry.status = U64Parse(buf);
	if (BufHasErrors(buf)) goto error;
	entry.valueList = SequenceParse(buf, ValueEntryParse, ValueEntryFree);
	if (BufHasErrors(buf)) goto error;
	entry.periodSignature = SignatureParse(buf);
	if (BufHasErrors(buf)) goto error;

	return entry;

error:
	buf.error = 1;
	ProfObjPeriodEntryFree(entry);
	return 0;
}

static void ProfObjPeriodEntryFree( void * p ) {
	ProfObjPeriodEntry * entry = p;

	if (entry) {
		TimeFree(entry.valTime);
		NumberFree(entry.status);
		SequenceFree(entry.valueList);
		SignatureFree(entry.periodSignature);

		free(entry);
	}
}

static void ProfObjHeaderEntryWrite( void * p, Buffer *buf) {
	ProfObjHeaderEntry * entry = p;

	BufSetTypeAndLength(buf, TYPELIST, 3);

	OctetStringWrite(entry.objName, buf);
	UnitWrite(entry.unit, buf);
	I8Write(entry.scaler, buf);
}

static void ProfObjPeriodEntryWrite( void * p, Buffer *buf) {
	ProfObjPeriodEntry * entry = p;

	BufSetTypeAndLength(buf, TYPELIST, 4);
	TimeWrite(entry.valTime, buf);
	U64Write(entry.status, buf);
	SequenceWrite(entry.valueList, buf, ValueEntryWrite);
	SignatureWrite(entry.periodSignature, buf);
}

static void * ValueEntryParse(Buffer *buf) {
	ValueEntry *entry = ValueEntryInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 2) {
		buf.error = 1;
		goto error;
	}

	entry.value = ValueParse(buf);
	if (BufHasErrors(buf)) goto error;
	entry.valueSignature = SignatureParse(buf);
	if (BufHasErrors(buf)) goto error;

	return entry;

error:
	buf.error = 1;
	ValueEntryFree(entry);
	return 0;
}

static void ValueEntryWrite( void * p, Buffer *buf) {
	ValueEntry * entry = p;

	BufSetTypeAndLength(buf, TYPELIST, 2);
	ValueWrite(entry.value, buf);
	SignatureWrite(entry.valueSignature, buf);
}

static void ValueEntryFree( void * p ) {
	ValueEntry * entry = p;

	if (entry) {
		ValueFree(entry.value);
		SignatureFree(entry.valueSignature);

		free(entry);
	}
}

// GetProfilePackResponse;

GetProfilePackResponse *GetProfilePackResponseInit() {
	GetProfilePackResponse *msg = (GetProfilePackResponse *) malloc(sizeof(GetProfilePackResponse));
	*msg = ( GetProfilePackResponse ) {
		.serverId = NULL,
		.actTime = NULL,
		.regPeriod = NULL,
		.parameterTreePath = NULL,
		.headerList = NULL,
		.periodList = NULL,
		.rawdata = NULL,
		.profileSignature = NULL
	};

	return msg;
}

GetProfilePackResponse *GetProfilePackResponseParse(Buffer *buf){
	GetProfilePackResponse *msg = GetProfilePackResponseInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 8) {
		buf.error = 1;
		goto error;
	}

	msg.serverId = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;

	msg.actTime = TimeParse(buf);
	if (BufHasErrors(buf)) goto error;

	msg.regPeriod = U32Parse(buf);
	if (BufHasErrors(buf)) goto error;

	msg.parameterTreePath = TreePathParse(buf);
	if (BufHasErrors(buf)) goto error;

	msg.headerList = SequenceParse(buf, ProfObjHeaderEntryParse, ProfObjHeaderEntryFree);
	if (BufHasErrors(buf)) goto error;

	msg.periodList = SequenceParse(buf, ProfObjPeriodEntryParse, ProfObjPeriodEntryFree);
	if (BufHasErrors(buf)) goto error;

	msg.rawdata = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;

	msg.profileSignature = SignatureParse(buf);
	if (BufHasErrors(buf)) goto error;

	return msg;

error:
	buf.error = 1;
	GetProfilePackResponseFree(msg);
	return 0;
}

void GetProfilePackResponseWrite(GetProfilePackResponse *msg, Buffer *buf) {
	BufSetTypeAndLength(buf, TYPELIST, 8);

	OctetStringWrite(msg.serverId, buf);
	TimeWrite(msg.actTime, buf);
	U32Write(msg.regPeriod, buf);
	TreePathWrite(msg.parameterTreePath, buf);
	SequenceWrite(msg.headerList, buf, ProfObjHeaderEntryWrite);
	SequenceWrite(msg.periodList, buf, ProfObjPeriodEntryWrite);
	OctetStringWrite(msg.rawdata, buf);
	SignatureWrite(msg.profileSignature, buf);
}

void GetProfilePackResponseFree(GetProfilePackResponse *msg){
	if (msg) {
		OctetStringFree(msg.serverId);
		TimeFree(msg.actTime);
		NumberFree(msg.regPeriod);
		TreePathFree(msg.parameterTreePath);
		SequenceFree(msg.headerList);
		SequenceFree(msg.periodList);
		OctetStringFree(msg.rawdata);
		SignatureFree(msg.profileSignature);

		free(msg);
	}
}


// ProfObjHeaderEntry;

ProfObjHeaderEntry *ProfObjHeaderEntryInit() {
	ProfObjHeaderEntry *entry = (ProfObjHeaderEntry *) malloc(sizeof(ProfObjHeaderEntry));
	*entry = ( ProfObjHeaderEntry ) {
		.objName = NULL,
		.unit = NULL,
		.scaler = NULL
	};
	return entry;
}

static void * ProfObjHeaderEntryParse(Buffer *buf) {
	ProfObjHeaderEntry *entry = ProfObjHeaderEntryInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 3) {
		buf.error = 1;
		goto error;
	}

	entry.objName = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;
	entry.unit = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;
	entry.scaler = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;

	return entry;
error:
	buf.error = 1;
	ProfObjHeaderEntryFree(entry);
	return 0;
}

ProfObjHeaderEntry * ProfObjHeaderEntryParse( Buffer * buf ) {
	return ProfObjHeaderEntryParse( buf );
}

void ProfObjHeaderEntryWrite( ProfObjHeaderEntry * entry, Buffer * buf ) {
	ProfObjHeaderEntryWrite( entry, buf );
}

void ProfObjHeaderEntryFree( ProfObjHeaderEntry * entry ) {
	ProfObjHeaderEntryFree( entry );
}


// ProfObjPeriodEntry;

ProfObjPeriodEntry *ProfObjPeriodEntryInit() {
	ProfObjPeriodEntry *entry = (ProfObjPeriodEntry *) malloc(sizeof(ProfObjPeriodEntry));
	*entry = ( ProfObjPeriodEntry ) {
		.valTime = NULL,
		.status = NULL,
		.valueList = NULL,
		.periodSignature = NULL
	};
	return entry;
}

ProfObjPeriodEntry * ProfObjPeriodEntryParse( Buffer * buf ) {
	return ProfObjPeriodEntryParse( buf );
}

// ValueEntry;

ValueEntry * ValueEntryParse( Buffer * buf ) {
	return ValueEntryParse( buf );
}
