package sml

import (
	"fmt"
)

func sml_status_parse(buf *sml_buffer) (int64, error) {
/*
	if (sml_buf_optional_is_skipped(buf)) {
		return 0;
	}

	int max = 1;
	int type = sml_buf_get_next_type(buf);
	unsigned char byte = sml_buf_get_current_byte(buf);

	sml_status *status = sml_status_init();
	status->type = type;
	switch (type) {
		case SML_TYPE_UNSIGNED:
			// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
			while (max < ((byte & SML_LENGTH_FIELD) - 1)) {
				max <<= 1;
			}

			status->data.status8 = sml_number_parse(buf, type, max);
			status->type |= max;
			break;
		default:
			buf->error = 1;
			break;
	}
*/
	// TODO proper type handling

	if skip := sml_buf_optional_is_skipped(buf); skip {
		return 0, nil
	}

	sml_debug(buf, "sml_status_parse")

	var max uint8 = 1
	var status8 int64
	typefield := sml_buf_get_next_type(buf)
	status_type := typefield
	b := sml_buf_get_current_byte(buf)

	if typefield == SML_TYPE_UNSIGNED {
		// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
		for max < ((b & SML_LENGTH_FIELD) - 1) {
			max = max << 1
		}

		if _, err := sml_number_parse(buf, typefield, int(max)); err != nil {
			return 0, err
		}

		status_type = status_type | max
	} else {
		return 0, fmt.Errorf("sml: Unexpected type %02x (expected %02x)", typefield, SML_TYPE_UNSIGNED)
	}

	return status8, nil
}
