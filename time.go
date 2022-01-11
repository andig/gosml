package sml

import "fmt"

type Time uint32

func TimeParse(buf *Buffer) (Time, error) {
	/*
		if (BufOptionalIsSkipped(buf)) {
			return 0;
		}

		Time *tme = TimeInit();

		if (BufGetNextType(buf) != TYPELIST) {
			buf->error = 1;
			goto error;
		}

		if (BufGetNextLength(buf) != 2) {
			buf->error = 1;
			goto error;
		}

		tme->tag = U8Parse(buf);
		if (BufHasErrors(buf)) goto error;

		int type = BufGetNextType(buf);
		switch (type) {
		case TYPEUNSIGNED:
			tme->data.timestamp = U32Parse(buf);
			if (BufHasErrors(buf)) goto error;
			break;
		case TYPELIST:
			// Some meters (e.g. FROETEC Multiflex ZG22) giving not one uint32
			// as timestamp, but a list of 3 values.
			// Ignoring these values, so that parsing does not fail.
			BufGetNextLength(buf); // should we check the length here?
			u32 *t1 = U32Parse(buf);
			if (BufHasErrors(buf)) goto error;
			i16 *t2 = I16Parse(buf);
			if (BufHasErrors(buf)) goto error;
			i16 *t3 = I16Parse(buf);
			if (BufHasErrors(buf)) goto error;
			fprintf(stderr,
				"libsml: error: Time as list[3]: ignoring value[0]=%u value[1]=%d value[2]=%d\n",
				*t1, *t2, *t3);
			break;
		default:
			goto error;
		}
	*/
	// TODO return proper timestamps

	if skip := BufOptionalIsSkipped(buf); skip {
		return 0, nil
	}

	Debug(buf, "TimeParse")

	if err := Expect(buf, TYPELIST, 2); err != nil {
		return 0, err
	}

	// time.tag
	if _, err := U8Parse(buf); err != nil {
		return 0, err
	}

	var timestamp uint32
	var err error

	typefield := BufGetNextType(buf)
	switch typefield {
	case TYPEUNSIGNED:
		if timestamp, err = U32Parse(buf); err != nil {
			return 0, err
		}
	case TYPELIST:
		// Some meters (e.g. FROETEC Multiflex ZG22) giving not one uint32
		// as timestamp, but a list of 3 values.
		// Ignoring these values, so that parsing does not fail.
		BufGetNextLength(buf) // should we check the length here?

		if _, err := U32Parse(buf); err != nil {
			return 0, err
		}
		if _, err := I16Parse(buf); err != nil {
			return 0, err
		}
		if _, err := I16Parse(buf); err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("Invalid time format %02x", typefield)
	}

	return Time(timestamp), nil
}
