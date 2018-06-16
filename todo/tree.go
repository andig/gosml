package sml


// what a messy tupel ...
type struct {
	octetString *serverId;
	Time *secIndex;
	u64 *status;

	Unit *unitPA;
	i8 *scalerPA;
	i64 *valuePA;

	Unit *unitR1;
	i8 *scalerR1;
	i64 *valueR1;

	Unit *unitR4;
	i8 *scalerR4;
	i64 *valueR4;

	octetString *signaturePAR1R4;

	Unit *unitMA;
	i8 *scalerMA;
	i64 *valueMA;

	Unit *unitR2;
	i8 *scalerR2;
	i64 *valueR2;

	Unit *unitR3;
	i8 *scalerR3;
	i64 *valueR3;

	octetString *signatureMAR2R3;
} TupelEntry;

type struct {
	octetString *objName;
	Unit *unit;
	i8 *scaler;
	Value *value;
	octetString *valueSignature;
} PeriodEntry;

type struct {
	u8 *tag;
	union {
		Value *value;
		PeriodEntry *periodEntry;
		TupelEntry *tupelEntry;
		Time *time;
	} data;
} ProcParValue;

type struct sTree{
	octetString *parameterName;
	ProcParValue *parameterValue; // optional
	struct sTree **childList; // optional

	int childListLen;
} Tree;

type struct {
	int pathEntriesLen;
	octetString **pathEntries;
} TreePath;

// TreePath;

TreePath *TreePathInit() {
	TreePath *treePath = (TreePath *) malloc(sizeof(TreePath));
	*treePath = ( TreePath ) {
		.pathEntriesLen = 0,
		.pathEntries = NULL
	};

	return treePath;
}

TreePath *TreePathParse(Buffer *buf) {
	if (BufOptionalIsSkipped(buf)) {
		return 0;
	}

	TreePath *treePath = TreePathInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		return 0;
	}

	octetString *s;
	int elems;
	for (elems = BufGetNextLength(buf); elems > 0; elems--) {
		s = OctetStringParse(buf);
		if (BufHasErrors(buf)) goto error;
		if (s) {
			TreePathAddPathEntry(treePath, s);
		}
	}

	return treePath;

error:
	buf.error = 1;
	TreePathFree(treePath);
	return 0;
}

void TreePathAddPathEntry(TreePath *treePath, octetString *entry) {
	treePath.pathEntriesLen++;
	treePath.pathEntries = (octetString **) realloc(treePath.pathEntries,
		sizeof(octetString *) * treePath.pathEntriesLen);

	treePath.pathEntries[treePath.pathEntriesLen - 1] = entry;
}

void TreePathWrite(TreePath *treePath, Buffer *buf) {
	if (treePath == 0) {
		BufOptionalWrite(buf);
		return;
	}

	if (treePath.pathEntries && treePath.pathEntriesLen > 0) {
		BufSetTypeAndLength(buf, TYPELIST, treePath.pathEntriesLen);

		int i;
		for (i = 0; i < treePath.pathEntriesLen; i++) {
			OctetStringWrite(treePath.pathEntries[i], buf);
		}
	}
}

void TreePathFree(TreePath *treePath) {
	if (treePath) {
		if (treePath.pathEntries && treePath.pathEntriesLen > 0) {
			int i;
			for (i = 0; i < treePath.pathEntriesLen; i++) {
				OctetStringFree(treePath.pathEntries[i]);
			}

			free(treePath.pathEntries);
		}

		free(treePath);
	}
}


// Tree;

Tree *TreeInit() {
	Tree *tree = (Tree *) malloc(sizeof(Tree));
	*tree = ( Tree ) {
		.parameterName = NULL,
		.parameterValue = NULL,
		.childList = NULL,
		.childListLen = 0
	};

	return tree;
}

Tree *TreeParse(Buffer *buf) {
	if (BufOptionalIsSkipped(buf)) {
		return 0;
	}

	Tree *tree = TreeInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 3) {
		buf.error = 1;
		goto error;
	}

	tree.parameterName = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;

	tree.parameterValue = ProcParValueParse(buf);
	if (BufHasErrors(buf)) goto error;

	if (!BufOptionalIsSkipped(buf)) {
		if (BufGetNextType(buf) != TYPELIST) {
			buf.error = 1;
			goto error;
		}

		Tree *c;
		int elems;
		for (elems = BufGetNextLength(buf); elems > 0; elems--) {
			c = TreeParse(buf);
			if (BufHasErrors(buf)) goto error;
			if (c) {
				TreeAddTree(tree, c);
			}
		}
	}

	return tree;

error:
	TreeFree(tree);
	return 0;
}

void TreeAddTree(Tree *baseTree, Tree *tree) {
	baseTree.childListLen++;
	baseTree.childList = (Tree **) realloc(baseTree.childList,
		sizeof(Tree *) * baseTree.childListLen);
	baseTree.childList[baseTree.childListLen - 1] = tree;
}

void TreeFree(Tree *tree) {
	if (tree) {
		OctetStringFree(tree.parameterName);
		ProcParValueFree(tree.parameterValue);
		int i;
		for (i = 0; i < tree.childListLen; i++) {
			TreeFree(tree.childList[i]);
		}

		free(tree.childList);
		free(tree);
	}
}

void TreeWrite(Tree *tree, Buffer *buf) {
	if (tree == 0) {
		BufOptionalWrite(buf);
		return;
	}

	BufSetTypeAndLength(buf, TYPELIST, 3);

	OctetStringWrite(tree.parameterName, buf);
	ProcParValueWrite(tree.parameterValue, buf);

	if (tree.childList && tree.childListLen > 0) {
		BufSetTypeAndLength(buf, TYPELIST, tree.childListLen);

		int i;
		for (i = 0; i < tree.childListLen; i++) {
			TreeWrite(tree.childList[i], buf);
		}
	}
	else {
		BufOptionalWrite(buf);
	}
}


// ProcParValue;

ProcParValue *ProcParValueInit() {
	ProcParValue *value = (ProcParValue *) malloc(sizeof(ProcParValue));
	*value = ( ProcParValue ) {
		.tag = NULL,
		.data.value = NULL
	};
	return value;
}

ProcParValue *ProcParValueParse(Buffer *buf) {
	if (BufOptionalIsSkipped(buf)) {
		return 0;
	}

	ProcParValue *ppv = ProcParValueInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 2) {
		buf.error = 1;
		goto error;
	}

	ppv.tag = U8Parse(buf);
	if (BufHasErrors(buf)) goto error;

	switch (*(ppv.tag)) {
		case PROCPARVALUETAGVALUE:
			ppv.data.value = ValueParse(buf);
			break;
		case PROCPARVALUETAGPERIODENTRY:
			ppv.data.periodEntry = PeriodEntryParse(buf);
			break;
		case PROCPARVALUETAGTUPELENTRY:
			ppv.data.tupelEntry = TupelEntryParse(buf);
			break;
		case PROCPARVALUETAGTIME:
			ppv.data.time = TimeParse(buf);
			break;
		default:
			buf.error = 1;
			goto error;
	}

	return ppv;

error:
	ProcParValueFree(ppv);
	return 0;
}

void ProcParValueWrite(ProcParValue *value, Buffer *buf) {
	if (value == 0) {
		BufOptionalWrite(buf);
		return;
	}

	BufSetTypeAndLength(buf, TYPELIST, 2);
	U8Write(value.tag, buf);

	switch (*(value.tag)) {
		case PROCPARVALUETAGVALUE:
			ValueWrite(value.data.value, buf);
			break;
		case PROCPARVALUETAGPERIODENTRY:
			PeriodEntryWrite(value.data.periodEntry, buf);
			break;
		case PROCPARVALUETAGTUPELENTRY:
			TupelEntryWrite(value.data.tupelEntry, buf);
			break;
		case PROCPARVALUETAGTIME:
			TimeWrite(value.data.time, buf);
			break;
		default:
			fprintf(stderr,"libsml: error: unknown tag in %s\n", "ProcParValueWrite");
	}
}

void ProcParValueFree(ProcParValue *ppv) {
	if (ppv) {
		if (ppv.tag) {
			switch (*(ppv.tag)) {
				case PROCPARVALUETAGVALUE:
					ValueFree(ppv.data.value);
					break;
				case PROCPARVALUETAGPERIODENTRY:
					PeriodEntryFree(ppv.data.periodEntry);
					break;
				case PROCPARVALUETAGTUPELENTRY:
					TupelEntryFree(ppv.data.tupelEntry);
					break;
				case PROCPARVALUETAGTIME:
					TimeFree(ppv.data.time);
					break;
				default:
					if (ppv.data.value) {
						free(ppv.data.value);
					}
			}
			NumberFree(ppv.tag);
		}
		else {
			// Without the tag, there might be a memory leak.
			if (ppv.data.value) {
				free(ppv.data.value);
			}
		}

		free(ppv);
	}
}


// TupleEntry;

TupelEntry *TupelEntryInit() {
	TupelEntry *tupel = (TupelEntry *) malloc(sizeof(TupelEntry));
	*tupel = ( TupelEntry ) {
		.serverId = NULL,
		.secIndex = NULL,
		.status = NULL,
		.unitPA = NULL,
		.scalerPA = NULL,
		.valuePA = NULL,
		.unitR1 = NULL,
		.scalerR1 = NULL,
		.valueR1 = NULL,
		.unitR4 = NULL,
		.scalerR4 = NULL,
		.valueR4 = NULL,
		.signaturePAR1R4 = NULL,
		.unitMA = NULL,
		.scalerMA = NULL,
		.valueMA = NULL,
		.unitR2 = NULL,
		.scalerR2 = NULL,
		.valueR2 = NULL,
		.unitR3 = NULL,
		.scalerR3 = NULL,
		.valueR3 = NULL,
		.signatureMAR2R3 = NULL
	};

	return tupel;
}

TupelEntry *TupelEntryParse(Buffer *buf) {
	if (BufOptionalIsSkipped(buf)) {
		return 0;
	}

	TupelEntry *tupel = TupelEntryInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 23) {
		buf.error = 1;
		goto error;
	}

	tupel.serverId = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.secIndex = TimeParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.status = U64Parse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.unitPA = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.scalerPA = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.valuePA = I64Parse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.unitR1 = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.scalerR1 = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.valueR1 = I64Parse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.unitR4 = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.scalerR4 = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.valueR4 = I64Parse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.signaturePAR1R4 = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.unitMA = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.scalerMA = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.valueMA = I64Parse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.unitR2 = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.scalerR2 = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.valueR2 = I64Parse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.unitR3 = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.scalerR3 = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;
	tupel.valueR3 = I64Parse(buf);
	if (BufHasErrors(buf)) goto error;

	tupel.signatureMAR2R3 = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;

	return tupel;

error:
	TupelEntryFree(tupel);
	return 0;
}

void TupelEntryWrite(TupelEntry *tupel, Buffer *buf) {
	if (tupel == 0) {
		BufOptionalWrite(buf);
		return;
	}

	BufSetTypeAndLength(buf, TYPELIST, 23);

	OctetStringWrite(tupel.serverId, buf);
	TimeWrite(tupel.secIndex, buf);
	U64Write(tupel.status, buf);

	UnitWrite(tupel.unitPA, buf);
	I8Write(tupel.scalerPA, buf);
	I64Write(tupel.valuePA, buf);

	UnitWrite(tupel.unitR1, buf);
	I8Write(tupel.scalerR1, buf);
	I64Write(tupel.valueR1, buf);

	UnitWrite(tupel.unitR4, buf);
	I8Write(tupel.scalerR4, buf);
	I64Write(tupel.valueR4, buf);

	OctetStringWrite(tupel.signaturePAR1R4, buf);

	UnitWrite(tupel.unitMA, buf);
	I8Write(tupel.scalerMA, buf);
	I64Write(tupel.valueMA, buf);

	UnitWrite(tupel.unitR2, buf);
	I8Write(tupel.scalerR2, buf);
	I64Write(tupel.valueR2, buf);

	UnitWrite(tupel.unitR3, buf);
	I8Write(tupel.scalerR3, buf);
	I64Write(tupel.valueR3, buf);

	OctetStringWrite(tupel.signatureMAR2R3, buf);
}

void TupelEntryFree(TupelEntry *tupel) {
	if (tupel) {
		OctetStringFree(tupel.serverId);
		TimeFree(tupel.secIndex);
		NumberFree(tupel.status);

		UnitFree(tupel.unitPA);
		NumberFree(tupel.scalerPA);
		NumberFree(tupel.valuePA);

		UnitFree(tupel.unitR1);
		NumberFree(tupel.scalerR1);
		NumberFree(tupel.valueR1);

		UnitFree(tupel.unitR4);
		NumberFree(tupel.scalerR4);
		NumberFree(tupel.valueR4);

		OctetStringFree(tupel.signaturePAR1R4);

		UnitFree(tupel.unitMA);
		NumberFree(tupel.scalerMA);
		NumberFree(tupel.valueMA);

		UnitFree(tupel.unitR2);
		NumberFree(tupel.scalerR2);
		NumberFree(tupel.valueR2);

		UnitFree(tupel.unitR3);
		NumberFree(tupel.scalerR3);
		NumberFree(tupel.valueR3);

		OctetStringFree(tupel.signatureMAR2R3);

		free(tupel);
	}
}



// PeriodEntry;

PeriodEntry *PeriodEntryInit() {
	PeriodEntry *period = (PeriodEntry *) malloc(sizeof(PeriodEntry));
	*period = ( PeriodEntry ) {
		.objName = NULL,
		.unit = NULL,
		.scaler = NULL,
		.value = NULL,
		.valueSignature = NULL
	};

	return period;
}

static void * PeriodEntryParse(Buffer *buf) {
	if (BufOptionalIsSkipped(buf)) {
		return 0;
	}

	PeriodEntry *period = PeriodEntryInit();

	if (BufGetNextType(buf) != TYPELIST) {
		buf.error = 1;
		goto error;
	}

	if (BufGetNextLength(buf) != 5) {
		buf.error = 1;
		goto error;
	}

	period.objName = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;

	period.unit = UnitParse(buf);
	if (BufHasErrors(buf)) goto error;

	period.scaler = I8Parse(buf);
	if (BufHasErrors(buf)) goto error;

	period.value = ValueParse(buf);
	if (BufHasErrors(buf)) goto error;

	period.valueSignature = OctetStringParse(buf);
	if (BufHasErrors(buf)) goto error;

	return period;

error:
	PeriodEntryFree(period);
	return 0;
}

PeriodEntry * PeriodEntryParse( Buffer * buf ) {
	return PeriodEntryParse( buf );
}
