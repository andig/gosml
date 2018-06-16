package sml

import (
	"github.com/pkg/errors"
)

func StatusParse(buf *Buffer) (int64, error) {
	/*
		if (BufOptionalIsSkipped(buf)) {
			return 0;
		}

		int max = 1;
		int type = BufGetNextType(buf);
		unsigned char byte = BufGetCurrentByte(buf);

		Status *status = StatusInit();
		status->type = type;
		switch (type) {
			case TYPEUNSIGNED:
				// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
				while (max < ((byte & LENGTHFIELD) - 1)) {
					max <<= 1;
				}

				status->data.status8 = NumberParse(buf, type, max);
				status->type |= max;
				break;
			default:
				buf->error = 1;
				break;
		}
	*/
	// TODO proper type handling

	if skip := BufOptionalIsSkipped(buf); skip {
		return 0, nil
	}

	Debug(buf, "StatusParse")

	var max uint8 = 1
	var status8 int64
	typefield := BufGetNextType(buf)
	statusType := typefield
	b := BufGetCurrentByte(buf)

	if typefield == TYPEUNSIGNED {
		// get maximal size, if not all bytes are used (example: only 6 bytes for a u64)
		for max < ((b & LENGTHFIELD) - 1) {
			max = max << 1
		}

		if _, err := NumberParse(buf, typefield, int(max)); err != nil {
			return 0, err
		}

		statusType = statusType | max
	} else {
		return 0, errors.Errorf("Unexpected type %02x (expected %02x)", typefield, TYPEUNSIGNED)
	}

	return status8, nil
}
