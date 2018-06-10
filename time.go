package sml

import (
	"fmt"
)

type sml_time uint32

func sml_time_parse(buf *sml_buffer) (sml_time, error) {
	/*
	if (sml_buf_optional_is_skipped(buf)) {
		return 0;
	}

	sml_time *tme = sml_time_init();

	if (sml_buf_get_next_type(buf) != SML_TYPE_LIST) {
		buf->error = 1;
		goto error;
	}

	if (sml_buf_get_next_length(buf) != 2) {
		buf->error = 1;
		goto error;
	}

	tme->tag = sml_u8_parse(buf);
	if (sml_buf_has_errors(buf)) goto error;

	int type = sml_buf_get_next_type(buf);
	switch (type) {
	case SML_TYPE_UNSIGNED:
		tme->data.timestamp = sml_u32_parse(buf);
		if (sml_buf_has_errors(buf)) goto error;
		break;
	case SML_TYPE_LIST:
		// Some meters (e.g. FROETEC Multiflex ZG22) giving not one uint32
		// as timestamp, but a list of 3 values.
		// Ignoring these values, so that parsing does not fail.
		sml_buf_get_next_length(buf); // should we check the length here?
		u32 *t1 = sml_u32_parse(buf);
		if (sml_buf_has_errors(buf)) goto error;
		i16 *t2 = sml_i16_parse(buf);
		if (sml_buf_has_errors(buf)) goto error;
		i16 *t3 = sml_i16_parse(buf);
		if (sml_buf_has_errors(buf)) goto error;
		fprintf(stderr,
			"libsml: error: sml_time as list[3]: ignoring value[0]=%u value[1]=%d value[2]=%d\n",
			*t1, *t2, *t3);
		break;
	default:
		goto error;
	}
*/
	// TODO return proper timestamps
	
	if skip := sml_buf_optional_is_skipped(buf); skip {
		return 0, nil
	}

	sml_debug(buf, "sml_time_parse")

	if err := sml_expect(buf, SML_TYPE_LIST, 2); err != nil {
		return 0, err
	}

	// time.tag
	if _, err := sml_u8_parse(buf); err != nil {
		return 0, err
	}

	var timestamp uint32
	var err error

	typefield := sml_buf_get_next_type(buf)
	switch (typefield) {
	case SML_TYPE_UNSIGNED:
		if timestamp, err = sml_u32_parse(buf); err != nil {
			return 0, err
		}
	case SML_TYPE_LIST:
		// Some meters (e.g. FROETEC Multiflex ZG22) giving not one uint32
		// as timestamp, but a list of 3 values.
		// Ignoring these values, so that parsing does not fail.
		sml_buf_get_next_length(buf) // should we check the length here?

		if _, err := sml_u32_parse(buf); err != nil {
			return 0, err
		}
		if _, err := sml_i16_parse(buf); err != nil {
			return 0, err
		}
		if _, err := sml_i16_parse(buf); err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("Invalid time format %02x", typefield)
	}

	return sml_time(timestamp), nil
}
